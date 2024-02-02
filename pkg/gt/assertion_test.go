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

	t.Describe("ToBe", func() {
		t.It("should pass when two values are equal", func() {
			// not all primitive types are tested,
			// but it's enough to test the functionality
			var intVal1 int = 1
			var intVal2 int = 1
			Expect(t, &intVal1).ToBe(intVal2)

			var strVal1 string = "hoge"
			var strVal2 string = "hoge"
			Expect(t, &strVal1).ToBe(strVal2)

			var boolVal1 bool = true
			var boolVal2 bool = true
			Expect(t, &boolVal1).ToBe(boolVal2)

			var floatVal1 float64 = 1.1
			var floatVal2 float64 = 1.1
			Expect(t, &floatVal1).ToBe(floatVal2)

			var complexVal1 complex128 = 1 + 1i
			var complexVal2 complex128 = 1 + 1i
			Expect(t, &complexVal1).ToBe(complexVal2)
		})

		t.It("should pass when two values are NOT equal", func() {
			var intVal1 int = 1
			var intVal2 int = 2
			Expect(t, &intVal1).Not().ToBe(intVal2)
		})
	})

	t.Describe("ToBeSamePointerAs", func() {
		t.It("should pass when two pointers are equal", func() {
			a := 1
			p := &a
			Expect(t, &a).ToBeSamePointerAs(p)
		})

		t.It("should pass when two pointers are NOT equal even if values are equal", func() {
			a := 1
			b := 1
			Expect(t, &a).Not().ToBeSamePointerAs(&b)
		})
	})

	t.Describe("ToDeepEqual", func() {
		t.It("should pass when two objects are equal", func() {
			a := Person{Name: "hoge", Age: 1}
			b := Person{Name: "hoge", Age: 1}

			Expect(t, &a).ToDeepEqual(b)
		})

		t.It("should pass when two objects are NOT equal", func() {
			a := Person{Name: "hoge", Age: 1}
			b := Person{Name: "foo", Age: 1}

			Expect(t, &a).Not().ToDeepEqual(b)
		})
	})

}
