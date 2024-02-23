package gt

import "testing"

func TestTypeAssertions(testingT *testing.T) {

	t := CreateTest(testingT)
	t.Describe("ToBe_ with type comparators", func() {
		t.It("should pass when the value is of the expected type", func() {
			valBool := true
			Expect(t, &valBool).ToBeType(OfBool)

			valInt := 1
			Expect(t, &valInt).ToBeType(OfInt)

			valStr := "abc"
			Expect(t, &valStr).ToBeType(OfString)

		})
	})
}
