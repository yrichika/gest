package gt

import (
	"fmt"
)

type PanicExpectation struct {
	test               *Test
	reverseExpectation bool
	failMsg            string
}

func ExpectPanic(test *Test) *PanicExpectation {
	return &PanicExpectation{
		test:               test,
		reverseExpectation: false,
		failMsg:            "",
	}
}

// errorやpanicが起きたことをアサート
func (p *PanicExpectation) ToHappen(panickyFunc func()) {
	p.test.testingT.Helper()
	p.test.subtotal++

	relPath, line := getTestInfo(1)

	defer func() {
		p.test.testingT.Helper() // DON'T FORGET THIS

		err := recover()
		if p.reverseExpectation {
			if err == nil {
				p.processPassed()
				return
			}
			msg := p.FailMsg("Expected panic DID happen")
			p.processFailure(relPath, line, msg)
			return
		}
		if err != nil {
			p.processPassed()
			return
		}
		msg := p.FailMsg("Expected panic did NOT happen")
		p.processFailure(relPath, line, msg)
	}()

	panickyFunc()
}

func (p *PanicExpectation) Not() *PanicExpectation {
	p.test.testingT.Helper()

	p.reverseExpectation = true
	return p
}

func (p *PanicExpectation) FailMsg(msg string) string {

	if p.failMsg != "" {
		return p.failMsg
	}
	return msg
}

func (p *PanicExpectation) processFailure(relPath string, line int, errorMsg string) {
	p.test.testingT.Helper()

	msg := p.failMessage(relPath, line, errorMsg)
	p.test.testingT.Errorf(RedMsg(msg))
	p.markAsFailed()
	p.resetNot()
}

func (p *PanicExpectation) processPassed() {
	p.test.testingT.Helper()

	p.test.passed++
	p.resetNot()
}

func (p *PanicExpectation) markAsFailed() {
	p.test.testingT.Helper()

	p.test.isThisTestFailed = true
	p.test.isAnyTestFailed = true
}

func (p *PanicExpectation) failMessage(relPath string, line int, errorMsg string) string {
	p.test.testingT.Helper()

	return fmt.Sprintf("Panic: [%s]:line %d: %s", relPath, line, fmt.Sprintf(errorMsg))
}

func (p *PanicExpectation) resetNot() {
	p.test.testingT.Helper()

	p.reverseExpectation = false
}
