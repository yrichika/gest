package gt

import "testing"

func TestComparableAssertions(testingT *testing.T) {
	t := CreateTest(testingT)

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

}
