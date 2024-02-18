package gt

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
func (expectation Expectation[A]) ToBe_(comparator func(A, A) bool, expected A) {
	expectation.test.testingT.Helper()

	expectation.test.subtotal++
	switch actual := any(*expectation.actual).(type) {
	case int:
		expectation.compareInt(actual, expected, comparator)
	case int8:
		expectation.compareInt8(actual, expected, comparator)
	case int16:
		expectation.compareInt16(actual, expected, comparator)
	case int32:
		expectation.compareInt32(actual, expected, comparator)
	case int64:
		expectation.compareInt64(actual, expected, comparator)
	case uint:
		expectation.compareUint(actual, expected, comparator)
	case uint8:
		expectation.compareUint8(actual, expected, comparator)
	case uint16:
		expectation.compareUint16(actual, expected, comparator)
	case uint32:
		expectation.compareUint32(actual, expected, comparator)
	case uint64:
		expectation.compareUint64(actual, expected, comparator)
	case uintptr:
		expectation.compareUintptr(actual, expected, comparator)
	case float32:
		expectation.compareFloat32(actual, expected, comparator)
	case float64:
		expectation.compareFloat64(actual, expected, comparator)
	default:
		relPath, line := getTestInfo(1)
		msg := expectation.FailMsg("!!ASSERTION ERROR!!: Type [%T] is not supported with `ToBe_` method. `ToBe_` is intended only for numeric types.")
		expectation.processFailure(relPath, line, msg, nil)
	}
}

func (expectation Expectation[A]) compareInt(actual int, expected A, comparator func(A, A) bool) {
	processFailure := any(expectation.processFailure).(func(string, int, string, *int))
	intComparator := any(comparator).(func(int, int) bool)
	assertOrder[int](
		actual,
		any(expected).(int),
		intComparator,
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation Expectation[A]) compareInt8(actual int8, expected A, comparator func(A, A) bool) {
	processFailure := any(expectation.processFailure).(func(string, int, string, *int8))
	typedComparator := any(comparator).(func(int8, int8) bool)
	assertOrder[int8](
		actual,
		any(expected).(int8),
		typedComparator,
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation Expectation[A]) compareInt16(actual int16, expected A, comparator func(A, A) bool) {
	processFailure := any(expectation.processFailure).(func(string, int, string, *int16))
	typedComparator := any(comparator).(func(int16, int16) bool)
	assertOrder[int16](
		actual,
		any(expected).(int16),
		typedComparator,
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation Expectation[A]) compareInt32(actual int32, expected A, comparator func(A, A) bool) {
	processFailure := any(expectation.processFailure).(func(string, int, string, *int32))
	typedComparator := any(comparator).(func(int32, int32) bool)
	assertOrder[int32](
		actual,
		any(expected).(int32),
		typedComparator,
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation Expectation[A]) compareInt64(actual int64, expected A, comparator func(A, A) bool) {
	processFailure := any(expectation.processFailure).(func(string, int, string, *int64))
	typedComparator := any(comparator).(func(int64, int64) bool)
	assertOrder[int64](
		actual,
		any(expected).(int64),
		typedComparator,
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation Expectation[A]) compareUint(actual uint, expected A, comparator func(A, A) bool) {
	processFailure := any(expectation.processFailure).(func(string, int, string, *uint))
	typedComparator := any(comparator).(func(uint, uint) bool)
	assertOrder[uint](
		actual,
		any(expected).(uint),
		typedComparator,
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation Expectation[A]) compareUint8(actual uint8, expected A, comparator func(A, A) bool) {
	processFailure := any(expectation.processFailure).(func(string, int, string, *uint8))
	typedComparator := any(comparator).(func(uint8, uint8) bool)
	assertOrder[uint8](
		actual,
		any(expected).(uint8),
		typedComparator,
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation Expectation[A]) compareUint16(actual uint16, expected A, comparator func(A, A) bool) {
	processFailure := any(expectation.processFailure).(func(string, int, string, *uint16))
	typedComparator := any(comparator).(func(uint16, uint16) bool)
	assertOrder[uint16](
		actual,
		any(expected).(uint16),
		typedComparator,
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation Expectation[A]) compareUint32(actual uint32, expected A, comparator func(A, A) bool) {
	processFailure := any(expectation.processFailure).(func(string, int, string, *uint32))
	typedComparator := any(comparator).(func(uint32, uint32) bool)
	assertOrder[uint32](
		actual,
		any(expected).(uint32),
		typedComparator,
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation Expectation[A]) compareUint64(actual uint64, expected A, comparator func(A, A) bool) {
	processFailure := any(expectation.processFailure).(func(string, int, string, *uint64))
	typedComparator := any(comparator).(func(uint64, uint64) bool)
	assertOrder[uint64](
		actual,
		any(expected).(uint64),
		typedComparator,
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation Expectation[A]) compareUintptr(actual uintptr, expected A, comparator func(A, A) bool) {
	processFailure := any(expectation.processFailure).(func(string, int, string, *uintptr))
	typedComparator := any(comparator).(func(uintptr, uintptr) bool)
	assertOrder[uintptr](
		actual,
		any(expected).(uintptr),
		typedComparator,
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation Expectation[A]) compareFloat32(actual float32, expected A, comparator func(A, A) bool) {
	processFailure := any(expectation.processFailure).(func(string, int, string, *float32))
	typedComparator := any(comparator).(func(float32, float32) bool)
	assertOrder[float32](
		actual,
		any(expected).(float32),
		typedComparator,
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

func (expectation Expectation[A]) compareFloat64(actual float64, expected A, comparator func(A, A) bool) {
	processFailure := any(expectation.processFailure).(func(string, int, string, *float64))
	typedComparator := any(comparator).(func(float64, float64) bool)
	assertOrder[float64](
		actual,
		any(expected).(float64),
		typedComparator,
		expectation.reverseExpectation,
		expectation.FailMsg,
		processFailure,
		expectation.processPassed,
		expectation.test.testingT.Helper,
	)
}

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

// REFACTOR: too many arguments. there should be a better way
func assertOrder[T orderable](
	actual T,
	expected T,
	comparator func(T, T) bool,
	reverse bool,
	failMsg func(string) string,
	failFunc func(string, int, string, *T),
	passFunc func(),
	Helper func(),
) {
	Helper()

	relPath, line := getTestInfo(3)

	if reverse {
		if !comparator(actual, expected) {
			passFunc()
			return
		}
		msg := failMsg("compared actual:[%#v] and expected:[%#v]")
		failFunc(relPath, line, msg, &expected)
		return
	}
	if comparator(actual, expected) {
		passFunc()
		return
	}
	msg := failMsg("compared actual:[%#v] and expected:[%#v]")
	failFunc(relPath, line, msg, &expected)
}
