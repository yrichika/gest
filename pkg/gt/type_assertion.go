package gt

// OfBool, OfInt, OfString などを使い、型のアサーションを行います。
// e.g. `Expect(t, &val).ToBeType(OfInt)`
func (expectation *Expectation[A]) ToBeType(typeComparator func(*A) bool) {
	expectation.test.testingT.Helper()

	relPath, line := getTestInfo(1)

	expectation.test.subtotal++
	if expectation.reverseExpectation {
		if !typeComparator(expectation.actual) {
			expectation.processPassed()
			return
		}
		failMsg := expectation.FailMsg("actual:[%T] IS expected type")
		expectation.processFailure(relPath, line, failMsg, nil)
		return
	}
	if typeComparator(expectation.actual) {
		expectation.processPassed()
		return
	}
	failMsg := expectation.FailMsg("actual:[%T] is NOT expected type")
	expectation.processFailure(relPath, line, failMsg, nil)
}
