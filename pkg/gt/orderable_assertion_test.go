package gt

import "testing"

func TestOrderableAssertions(testingT *testing.T) {
	t := CreateTest(testingT)

	t.Describe("ToBe_", func() {
		t.It("should pass when the value is in the comparator range", func() {
			var intVal1 int = 10
			Expect(t, &intVal1).ToBe_(GreaterThan, 1)

			var int8Val1 int8 = 10
			Expect(t, &int8Val1).ToBe_(GreaterThan, 1)

			var int16Val1 int16 = 10
			Expect(t, &int16Val1).ToBe_(GreaterThan, 1)

			var float32Val1 float32 = 10.2
			Expect(t, &float32Val1).ToBe_(GreaterThan, 10.1)

			valBt := 1
			Expect(t, &valBt).ToBe_(Between(1), 11)
		})

		t.It("should pass when the value is NOT in the comparator range", func() {

			var intVal int = 10
			Expect(t, &intVal).Not().ToBe_(GreaterThan, 11)
		})
	})

	t2 := CreateTest(testingT)
	t2.Describe("Between", func() {
		t2.It("should pass when the value is in between", func() {

			valBt1 := 1
			Expect(t2, &valBt1).ToBe_(Between(1), 3)

			valBt2 := 2
			Expect(t2, &valBt2).ToBe_(Between(1), 3)

			valBt3 := 3
			Expect(t2, &valBt3).ToBe_(Between(1), 3)
		})

		t2.It("should pass when value is NOT in between", func() {
			valBt0 := 0
			Expect(t2, &valBt0).Not().ToBe_(Between(1), 3)

			valBt4 := 4
			Expect(t2, &valBt4).Not().ToBe_(Between(1), 3)
		})
	})
}
