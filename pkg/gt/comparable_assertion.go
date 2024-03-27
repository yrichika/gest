package gt

import (
	"reflect"
	"time"
)

const equalityFailedMsg = "actual:[%#v] is NOT expected:[%#v]"
const equalityReverseFailedMsg = "actual:[%#v] IS expected:[%#v]"

// Assertion for primitive values, time.Time and custom struct types.
// Types must be the same.
// プリミティブ型の値を比較します。値が同じ場合はアサートがpassします。
// アサート対象の値と、想定している値は同じ型である必要があります。
func (expectation *Expectation[A]) ToBe(expected A) {
	expectation.test.testingT.Helper()

	switch actual := any(*expectation.actual).(type) {
	case
		int,
		int8,
		int16,
		int32,
		int64,
		uint,
		uint8,
		uint16,
		uint32,
		uint64,
		uintptr,
		float32,
		float64,
		bool,
		string,
		complex64,
		complex128:
		expectation.comparableEq(actual, expected)
	case time.Duration:
		expectation.durationEq(actual, expected)
	case time.Time:
		expectation.timeEq(actual, expected)
	case any:
		expectation.deepEq(&actual, &expected)
	default:
		msg := expectation.FailMsg("!!ASSERTION ERROR!!: Type [%T] is not supported with `ToBe` method.")
		expectation.processFailure(msg, actual, nil)
	}
}

func (expectation *Expectation[A]) comparableEq(actual any, expected any) {
	expectation.test.testingT.Helper()

	asserting(
		expectation,
		actual,
		expected,
		equalityFailedMsg,
		equalityReverseFailedMsg,
		equalityAssertion,
	)
}

func (expectation *Expectation[A]) durationEq(actual time.Duration, expected A) {
	expectation.test.testingT.Helper()

	asserting(
		expectation,
		actual,
		any(expected).(time.Duration),
		"actual:[%v] is NOT expected:[%v]",
		"actual:[%v] IS expected:[%v]",
		equalityAssertion,
	)
}

func (expectation *Expectation[A]) timeEq(actual time.Time, expected A) {
	expectation.test.testingT.Helper()

	asserting(
		expectation,
		actual,
		any(expected).(time.Time),
		equalityFailedMsg,
		equalityReverseFailedMsg,
		timeAssertion,
	)
}

func (expectation *Expectation[A]) deepEq(actual *any, expected *A) {
	expectation.test.testingT.Helper()

	converted := any(*expected)
	asserting(
		expectation,
		*actual,
		converted,
		equalityFailedMsg,
		equalityReverseFailedMsg,
		reflect.DeepEqual,
	)
}
