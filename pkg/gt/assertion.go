// Main assertion methods
package gt

import (
	"fmt"
	"regexp"
	"strings"
	"time"
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
func LogWhenFail[A any](test *Test, failMsg string) *Expectation[A] {
	return &Expectation[A]{
		test:               test,
		failMsg:            failMsg,
		reverseExpectation: false,
	}
}

// Use this ONLY AFTER you call `LogWhenFail()` or any other Expectation constructors
// if exist
// `LogWhenFail()`を呼び出した後にのみ、この関数を呼び出してください。
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

	expectation.test.subtotal++
	if expectation.reverseExpectation {
		if expectation.actual != nil {
			expectation.processPassed()
			return
		}
		failMsg := expectation.FailMsg("value IS nil: [%#v]")
		expectation.processFailure(failMsg, expectation.actual, nil)
		return
	}
	if expectation.actual == nil {
		expectation.processPassed()
		return
	}
	failMsg := expectation.FailMsg("value is NOT nil: [%#v]")
	expectation.processFailure(failMsg, *expectation.actual, nil)

}

// This function asserts whether the interface is nil or not.
// It's primarily used to determine if the error is nil or not.
func (expectation *Expectation[A]) ToBeNilInterface() {
	expectation.test.testingT.Helper()

	// interface is passed as a pointer, so we need to dereference it
	convertedActual := any(*expectation.actual)

	expectation.test.subtotal++
	if expectation.reverseExpectation {
		if convertedActual != nil {
			expectation.processPassed()
			return
		}
		failMsg := expectation.FailMsg("interface IS nil")
		expectation.processFailure(failMsg, nil, nil)
		return
	}
	if convertedActual == nil {
		expectation.processPassed()
		return
	}
	failMsg := expectation.FailMsg("the interface is NOT nil: [%#v]")
	expectation.processFailure(failMsg, convertedActual, nil)

}

// Assertion for pointer values.
// ポインタ型の値を比較します。ポインタが同じ場合はアサートがpassします。
func (expectation *Expectation[A]) ToBeSamePointerAs(expected *A) {
	expectation.test.testingT.Helper()

	asserting(
		expectation,
		expectation.actual,
		expected,
		"[%#v] is NOT the same. Expected: [%#v]",
		"[%#v] IS the same. Expected: [%#v]",
		equalityAssertion,
	)
}

// 文字列が正規表現にマッチする場合はアサートがpassします。
func (expectation *Expectation[A]) ToMatchRegex(expected string) {
	expectation.test.testingT.Helper()

	actualString := fmt.Sprintf("%v", *expectation.actual)
	asserting(
		expectation,
		actualString,
		expected,
		"actual:[%#v] does NOT match with regex expected:[%#v]",
		"actual:[%#v] DOES match with regex expected:[%#v]",
		regexAssertion,
	)
}

// 文字列が含まれている場合はアサートがpassします。
func (expectation *Expectation[A]) ToContainString(expected string) {
	expectation.test.testingT.Helper()

	actualString := fmt.Sprintf("%v", *expectation.actual)
	asserting(
		expectation,
		actualString,
		expected,
		"actual:[%#v] does NOT contain expected:[%#v]",
		"actual:[%#v] DOES contain expected:[%#v]",
		strings.Contains,
	)
}

// TODO: implement:
// assert json value the same

// ----------------------------------------------------

func asserting[A any, T any](
	expectation *Expectation[A],
	convertedActual T,
	convertedExpected T,
	failMessage string,
	reverseFailMessage string,
	assertion func(T, T) bool,
) {
	expectation.test.testingT.Helper()

	expectation.test.subtotal++
	if expectation.reverseExpectation {
		if !assertion(convertedActual, convertedExpected) {
			expectation.processPassed()
			return
		}
		failMsg := expectation.FailMsg(reverseFailMessage)
		expectation.processFailure(failMsg, convertedActual, convertedExpected)
		return
	}
	if assertion(convertedActual, convertedExpected) {
		expectation.processPassed()
		return
	}
	failMsg := expectation.FailMsg(failMessage)
	expectation.processFailure(failMsg, convertedActual, convertedExpected)
}

func assertingComparable[A any, T any](
	expectation *Expectation[A],
	convertedActual T,
	failMessage string,
	reverseFailMessage string,
	assertion func(T) bool,
) {
	expectation.test.testingT.Helper()

	expectation.test.subtotal++
	if expectation.reverseExpectation {
		if !assertion(convertedActual) {
			expectation.processPassed()
			return
		}
		failMsg := expectation.FailMsg(reverseFailMessage)
		expectation.processFailure(failMsg, convertedActual, nil)
		return
	}
	if assertion(convertedActual) {
		expectation.processPassed()
		return
	}
	failMsg := expectation.FailMsg(failMessage)
	expectation.processFailure(failMsg, convertedActual, nil)
}

func (expectation *Expectation[A]) FailMsg(msg string) string {

	if expectation.failMsg != "" {
		return expectation.failMsg
	}
	return msg
}

func regexAssertion(actual string, expected string) bool {
	matched, _ := regexp.MatchString(expected, actual)
	return matched
}

func equalityAssertion[T comparable](actual T, expected T) bool {
	return actual == expected
}

func timeAssertion(actual time.Time, expected time.Time) bool {
	return actual.Equal(any(expected).(time.Time))
}

func (expectation Expectation[A]) processFailure(
	errorMsg string,
	actual any,
	expected any,
) {
	expectation.test.testingT.Helper()

	msg := expectation.formatFailMessage(errorMsg, actual, expected)
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
	errorMsg string,
	actual any,
	expected any,
) string {
	expectation.test.testingT.Helper()

	if actual != nil {
		if expected != nil {
			return fmt.Sprintf(errorMsg, actual, expected)
		}
		return fmt.Sprintf(errorMsg, actual)
	}
	return errorMsg
}

func (expectation *Expectation[A]) resetNot() {
	expectation.test.testingT.Helper()

	expectation.reverseExpectation = false
}
