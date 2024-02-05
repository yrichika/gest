// Main assertion methods
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

// アサートする際に、`*Test`を第1引数に、アサート対象の値のポインタを第2引数に渡してください。
func Expect[A any](test *Test, actual *A) *Expectation[A] {
	// TODO: Testを常に引数に渡さないようにしたい。
	// ただしメソッドに型パラメータをもたせられないため、方法がわからない。
	return &Expectation[A]{
		actual:             actual,
		test:               test,
		failMsg:            "",
		reverseExpectation: false,
	}
}

// アサートがFailした際に出力する文字列を変更したい場合に、第2引数にその文字列を渡してください。
// この関数を呼び出した後に、`Expect(*A)`を呼び出してください。この後に呼び出すExpectは
// 通常の`Expect(*Test, *A)`とは違い、`*Test`を第1引数に取りません。
func PrintWhenFail[A any](test *Test, failMsg string) *Expectation[A] {
	return &Expectation[A]{
		test:               test,
		failMsg:            failMsg,
		reverseExpectation: false,
	}
}

// Use this ONLY AFTER you call `PrintWhenFail()` or any other Expectation constructors
// if exist
// `PrintWhenFail()`を呼び出した後にのみ、この関数を呼び出してください。
// それ以外の場合では意味がない、もしくは正しく動作しない可能性があります。
func (expectation *Expectation[A]) Expect(actual *A) *Expectation[A] {
	expectation.test.testingT.Helper()

	expectation.actual = actual
	return expectation
}

// アサートの結果が逆転します。`ToBe()`などのアサートするメソッドの直前で
// `Not()`を呼び出してください。
// 例: `Expect(test, &actual).Not().ToBe(expected)`
func (expectation *Expectation[A]) Not() *Expectation[A] {
	expectation.test.testingT.Helper()

	expectation.reverseExpectation = true
	return expectation
}

// Assertion for nil values.
// アサート対象の値がnilかどうかを確認します。
// 対象の値がnilの場合はアサートがpassします。
func (expectation *Expectation[A]) ToBeNil() {
	expectation.test.testingT.Helper()

	relPath, line := getTestInfo(1)

	expectation.test.subtotal++
	if expectation.reverseExpectation {
		if expectation.actual != nil {
			expectation.processPassed()
			return
		}
		failMsg := expectation.FailMsg("Value IS nil")
		expectation.processFailure(relPath, line, failMsg, nil)
		return
	}

	if expectation.actual == nil {
		expectation.processPassed()
		return
	}
	failMsg := expectation.FailMsg("[%#v] is NOT nil")
	expectation.processFailure(relPath, line, failMsg, nil)
}

// Assertion for primitive values.
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
	default:
		relPath, line := getTestInfo(1)
		msg := expectation.FailMsg("!!ASSERTION ERROR!!: Type [%T] is not supported with `ToBe` method. `ToBe` is intended only for primitive types. Please use `ToDeepEqual` method if it's a struct type.")
		expectation.processFailure(relPath, line, msg, nil)
	}
}

// two structs equality
// 2つの構造体の等価性を確認します。等価性は、構造体のフィールドの値が全て等しいことを意味します。
// 2つの構造体が等しい場合はアサートがpassします。
// 等価性の確認に、内部では`reflect.DeepEqual()`が使用されます。
func (expectation *Expectation[A]) ToDeepEqual(expected A) {
	expectation.test.testingT.Helper()

	relPath, line := getTestInfo(1)

	expectation.test.subtotal++
	if expectation.reverseExpectation {
		if !reflect.DeepEqual(*expectation.actual, expected) {
			expectation.processPassed()
			return
		}
		failMsg := expectation.FailMsg("actual:[%#v] IS expected:[%#v]")
		expectation.processFailure(relPath, line, failMsg, &expected)
		return
	}

	if reflect.DeepEqual(*expectation.actual, expected) {
		expectation.processPassed()
		return
	}
	failMsg := expectation.FailMsg("actual:[%#v] is NOT expected:[%#v]")
	expectation.processFailure(relPath, line, failMsg, &expected)
}

// Assertion for pointer values.
// ポインタ型の値を比較します。ポインタが同じ場合はアサートがpassします。
func (expectation *Expectation[A]) ToBeSamePointerAs(expected *A) {
	expectation.test.testingT.Helper()

	relPath, line := getTestInfo(1)

	expectation.test.subtotal++
	if expectation.reverseExpectation {
		if expectation.actual != expected {
			expectation.processPassed()
			return
		}
		failMsg := expectation.FailMsg("Pointer to [%#v] IS the same")
		expectation.processFailure(relPath, line, failMsg, nil)
		return
	}
	if expectation.actual == expected {
		expectation.processPassed()
		return
	}
	failMsg := expectation.FailMsg("Pointer to [%#v] is NOT the same")
	expectation.processFailure(relPath, line, failMsg, nil)

}

// TODO: implement:
// func ToContainString()
// func ToMatchRegex()
// assert json value the same

func (expectation *Expectation[A]) FailMsg(msg string) string {

	if expectation.failMsg != "" {
		return expectation.failMsg
	}
	return msg
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
	}
	if actual == expected {
		passFunc()
		return
	}
	msg := failMsg("actual:[%#v] is NOT expected:[%#v]")
	failFunc(relPath, line, msg, &expected)
	return

}

func getTestInfo(skip int) (string, int) {
	skip++ // add one for this function
	_, file, line, _ := runtime.Caller(skip)
	relPath := extractRelPath(file)
	return relPath, line
}

func (expectation Expectation[A]) processFailure(
	relPath string,
	line int,
	errorMsg string,
	expected *A,
) {
	expectation.test.testingT.Helper()

	msg := expectation.formatFailMessage(relPath, line, errorMsg, expected)
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

func (expectation *Expectation[A]) formatFailMessage(
	relPath string,
	line int,
	errorMsg string,
	expected *A,
) string {
	expectation.test.testingT.Helper()

	if expectation.actual != nil {
		if expected != nil {
			return fmt.Sprintf(
				"Failed at [%s]:line %d: %s",
				relPath, line, fmt.Sprintf(errorMsg, *expectation.actual, *expected),
			)
		}
		return fmt.Sprintf(
			"Failed at [%s]:line %d: %s",
			relPath, line, fmt.Sprintf(errorMsg, *expectation.actual),
		)

	}
	return fmt.Sprintf("Failed at [%s]:line %d: %s", relPath, line, fmt.Sprintf(errorMsg))
}

func (expectation *Expectation[A]) resetNot() {
	expectation.test.testingT.Helper()

	expectation.reverseExpectation = false
}
