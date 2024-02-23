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
			bt := Between(1)
			r := bt(2, 3)
			Expect(t1, &r).ToBe(true)
		})

		t1.It("Between with time.Duration", func() {
			bt := Between(time.Second)
			r := bt(2*time.Second, 3*time.Second)
			Expect(t1, &r).ToBe(true)
		})

		t1.It("TimeBetween", func() {
			oneMinuteAgo := time.Now().Add(-1 * time.Minute)
			oneMinuteLater := time.Now().Add(1 * time.Minute)
			bt := TimeBetween(oneMinuteAgo)
			r := bt(time.Now(), oneMinuteLater)
			Expect(t1, &r).ToBe(true)
		})
	})

}
