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

	asserting(
		expectation,
		*expectation.actual,
		expected,
		"compared actual:[%#v] and expected:[%#v]",
		"compared actual:[%#v] and expected:[%#v]",
		1,
		comparator,
	)
}
