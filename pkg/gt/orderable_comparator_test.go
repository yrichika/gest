package gt

import (
	"testing"
	"time"
)

func TestComparators(testingT *testing.T) {

	// Not testing very simple functions like GreaterThan, LessThan, etc
	// they are too simple and clear to test

	t1 := CreateTest(testingT)
	t1.Describe("Comparators", func() {
		t1.It("Between with int", func() {
			bt := Between(1, 3)
			r := bt(2)
			Expect(t1, &r).ToBe(true)
		})

		t1.It("Between with time.Duration", func() {
			bt := Between(time.Second, 3*time.Second)
			r := bt(2 * time.Second)
			Expect(t1, &r).ToBe(true)
		})

		t1.It("TimeBetween", func() {
			oneMinuteAgo := time.Now().Add(-1 * time.Minute)
			oneMinuteLater := time.Now().Add(1 * time.Minute)
			bt := TimeBetween(oneMinuteAgo, oneMinuteLater)
			r := bt(time.Now())
			Expect(t1, &r).ToBe(true)
		})
	})

}
