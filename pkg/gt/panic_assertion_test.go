package gt

import (
	"testing"
)

func TestPanicAssertion(testingT *testing.T) {
	t := CreateTest(testingT)

	t.Describe("ExpectPanic", func() {
		t.Todo("should pass when panic happens")
	})
}
