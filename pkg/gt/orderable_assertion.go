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
func (expectation *Expectation[A]) ToBe_(comparator func(A) bool) {
	expectation.test.testingT.Helper()

	assertingComparable(
		expectation,
		*expectation.actual,
		"compared values: actual:[%#v]",
		"compared values: actual:[%#v]",
		comparator,
	)
}
