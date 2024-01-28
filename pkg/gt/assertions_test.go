package gt

import (
	"testing"
)

func TestAssertions(testingT *testing.T) {
	t := CreateTest(testingT)

	t.Describe("assertions", func() {
		// TODO: test assertions.go functions
	})

	t.Describe("toBeNil", func() {
		t.It("should pass when nil", func() {
			// var a int = 1
			// var b string = "b"
			var nilValue *int = nil // or &a to fail
			// var nilValue *string = &b // or &b to fail
			Expect(t, nilValue).ToBeNil()
		})

		t.It("should fail when nil", func() {
			var nilValue *int = nil
			Expect(t, nilValue).Not().ToBeNil()
		})
	})
}
