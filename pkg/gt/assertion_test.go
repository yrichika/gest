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

	t2 := CreateTest(testingT)
	t2.Describe("ToBe", func() {
		t2.It("should pass when two values are equal", func() {
			// not all primitive types are tested,
			// but it's enough to test the functionality
			var intVal1 int = 1
			var intVal2 int = 1
			Expect(t2, &intVal1).ToBe(intVal2)

			var strVal1 string = "hoge"
			var strVal2 string = "hoge"
			Expect(t2, &strVal1).ToBe(strVal2)

			var boolVal1 bool = true
			var boolVal2 bool = true
			Expect(t2, &boolVal1).ToBe(boolVal2)

			var floatVal1 float64 = 1.1
			var floatVal2 float64 = 1.1
			Expect(t2, &floatVal1).ToBe(floatVal2)

			var complexVal1 complex128 = 1 + 1i
			var complexVal2 complex128 = 1 + 1i
			Expect(t2, &complexVal1).ToBe(complexVal2)
		})

		t2.It("should pass when two values are NOT equal", func() {
			var intVal1 int = 1
			var intVal2 int = 2
			Expect(t2, &intVal1).Not().ToBe(intVal2)
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

	t4 := CreateTest(testingT)
	t4.Describe("ToDeepEqual", func() {
		t4.It("should pass when two objects are equal", func() {
			a := Person{Name: "hoge", Age: 1}
			b := Person{Name: "hoge", Age: 1}

			Expect(t4, &a).ToDeepEqual(b)
		})

		t4.It("should pass when two objects are NOT equal", func() {
			a := Person{Name: "hoge", Age: 1}
			b := Person{Name: "foo", Age: 1}

			Expect(t4, &a).Not().ToDeepEqual(b)
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
