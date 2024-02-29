package gt

import (
	"testing"
	"time"
)

func TestArrayAssertions(testingT *testing.T) {

	t := CreateTest(testingT)
	t.Describe("ToBeIn", func() {
		t.It("should pass when value is in []int", func() {
			intArr := []int{1, 2, 3, 4, 5}
			val := 3
			Expect(t, &val).ToBeIn(intArr)
		})
		t.It("should pass when value is NOT in []int", func() {
			intArr := []int{1, 2, 3, 4, 5}
			val := 6
			Expect(t, &val).Not().ToBeIn(intArr)
		})

		t.It("should pass when value is in []bool", func() {
			boolArr := []bool{false, false, false, true}
			val := true
			Expect(t, &val).ToBeIn(boolArr)
		})

		t.It("should pass when value is NOT in []bool", func() {
			boolArr := []bool{false, false, false}
			val := true
			Expect(t, &val).Not().ToBeIn(boolArr)
		})

		t.It("should pass when value is in []string", func() {
			strArr := []string{"foo", "bar", "baz"}
			val := "bar"
			Expect(t, &val).ToBeIn(strArr)
		})

		t.It("should pass when value is NOT in []string", func() {
			strArr := []string{"foo", "bar", "baz"}
			val := "hoge"
			Expect(t, &val).Not().ToBeIn(strArr)
		})

		t.It("should pass when value is in []time.Duration", func() {
			durationArr := []time.Duration{
				1 * time.Second,
				2 * time.Second,
				3 * time.Second,
			}
			val := 3 * time.Second
			Expect(t, &val).ToBeIn(durationArr)
		})

		t.It("should pass when value is NOT in []time.Duration", func() {
			durationArr := []time.Duration{
				1 * time.Second,
				2 * time.Second,
				3 * time.Second,
			}
			val := 4 * time.Second
			Expect(t, &val).Not().ToBeIn(durationArr)
		})

		t.It("should pass when value is in []time.Time", func() {
			now := time.Now()
			timeArr := []time.Time{
				now,
				now.Add(1 * time.Hour),
				now.Add(2 * time.Hour),
			}
			val := now.Add(2 * time.Hour)
			Expect(t, &val).ToBeIn(timeArr)
		})

		t.It("should pass when value is NOT in []time.Time", func() {
			now := time.Now()
			timeArr := []time.Time{
				now,
				now.Add(1 * time.Hour),
				now.Add(2 * time.Hour),
			}
			val := now.Add(3 * time.Hour)
			Expect(t, &val).Not().ToBeIn(timeArr)
		})

		t.It("should pass when value is in custom struct slice", func() {
			personArr := []Person{
				{Name: "foo", Age: 20},
				{Name: "bar", Age: 30},
				{Name: "baz", Age: 40},
			}
			val := Person{Name: "bar", Age: 30}
			Expect(t, &val).ToBeIn(personArr)
		})

		t.It("should pass when value is NOT in custom struct slice", func() {
			personArr := []Person{
				{Name: "foo", Age: 20},
				{Name: "bar", Age: 30},
				{Name: "baz", Age: 40},
			}
			val := Person{Name: "bar", Age: 40}
			Expect(t, &val).Not().ToBeIn(personArr)
		})
	})
}
