package gt

import "time"

type orderable interface {
	int |
		int8 |
		int16 |
		int32 |
		int64 |
		uint |
		uint8 |
		uint16 |
		uint32 |
		uint64 |
		uintptr |
		float32 |
		float64
}

// GreaterThan, LessThanなどを使い、値の比較のアサーションを行います。
// e.g. `Expect(t, &val).ToBe_(GreaterThan, 1)`
func (expectation *Expectation[A]) ToBe_(comparator func(A, A) bool, expected A) {
	expectation.test.testingT.Helper()

	relPath, line := getTestInfo(1)

	expectation.test.subtotal++
	if expectation.reverseExpectation {
		if !comparator(*expectation.actual, expected) {
			expectation.processPassed()
			return
		}
		failMsg := expectation.FailMsg("compared actual:[%#v] and expected:[%#v]")
		expectedForFailMsg := any(expected).(A)
		expectation.processFailure(relPath, line, failMsg, &expectedForFailMsg)
		return
	}
	if comparator(*expectation.actual, expected) {
		expectation.processPassed()
		return
	}
	failMsg := expectation.FailMsg("compared actual:[%#v] and expected:[%#v]")
	expectedForFailMsg := any(expected).(A)
	expectation.processFailure(relPath, line, failMsg, &expectedForFailMsg)
}

// this works too, but might be confusing. Not applying for now.
// func EqTo[T comparable](actual T, expected T) bool {
// 	return actual == expected
// }

func GreaterThan[T orderable](actual T, expected T) bool {
	return actual > expected
}

func GreaterThanOrEq[T orderable](actual T, expected T) bool {
	return actual >= expected
}

func LessThan[T orderable](actual T, expected T) bool {
	return actual < expected
}

func LessThanOrEq[T orderable](actual T, expected T) bool {
	return actual <= expected
}

// minに指定した数値と、expectedに指定した数値は範囲に含まれます
func Between[T orderable](min T) func(T, T) bool {
	return func(actual, max T) bool {
		return GreaterThanOrEq(actual, min) && LessThanOrEq(actual, max)
	}
}

func After(actual time.Time, expected time.Time) bool {
	return actual.After(expected)
}

func AfterOrEq(actual time.Time, expected time.Time) bool {
	return actual.After(expected) || actual.Equal(expected)
}

func Before(actual time.Time, expected time.Time) bool {
	return actual.Before(expected)
}

func BeforeOrEq(actual time.Time, expected time.Time) bool {
	return actual.Before(expected) || actual.Equal(expected)
}

// minに指定した数値と、expectedに指定した数値は範囲に含まれます
func TimeBetween[T time.Time](from time.Time) func(time.Time, time.Time) bool {
	return func(actual, to time.Time) bool {
		return AfterOrEq(actual, from) && BeforeOrEq(actual, to)
	}
}
