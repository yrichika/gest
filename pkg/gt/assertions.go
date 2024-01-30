package gt

import (
	"fmt"
	"reflect"
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
// Expectation constructor
func Expect[A any](test *Test, actual *A) *Expectation[A] {
	return &Expectation[A]{
		actual:             actual,
		test:               test,
		failMsg:            "",
		reverseExpectation: false,
	}
}

// TODO: まだ調整が必要: failMessageが %sになっているため、文字列がおかしくなる
// Expectation constructor with fail message
func WhenFailPrint[A any](test *Test, failMsg string) *Expectation[A] {
	return &Expectation[A]{
		test:               test,
		failMsg:            failMsg,
		reverseExpectation: false,
	}
}

// Use this ONLY AFTER you call `WhenFailPrint()` or any other Expectation constructors if exist
func (expectation *Expectation[A]) Expect(actual *A) *Expectation[A] {
	expectation.test.testingT.Helper()

	expectation.actual = actual
	return expectation
}

func (expectation *Expectation[A]) Not() *Expectation[A] {
	expectation.test.testingT.Helper()

	expectation.reverseExpectation = true
	return expectation
}

func (expectation *Expectation[A]) ToBeNil() {
	expectation.test.testingT.Helper()

	relPath, line := getTestInfo(1)

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
		return
	}

	if expectation.actual == nil {
		expectation.processPassed()
		return
	}
	expectation.processFailure(relPath, line, failMsg)

}

// assertion for primitive values
func (expectation *Expectation[A]) ToBe(expected A) {
	expectation.test.testingT.Helper()

	failMsg := "Expected [%v] to be [%v]"
	if expectation.failMsg != "" {
		failMsg = expectation.failMsg
	}

	expectation.test.subtotal++
	switch actual := any(*expectation.actual).(type) {
	case int:
		expectation.intEq(actual, expected, failMsg)
	case int8:
		expectation.int8Eq(actual, expected, failMsg)
	case int16:
		expectation.int16Eq(actual, expected, failMsg)
	case int32: // rune
		expectation.int32Eq(actual, expected, failMsg)
	case int64:
		expectation.int64Eq(actual, expected, failMsg)
	case uint:
		expectation.uintEq(actual, expected, failMsg)
	case uint8: // byte
		expectation.uint8Eq(actual, expected, failMsg)
	case uint16:
		expectation.uint16Eq(actual, expected, failMsg)
	case uint32:
		expectation.uint32Eq(actual, expected, failMsg)
	case uint64:
		expectation.uint64Eq(actual, expected, failMsg)
	case uintptr:
		expectation.uintptrEq(actual, expected, failMsg)
	case float32:
		expectation.float32Eq(actual, expected, failMsg)
	case float64:
		expectation.float64Eq(actual, expected, failMsg)
	case bool:
		expectation.boolEq(actual, expected, failMsg)
	case string:
		expectation.stringEq(actual, expected, failMsg)
	case complex64:
		expectation.complex64Eq(actual, expected, failMsg)
	case complex128:
		expectation.complex128Eq(actual, expected, failMsg)
	default:
		// TODO:
	}
}

// two structs equality
func (expectation *Expectation[A]) ToDeepEqual(expected A) {
	expectation.test.testingT.Helper()

	relPath, line := getTestInfo(1)
	// FIXME: メッセージ修正
	failMsg := "Expected [%v] to be [%v]"
	if expectation.failMsg != "" {
		failMsg = expectation.failMsg
	}

	expectation.test.subtotal++
	if expectation.reverseExpectation {
		if !reflect.DeepEqual(*expectation.actual, expected) {
			expectation.processPassed()
			return
		}
		expectation.processFailure(relPath, line, failMsg)
		return
	}

	if reflect.DeepEqual(*expectation.actual, expected) {
		expectation.processPassed()
		return
	}
	expectation.processFailure(relPath, line, failMsg)
}

func (expectation *Expectation[A]) ToBeSamePointerAs(expected *A) {
	expectation.test.testingT.Helper()

	relPath, line := getTestInfo(1)
	// FIXME: メッセージ修正
	failMsg := "Expected [%v] to be [%v]"
	if expectation.failMsg != "" {
		failMsg = expectation.failMsg
	}

	expectation.test.subtotal++
	if expectation.reverseExpectation {
		if expectation.actual != expected {
			expectation.processPassed()
			return
		}
		expectation.processFailure(relPath, line, failMsg)
		return
	}

	if expectation.actual == expected {
		expectation.processPassed()
		return
	}
	expectation.processFailure(relPath, line, failMsg)

}

// func ToContainString()
// func ToMatchRegex()
// assert json value the same

func (expectation *Expectation[A]) intEq(actual int, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[int](
		actual,
		any(expected).(int),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation *Expectation[A]) int8Eq(actual int8, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[int8](
		actual,
		any(expected).(int8),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation *Expectation[A]) int16Eq(actual int16, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[int16](
		actual,
		any(expected).(int16),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) int32Eq(actual int32, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[int32](
		actual,
		any(expected).(int32),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) int64Eq(actual int64, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[int64](
		actual,
		any(expected).(int64),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) uintEq(actual uint, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[uint](
		actual,
		any(expected).(uint),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) uint8Eq(actual uint8, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[uint8](
		actual,
		any(expected).(uint8),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) uint16Eq(actual uint16, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[uint16](
		actual,
		any(expected).(uint16),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) uint32Eq(actual uint32, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[uint32](
		actual,
		any(expected).(uint32),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) uint64Eq(actual uint64, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[uint64](
		actual,
		any(expected).(uint64),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) uintptrEq(actual uintptr, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[uintptr](
		actual,
		any(expected).(uintptr),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) float32Eq(actual float32, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[float32](
		actual,
		any(expected).(float32),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) float64Eq(actual float64, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[float64](
		actual,
		any(expected).(float64),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) boolEq(actual bool, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[bool](
		actual,
		any(expected).(bool),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) complex64Eq(actual complex64, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[complex64](
		actual,
		any(expected).(complex64),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) complex128Eq(actual complex128, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[complex128](
		actual,
		any(expected).(complex128),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation *Expectation[A]) stringEq(actual string, expected A, failMsg string) {
	expectation.test.testingT.Helper()

	assertEq[string](
		actual,
		any(expected).(string),
		expectation.reverseExpectation,
		failMsg,
		expectation.processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

// REFACTOR: too many arguments. there should be a better way
func assertEq[T comparable](
	actual T,
	expected T,
	reverse bool,
	failMsg string,
	failFunc func(string, int, string),
	passFunc func(),
	Helper func(),
) {
	Helper()

	relPath, line := getTestInfo(3)

	if reverse {
		if actual == expected {
			failFunc(relPath, line, failMsg)
			return
		}
		passFunc()
	} else {
		if actual != expected {
			failFunc(relPath, line, failMsg)
			return
		}
		passFunc()
	}
}

func getTestInfo(skip int) (string, int) {
	skip++ // add one for this function
	_, file, line, _ := runtime.Caller(skip)
	relPath := extractRelPath(file)
	return relPath, line
}

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
