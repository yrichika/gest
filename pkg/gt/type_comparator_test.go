package gt

import (
	"testing"
)

func TestTypeComparators(testingT *testing.T) {

	// Not all types are tested. Some of them should be enough for now.

	t := CreateTest(testingT)
	t.Describe("Type Comparators", func() {
		t.It("BoolType", func() {
			val := true
			r := OfBool(&val)
			Expect(t, &r).ToBe(true)
		})

		t.It("IntType", func() {
			val := 1
			r := OfInt(&val)
			Expect(t, &r).ToBe(true)
		})

		t.It("StringType", func() {
			val := "a"
			r := OfString(&val)
			Expect(t, &r).ToBe(true)
		})

		t.It("ArrayType", func() {
			val := [3]int{1, 2, 3}
			r := OfArray(&val)
			Expect(t, &r).ToBe(true)
		})

		t.It("SliceType", func() {
			val := []int{1, 2, 3}
			r := OfSlice(&val)
			Expect(t, &r).ToBe(true)
		})

		t.It("MapType", func() {
			val := map[string]int{"a": 1}
			r := OfMap(&val)
			Expect(t, &r).ToBe(true)
		})
	})
}
