package gt

import (
	"fmt"
	"runtime"
)

type Expectation[A any] struct {
	actual             *A
	test               *Test
	failMsg            string
	reverseExpectation bool
}

// TODO: testを常に引数に渡さないようにしたい。
// ただし、メソッドに型パラメータをもたせられないため、方法がわからない。
func Expect[A any](test *Test, actual *A) *Expectation[A] {
	return &Expectation[A]{
		actual:             actual,
		test:               test,
		failMsg:            "",
		reverseExpectation: false,
	}
}

func (expected *Expectation[A]) Not() *Expectation[A] {
	expected.test.testingT.Helper()

	expected.reverseExpectation = true
	return expected
}

// TODO: まだ調整が必要: failMessageが %sになっているため、文字列がおかしくなる
func (expected *Expectation[A]) FailMsg(msg string) *Expectation[A] {
	expected.test.testingT.Helper()

	expected.failMsg = msg
	return expected
}

// TODO: 引数にカスタムのメッセージをいれられるようにする
func (expected *Expectation[A]) ToBeTrue() {
	expected.test.testingT.Helper()
	expected.handleBoolReverseExpectation(true)
}

func (expected *Expectation[A]) ToBeFalse() {
	expected.test.testingT.Helper()
	expected.handleBoolReverseExpectation(false)
}

func (expected *Expectation[A]) handleBoolReverseExpectation(value bool) {
	expected.test.testingT.Helper()

	failMsg := "Expected [%%v] to be %v"
	if expected.failMsg != "" {
		failMsg = expected.failMsg
	}

	if expected.reverseExpectation {
		expected.checkBool(!value, fmt.Sprintf(failMsg, !value))
		return
	}
	expected.checkBool(value, fmt.Sprintf(failMsg, value))
}

func (expected *Expectation[A]) checkBool(value bool, errorMsg string) {
	expected.test.testingT.Helper()

	// `3` is where ToBeTrue() or ToBeFalse() is called
	_, file, line, _ := runtime.Caller(3)
	relPath := extractRelPath(file)

	expected.test.subtotal++
	switch v := any(*expected.actual).(type) {
	case bool:
		if v != value {
			expected.processFailure(relPath, line, errorMsg)
			return
		}
		expected.processPassed()
	default:
		// TODO: ここもメッセージ変えられるようにする
		msgNotBool := "Expected [%v] to be a bool value"
		expected.processFailure(relPath, line, msgNotBool)
	}
}

func (expected *Expectation[A]) ToBeNil() {
	expected.test.testingT.Helper()

	_, file, line, _ := runtime.Caller(1)
	relPath := extractRelPath(file)

	failMsg := "Expected [%v] to be nil"
	if expected.failMsg != "" {
		failMsg = expected.failMsg
	}

	expected.test.subtotal++
	// REFACTOR:
	if expected.reverseExpectation {
		if expected.actual != nil {
			expected.processPassed()
			return
		}
		// FIXME: Not()を使って、ここの場合メッセージが %v の部分が正しくないため修正
		expected.processFailure(relPath, line, failMsg)
	} else {
		if expected.actual != nil {
			expected.processFailure(relPath, line, failMsg)
			return
		}
		expected.processPassed()
	}
}

// assertion for primitive values
func (expected *Expectation[A]) ToBe(expectedValue A) {
	expected.test.testingT.Helper()

	_, file, line, _ := runtime.Caller(1)
	relPath := extractRelPath(file)

	failMsg := "Expected [%v] to be [%v]"
	if expected.failMsg != "" {
		failMsg = expected.failMsg
	}

	expected.test.subtotal++
	switch v := any(*expected.actual).(type) {
	case int:
		expectedValueInt := any(expectedValue).(int)
		if expected.reverseExpectation {
			if v == expectedValueInt {
				expected.processFailure(relPath, line, failMsg)
				return
			}
			expected.processPassed()
		} else {
			if v != expectedValueInt {
				expected.processFailure(relPath, line, failMsg)
				return
			}
			expected.processPassed()
		}
	case string:
		expectedValueString := any(expectedValue).(string)
		if expected.reverseExpectation {
			if v == expectedValueString {
				expected.processFailure(relPath, line, failMsg)
				return
			}
			expected.processPassed()
		} else {
			if v != expectedValueString {
				expected.processFailure(relPath, line, failMsg)
				return
			}
			expected.processPassed()
		}
	}
}

// TODO: expectがfailしたファイル・行を出力するには、expected.test.testingT.Helper()を使う必要がある。
// すべての関数で、最初に expected.test.testingT.Helper() を呼ぶこと
// TODO: それぞれのアサーションには、引数でカスタムのメッセージを指定できるようにすること

// func ToEqual()
// func PointerEquals()
// func ContainString()
// func ToMatchRegex()
// assert json value the same

func (expected Expectation[A]) processFailure(relPath string, line int, errorMsg string) {
	expected.test.testingT.Helper()

	msg := expected.failMessage(relPath, line, errorMsg)
	expected.test.testingT.Errorf(RedMsg(msg))
	expected.markAsFailed()
	expected.resetNot()
}

func (expected *Expectation[A]) processPassed() {
	expected.test.testingT.Helper()

	expected.test.passed++
	expected.resetNot()
}

func (expected *Expectation[A]) markAsFailed() {
	expected.test.testingT.Helper()

	expected.test.isThisTestFailed = true
	expected.test.isAnyTestFailed = true
}

func (expected *Expectation[A]) failMessage(relPath string, line int, errorMsg string) string {
	expected.test.testingT.Helper()

	if expected.actual != nil {
		return fmt.Sprintf("Failed at [%s]:line %d: %s", relPath, line, fmt.Sprintf(errorMsg, *expected.actual))
	}
	return fmt.Sprintf("Failed at [%s]:line %d: %s", relPath, line, fmt.Sprintf(errorMsg))
}

func (expected *Expectation[A]) resetNot() {
	expected.test.testingT.Helper()

	expected.reverseExpectation = false
}
