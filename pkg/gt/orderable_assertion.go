package gt

import (
	"time"
)

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
		float64 |
		time.Duration
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
