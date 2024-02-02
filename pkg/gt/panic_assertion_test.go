package gt

import (
	"testing"
)

func TestPanicAssertions(testingT *testing.T) {
	t := CreateTest(testingT)

	t.Describe("ExpectPanic", func() {
		t.It("should pass when panic happens", func() {
			ExpectPanic(t).ToHappen(func() {
				panic("test")
			})
		})

		t.It("should pass when panic does not happen", func() {
			ExpectPanic(t).Not().ToHappen(func() {
				// do nothing
			})
		})
	})
}
