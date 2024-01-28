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

func (expectation *Expectation[A]) Not() *Expectation[A] {
	expectation.test.testingT.Helper()

	expectation.reverseExpectation = true
	return expectation
}

// TODO: まだ調整が必要: failMessageが %sになっているため、文字列がおかしくなる
func (expectation *Expectation[A]) FailMsg(msg string) *Expectation[A] {
	expectation.test.testingT.Helper()

	expectation.failMsg = msg
	return expectation
}

// TODO: 引数にカスタムのメッセージをいれられるようにする
func (expectation *Expectation[A]) ToBeTrue() {
	expectation.test.testingT.Helper()
	expectation.handleBoolReverseExpectation(true)
}

func (expectation *Expectation[A]) ToBeFalse() {
	expectation.test.testingT.Helper()
	expectation.handleBoolReverseExpectation(false)
}

func (expectation *Expectation[A]) handleBoolReverseExpectation(value bool) {
	expectation.test.testingT.Helper()

	failMsg := "Expected [%%v] to be %v"
	if expectation.failMsg != "" {
		failMsg = expectation.failMsg
	}

	if expectation.reverseExpectation {
		expectation.checkBool(!value, fmt.Sprintf(failMsg, !value))
		return
	}
	expectation.checkBool(value, fmt.Sprintf(failMsg, value))
}

func (expectation *Expectation[A]) checkBool(value bool, errorMsg string) {
	expectation.test.testingT.Helper()

	// `3` is where ToBeTrue() or ToBeFalse() is called
	_, file, line, _ := runtime.Caller(3)
	relPath := extractRelPath(file)

	expectation.test.subtotal++
	switch v := any(*expectation.actual).(type) {
	case bool:
		if v != value {
			expectation.processFailure(relPath, line, errorMsg)
			return
		}
		expectation.processPassed()
	default:
		// TODO: ここもメッセージ変えられるようにする
		msgNotBool := "Expected [%v] to be a bool value"
		expectation.processFailure(relPath, line, msgNotBool)
	}
}

func (expectation *Expectation[A]) ToBeNil() {
	expectation.test.testingT.Helper()

	_, file, line, _ := runtime.Caller(1)
	relPath := extractRelPath(file)

	failMsg := "Expected [%v] to be nil"
	if expectation.failMsg != "" {
		failMsg = expectation.failMsg
	}

	expectation.test.subtotal++
	// REFACTOR:
	if expectation.reverseExpectation {
		if expectation.actual != nil {
			expectation.processPassed()
			return
		}
		// FIXME: Not()を使って、ここの場合メッセージが %v の部分が正しくないため修正
		expectation.processFailure(relPath, line, failMsg)
	} else {
		if expectation.actual != nil {
			expectation.processFailure(relPath, line, failMsg)
			return
		}
		expectation.processPassed()
	}
}

// assertion for primitive values
func (expectation *Expectation[A]) ToBe(expectationValue A) {
	expectation.test.testingT.Helper()

	_, file, line, _ := runtime.Caller(1)
	relPath := extractRelPath(file)

	failMsg := "Expected [%v] to be [%v]"
	if expectation.failMsg != "" {
		failMsg = expectation.failMsg
	}

	expectation.test.subtotal++
	switch v := any(*expectation.actual).(type) {
	case int:
		expectationValueInt := any(expectationValue).(int)
		if expectation.reverseExpectation {
			if v == expectationValueInt {
				expectation.processFailure(relPath, line, failMsg)
				return
			}
			expectation.processPassed()
		} else {
			if v != expectationValueInt {
				expectation.processFailure(relPath, line, failMsg)
				return
			}
			expectation.processPassed()
		}
	case string:
		expectationValueString := any(expectationValue).(string)
		if expectation.reverseExpectation {
			if v == expectationValueString {
				expectation.processFailure(relPath, line, failMsg)
				return
			}
			expectation.processPassed()
		} else {
			if v != expectationValueString {
				expectation.processFailure(relPath, line, failMsg)
				return
			}
			expectation.processPassed()
		}
	}
}

// TODO: expectがfailしたファイル・行を出力するには、expectation.test.testingT.Helper()を使う必要がある。
// すべての関数で、最初に expectation.test.testingT.Helper() を呼ぶこと
// TODO: それぞれのアサーションには、引数でカスタムのメッセージを指定できるようにすること

// func ToEqual()
// func PointerEquals()
// func ContainString()
// func ToMatchRegex()
// assert json value the same

func (expectation Expectation[A]) processFailure(relPath string, line int, errorMsg string) {
	expectation.test.testingT.Helper()

	msg := expectation.failMessage(relPath, line, errorMsg)
	expectation.test.testingT.Errorf(RedMsg(msg))
	expectation.markAsFailed()
	expectation.resetNot()
}

func (expectation *Expectation[A]) processPassed() {
	expectation.test.testingT.Helper()

	expectation.test.passed++
	expectation.resetNot()
}

func (expectation *Expectation[A]) markAsFailed() {
	expectation.test.testingT.Helper()

	expectation.test.isThisTestFailed = true
	expectation.test.isAnyTestFailed = true
}

func (expectation *Expectation[A]) failMessage(relPath string, line int, errorMsg string) string {
	expectation.test.testingT.Helper()

	if expectation.actual != nil {
		return fmt.Sprintf("Failed at [%s]:line %d: %s", relPath, line, fmt.Sprintf(errorMsg, *expectation.actual))
	}
	return fmt.Sprintf("Failed at [%s]:line %d: %s", relPath, line, fmt.Sprintf(errorMsg))
}

func (expectation *Expectation[A]) resetNot() {
	expectation.test.testingT.Helper()

	expectation.reverseExpectation = false
}
