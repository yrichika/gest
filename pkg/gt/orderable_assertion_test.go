package gt

import (
	"testing"
	"time"
)

func TestOrderableAssertions(testingT *testing.T) {
	t := CreateTest(testingT)

	t.Describe("ToBe_ with number comparators", func() {
		t.It("should pass when the value is in the comparator range", func() {
			var intVal1 int = 10
			Expect(t, &intVal1).ToBe_(GreaterThan, 1)

			var int8Val1 int8 = 10
			Expect(t, &int8Val1).ToBe_(GreaterThan, 1)

			var int16Val1 int16 = 10
			Expect(t, &int16Val1).ToBe_(GreaterThan, 1)

			var float32Val1 float32 = 10.2
			Expect(t, &float32Val1).ToBe_(GreaterThan, 10.1)

			var duration time.Duration = 2 * time.Second
			Expect(t, &duration).ToBe_(GreaterThan, 1*time.Second)

			valBt := 1
			Expect(t, &valBt).ToBe_(Between(1), 11)

			var durationBt time.Duration = 2 * time.Second
			Expect(t, &durationBt).ToBe_(Between(1*time.Second), 3*time.Second)
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

	t3 := CreateTest(testingT)
	t3.Describe("ToBe_ with time comparators", func() {
		t3.It("should pass when value is in the comparator range", func() {
			now := time.Now()

			past := time.Now().Add(-1 * time.Minute)
			Expect(t3, &past).ToBe_(Before, now)

			future := time.Now().Add(1 * time.Minute)
			Expect(t3, &future).ToBe_(After, now)

			// Equality assertions
			Expect(t3, &now).ToBe_(BeforeOrEq, now)
			Expect(t3, &now).ToBe_(AfterOrEq, now)
		})

		t3.It("should pass when value is NOT in the comparator range", func() {
			past := time.Now().Add(-1 * time.Minute)
			Expect(t3, &past).Not().ToBe_(After, time.Now())
		})

		t3.It("should pass when time value is in between", func() {
			past := time.Now().Add(-1 * time.Minute)
			future := time.Now().Add(1 * time.Minute)
			now := time.Now()
			Expect(t3, &now).ToBe_(TimeBetween(past), future)
		})

	})
}
