package gt

import (
	"reflect"
	"time"
)

// actualが指定したexpectedのスライスの要素の中にあるかどうかを確認
func (expectation Expectation[A]) ToBeIn(expected []A) {
	expectation.test.testingT.Helper()

	switch actual := any(*expectation.actual).(type) {
	case int,
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
		expectation.comparableIn(actual, expected)
	case time.Duration:
		expectation.durationIn(actual, expected)
	case time.Time:
		expectation.timeIn(actual, expected)
	case any:
		expectation.deepIn(actual, expected)
	default:
		relPath, line := getTestInfo(1)
		msg := expectation.FailMsg("!!ASSERTION ERROR!!: Type [%T] is not supported with `ToBeIn` method.")
		expectation.processFailure(relPath, line, msg, nil)
	}
}

func (expectation *Expectation[A]) comparableIn(actual any, expected []A) {
	expectation.test.testingT.Helper()

	convertedExpected := make([]any, len(expected))
	for i, v := range expected {
		convertedExpected[i] = v
	}
	assertingSlice(
		expectation,
		actual,
		convertedExpected,
		"actual:[%#v] is NOT in expected:[%v]",
		"actual:[%#v] IS in expected:[%v]",
		2,
		IsInSlice,
	)
}

func (expectation *Expectation[A]) durationIn(actual time.Duration, expected []A) {
	expectation.test.testingT.Helper()

	convertedExpected := make([]time.Duration, len(expected))
	for i, v := range expected {
		convertedExpected[i] = any(v).(time.Duration)
	}

	assertingSlice(
		expectation,
		actual,
		convertedExpected,
		"actual:[%v] is NOT in expected:[%v]",
		"actual:[%v] IS in expected:[%v]",
		2,
		IsInSlice,
	)
}

func (expectation *Expectation[A]) timeIn(actual time.Time, expected []A) {
	expectation.test.testingT.Helper()

	convertedExpected := make([]time.Time, len(expected))
	for i, v := range expected {
		convertedExpected[i] = any(v).(time.Time)
	}
	assertingSlice(
		expectation,
		actual,
		convertedExpected,
		"actual:[%#v] is NOT in expected:[%#v]",
		"actual:[%#v] IS in expected:[%#v]",
		2,
		isTimeInSlice,
	)
}

func isTimeInSlice(actual time.Time, expected []time.Time) bool {
	eq := func(a, b time.Time) bool {
		return a.Equal(b)
	}
	return ContainsElement(actual, expected, eq)
}

func (expectation *Expectation[A]) deepIn(actual any, expected []A) {
	expectation.test.testingT.Helper()

	convertedExpected := make([]any, len(expected))
	for i, v := range expected {
		convertedExpected[i] = v
	}
	assertingSlice(
		expectation,
		actual,
		convertedExpected,
		"actual:[%#v] is NOT in expected:[%#v]",
		"actual:[%#v] IS in expected:[%#v]",
		2,
		isStructInSlice,
	)
}

func isStructInSlice(actual any, expected []any) bool {
	eq := func(a, b any) bool {
		return reflect.DeepEqual(a, b)
	}
	return ContainsElement(actual, expected, eq)
}

// assertingとほとんど同じだが、expectedの型がスライスになっているところが違い
func assertingSlice[A any, T any](
	expectation *Expectation[A],
	convertedActual T,
	convertedExpected []T,
	failMessage string,
	reverseFailMessage string,
	skip int,
	assertion func(T, []T) bool,
) {
	expectation.test.testingT.Helper()

	relPath, line := getTestInfo(skip + 1)

	expectation.test.subtotal++
	if expectation.reverseExpectation {
		if !assertion(convertedActual, convertedExpected) {
			expectation.processPassed()
			return
		}
		failMsg := expectation.FailMsg(reverseFailMessage)
		expectation.processFailure(relPath, line, failMsg, convertedExpected)
		return
	}
	if assertion(convertedActual, convertedExpected) {
		expectation.processPassed()
		return
	}
	failMsg := expectation.FailMsg(failMessage)
	expectation.processFailure(relPath, line, failMsg, convertedExpected)
}

// TODO: array/slice/mapが同じかどうか
