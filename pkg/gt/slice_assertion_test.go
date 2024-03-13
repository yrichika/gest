package gt

import (
	"testing"
	"time"
)

func TestSliceAssertions(testingT *testing.T) {

	t := CreateTest(testingT)
	t.Describe("ToBeIn", func() {
		t.It("should pass when value is in []int", func() {
			intSlice := []int{1, 2, 3, 4, 5}
			val := 3
			Expect(t, &val).ToBeIn(intSlice)
		})
		t.It("should pass when value is NOT in []int", func() {
			intSlice := []int{1, 2, 3, 4, 5}
			val := 6
			Expect(t, &val).Not().ToBeIn(intSlice)
		})

		t.It("should pass when value is in []bool", func() {
			boolSlice := []bool{false, false, false, true}
			val := true
			Expect(t, &val).ToBeIn(boolSlice)
		})

		t.It("should pass when value is NOT in []bool", func() {
			boolSlice := []bool{false, false, false}
			val := true
			Expect(t, &val).Not().ToBeIn(boolSlice)
		})

		t.It("should pass when value is in []string", func() {
			strSlice := []string{"foo", "bar", "baz"}
			val := "bar"
			Expect(t, &val).ToBeIn(strSlice)
		})

		t.It("should pass when value is NOT in []string", func() {
			strSlice := []string{"foo", "bar", "baz"}
			val := "hoge"
			Expect(t, &val).Not().ToBeIn(strSlice)
		})

		t.It("should pass when value is in []time.Duration", func() {
			durationSlice := []time.Duration{
				1 * time.Second,
				2 * time.Second,
				3 * time.Second,
			}
			val := 3 * time.Second
			Expect(t, &val).ToBeIn(durationSlice)
		})

		t.It("should pass when value is NOT in []time.Duration", func() {
			durationSlice := []time.Duration{
				1 * time.Second,
				2 * time.Second,
				3 * time.Second,
			}
			val := 4 * time.Second
			Expect(t, &val).Not().ToBeIn(durationSlice)
		})

		t.It("should pass when value is in []time.Time", func() {
			now := time.Now()
			timeSlice := []time.Time{
				now,
				now.Add(1 * time.Hour),
				now.Add(2 * time.Hour),
			}
			val := now.Add(2 * time.Hour)
			Expect(t, &val).ToBeIn(timeSlice)
		})

		t.It("should pass when value is NOT in []time.Time", func() {
			now := time.Now()
			timeSlice := []time.Time{
				now,
				now.Add(1 * time.Hour),
				now.Add(2 * time.Hour),
			}
			val := now.Add(3 * time.Hour)
			Expect(t, &val).Not().ToBeIn(timeSlice)
		})

		t.It("should pass when value is in custom struct slice", func() {
			personSlice := []Person{
				{Name: "foo", Age: 20},
				{Name: "bar", Age: 30},
				{Name: "baz", Age: 40},
			}
			val := Person{Name: "bar", Age: 30}
			Expect(t, &val).ToBeIn(personSlice)
		})

		t.It("should pass when value is NOT in custom struct slice", func() {
			personSlice := []Person{
				{Name: "foo", Age: 20},
				{Name: "bar", Age: 30},
				{Name: "baz", Age: 40},
			}
			val := Person{Name: "bar", Age: 40}
			Expect(t, &val).Not().ToBeIn(personSlice)
		})
	})
}
