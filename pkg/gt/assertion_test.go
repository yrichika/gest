package gt

import (
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestExpectAssertions(testingT *testing.T) {
	t := CreateTest(testingT)

	t.Describe("ToBeNil", func() {
		t.It("should pass when value is nil", func() {
			var nilValue *int
			Expect(t, nilValue).ToBeNil()
		})

		t.It("should pass when value is NOT nil", func() {
			notNilValue := 1
			Expect(t, &notNilValue).Not().ToBeNil()
		})
	})

	t3 := CreateTest(testingT)
	t3.Describe("ToBeSamePointerAs", func() {
		t3.It("should pass when two pointers are equal", func() {
			a := 1
			p := &a
			Expect(t3, &a).ToBeSamePointerAs(p)
		})

		t3.It("should pass when two pointers are NOT equal even if values are equal", func() {
			a := 1
			b := 1
			Expect(t3, &a).Not().ToBeSamePointerAs(&b)
		})
	})

	t5 := CreateTest(testingT)
	t5.Describe("ToMatchRegex", func() {
		t5.It("should pass when string matches regex", func() {
			a := "foo"
			Expect(t5, &a).ToMatchRegex("^foo$")
		})

		t5.It("should pass when string doet NOT matches regex", func() {
			a := "foo"
			Expect(t5, &a).Not().ToMatchRegex("^bar$")
		})
	})

	t6 := CreateTest(testingT)
	t6.Describe("ToContainString", func() {
		t6.It("should pass when string contains substring", func() {
			a := "hello world"
			Expect(t6, &a).ToContainString("o w")
		})

		t6.It("should pass when string does NOT contain substring", func() {
			a := "hello world"
			Expect(t6, &a).Not().ToContainString("foo")
		})
	})

}
