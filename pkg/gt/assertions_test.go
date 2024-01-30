package gt

import (
	"testing"
)

type MockObject struct {
	Name string
	Age  int
}

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

		t.It("should pass when int object equals", func() {
			var a int = 1
			var b int = 1
			Expect(t, &a).ToBe(b)

			var c string = "hoge"
			var d string = "hoge"
			Expect(t, &c).ToBe(d)
		})

		t.It("compares pointers", func() {
			var a int = 1
			var b int = 1
			var aPointer *int = &a
			// var bPointer *int = &b
			Expect(t, aPointer).ToBeSamePointerAs(&b)
		})

		t.It("should pass when two objects are equal", func() {
			a := MockObject{Name: "hoge", Age: 1}
			b := MockObject{Name: "hoge", Age: 1}

			Expect(t, &a).ToDeepEqual(b)
		})

		t.It("WhenFailPrint", func() {
			var a int = 1
			var b int = 2
			WhenFailPrint[int](t, "show this message when fail").Expect(&a).ToBe(b)
		})

		t.It("should pass when panic happens", func() {
			ExpectPanic(t).ToHappen(func() {
				panic("panic")
			})
		})
	})
}
