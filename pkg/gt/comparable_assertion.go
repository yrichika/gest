package gt

import (
	"reflect"
	"time"
)

// Assertion for primitive values, time.Time and custom struct types.
// Types must be the same.
// プリミティブ型の値を比較します。値が同じ場合はアサートがpassします。
// アサート対象の値と、想定している値は同じ型である必要があります。
func (expectation *Expectation[A]) ToBe(expected A) {
	expectation.test.testingT.Helper()

	expectation.test.subtotal++
	switch actual := any(*expectation.actual).(type) {
	case int:
		expectation.intEq(actual, expected)
	case int8:
		expectation.int8Eq(actual, expected)
	case int16:
		expectation.int16Eq(actual, expected)
	case int32: // rune
		expectation.int32Eq(actual, expected)
	case int64:
		expectation.int64Eq(actual, expected)
	case uint:
		expectation.uintEq(actual, expected)
	case uint8: // byte
		expectation.uint8Eq(actual, expected)
	case uint16:
		expectation.uint16Eq(actual, expected)
	case uint32:
		expectation.uint32Eq(actual, expected)
	case uint64:
		expectation.uint64Eq(actual, expected)
	case uintptr:
		expectation.uintptrEq(actual, expected)
	case float32:
		expectation.float32Eq(actual, expected)
	case float64:
		expectation.float64Eq(actual, expected)
	case bool:
		expectation.boolEq(actual, expected)
	case string:
		expectation.stringEq(actual, expected)
	case complex64:
		expectation.complex64Eq(actual, expected)
	case complex128:
		expectation.complex128Eq(actual, expected)
	case time.Time:
		expectation.timeEq(actual, expected)
	case any:
		expectation.deepEq(&actual, &expected)
	default:
		relPath, line := getTestInfo(1)
		msg := expectation.FailMsg("!!ASSERTION ERROR!!: Type [%T] is not supported with `ToBe` method.")
		expectation.processFailure(relPath, line, msg, nil)
	}
}

func (expectation *Expectation[A]) intEq(
	actual int,
	expected A,
) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *int))
	assertEq[int](
		actual,
		any(expected).(int),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation *Expectation[A]) int8Eq(actual int8, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *int8))
	assertEq[int8](
		actual,
		any(expected).(int8),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation *Expectation[A]) int16Eq(actual int16, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *int16))
	assertEq[int16](
		actual,
		any(expected).(int16),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) int32Eq(actual int32, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *int32))
	assertEq[int32](
		actual,
		any(expected).(int32),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) int64Eq(actual int64, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *int64))
	assertEq[int64](
		actual,
		any(expected).(int64),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) uintEq(actual uint, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *uint))
	assertEq[uint](
		actual,
		any(expected).(uint),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) uint8Eq(actual uint8, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *uint8))
	assertEq[uint8](
		actual,
		any(expected).(uint8),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) uint16Eq(actual uint16, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *uint16))
	assertEq[uint16](
		actual,
		any(expected).(uint16),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) uint32Eq(actual uint32, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *uint32))
	assertEq[uint32](
		actual,
		any(expected).(uint32),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) uint64Eq(actual uint64, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *uint64))
	assertEq[uint64](
		actual,
		any(expected).(uint64),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) uintptrEq(actual uintptr, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *uintptr))
	assertEq[uintptr](
		actual,
		any(expected).(uintptr),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) float32Eq(actual float32, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *float32))
	assertEq[float32](
		actual,
		any(expected).(float32),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) float64Eq(actual float64, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *float64))
	assertEq[float64](
		actual,
		any(expected).(float64),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) boolEq(actual bool, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *bool))
	assertEq[bool](
		actual,
		any(expected).(bool),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) complex64Eq(actual complex64, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *complex64))
	assertEq[complex64](
		actual,
		any(expected).(complex64),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}
func (expectation *Expectation[A]) complex128Eq(actual complex128, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *complex128))
	assertEq[complex128](
		actual,
		any(expected).(complex128),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation *Expectation[A]) stringEq(actual string, expected A) {
	expectation.test.testingT.Helper()

	processFailure := any(expectation.processFailure).(func(string, int, string, *string))
	assertEq[string](
		actual,
		any(expected).(string),
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

// REFACTOR: too many arguments. there should be a better way
func assertEq[T comparable](
	actual T,
	expected T,
	reverse bool,
	failMsg func(string) string,
	failFunc func(string, int, string, *T),
	passFunc func(),
	Helper func(),
) {
	Helper()

	relPath, line := getTestInfo(3)

	if reverse {
		if actual != expected {
			passFunc()
			return
		}
		msg := failMsg("actual:[%#v] IS expected:[%#v]")
		failFunc(relPath, line, msg, &expected)
		return
	}
	if actual == expected {
		passFunc()
		return
	}
	msg := failMsg("actual:[%#v] is NOT expected:[%#v]")
	failFunc(relPath, line, msg, &expected)
}

// REFACTOR:
func (expectation *Expectation[A]) timeEq(actual time.Time, expected A) {
	expectation.test.testingT.Helper()

	relPath, line := getTestInfo(2)

	if expectation.reverseExpectation {
		if !actual.Equal(any(expected).(time.Time)) {
			expectation.processPassed()
			return
		}
		failMsg := expectation.FailMsg("actual:[%#v] IS expected:[%#v]")
		expectation.processFailure(relPath, line, failMsg, &expected)
		return
	}

	if actual.Equal(any(expected).(time.Time)) {
		expectation.processPassed()
		return
	}
	failMsg := expectation.FailMsg("actual:[%#v] is NOT expected:[%#v]")
	expectation.processFailure(relPath, line, failMsg, &expected)
}

// REFACTOR:
func (expectation *Expectation[A]) deepEq(actual *any, expected *A) {
	expectation.test.testingT.Helper()

	relPath, line := getTestInfo(2)

	if expectation.reverseExpectation {
		if !reflect.DeepEqual(*actual, *expected) {
			expectation.processPassed()
			return
		}
		failMsg := expectation.FailMsg("actual:[%#v] IS expected:[%#v]")
		expectation.processFailure(relPath, line, failMsg, expected)
		return
	}

	if reflect.DeepEqual(*actual, *expected) {
		expectation.processPassed()
		return
	}
	failMsg := expectation.FailMsg("actual:[%#v] is NOT expected:[%#v]")
	expectation.processFailure(relPath, line, failMsg, expected)
}
