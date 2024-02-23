package examples

import (
	"testing"
	"time"

	"github.com/yrichika/gest/pkg/gt"
)

type Person struct {
	Name string
	Age  int
}

// Testing for failure behaviors
// This is not a usual example or test case, but it's just checking for failure behaviors.
func TestFailureBehaviors(testingT *testing.T) {

	t := gt.CreateTest(testingT)

	t.Describe("Tests for failure behaviors", func() {

		t.It("fails with ToBe", func() {
			v := false
			gt.Expect(t, &v).ToBe(true)
			// Output: Failed at [failing_test.go]:line 23: actual:[false] is NOT expected:[true]
		})

		t.It("fails with Not.ToBe", func() {
			v := true
			gt.Expect(t, &v).Not().ToBe(true)
			// Output: Failed at [failing_test.go]:line 29: actual:[true] IS expected:[true]
		})

		t.It("fails with ToBeNil", func() {
			val := 1
			gt.Expect(t, &val).ToBeNil()
			// Output: Failed at [failing_test.go]:line 43: [1] is NOT nil
		})

		t.It("fails with Not.ToBeNil", func() {
			var val *int = nil
			gt.Expect(t, val).Not().ToBeNil()
			// Output: Failed at [failing_test.go]:line 49: Value IS nil
		})

		t.It("fails with ToBeSamePointerAs", func() {
			p := Person{Name: "hoge", Age: 1}
			o := p // not same pointer, copying

			gt.Expect(t, &p).ToBeSamePointerAs(&o)
			// Output: Failed at [failing_test.go]:line 57: Pointer to [examples.Person{Name:"hoge", Age:1}] is NOT the same
		})

		t.It("fails with Not.ToBeSamePointerAs", func() {
			p := Person{Name: "hoge", Age: 1}
			o := &p

			gt.Expect(t, &p).Not().ToBeSamePointerAs(o)
			// Output: Failed at [failing_test.go]:line 65: Pointer to [examples.Person{Name:"hoge", Age:1}] IS the same
		})

		t.It("fails with ToBe with time.Time", func() {
			time1 := time.Now()
			time2, _ := time.Parse("2006-01-02", "2021-01-01")
			gt.Expect(t, &time1).ToBe(time2)
			// Output: Failed at [failure_behavior_test.go]:line 66: actual:[time.Date(2024, time.February, 21, 18, 12, 6, 142277000, time.Local)] is NOT expected:[time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)]
		})

		t.It("fails with Not.ToBe with time.Time", func() {
			time1 := time.Now()
			time2 := time1
			gt.Expect(t, &time1).Not().ToBe(time2)
			// Output: Failed at [failure_behavior_test.go]:line 74: actual:[time.Date(2024, time.February, 21, 18, 12, 6, 142597000, time.Local)] IS expected:[time.Date(2024, time.February, 21, 18, 12, 6, 142597000, time.Local)]
		})

		t.It("fails with ToBe with struct type", func() {
			a := Person{Name: "hoge", Age: 1}
			b := Person{Name: "hog", Age: 1}

			gt.Expect(t, &a).ToBe(b)
			// Output: Failed at [failure_behavior_test.go]:line 87: actual:[examples.Person{Name:"hoge", Age:1}] is NOT expected:[examples.Person{Name:"hog", Age:1}]
		})

		t.It("fails with Not.ToBe with struct type", func() {
			a := Person{Name: "hoge", Age: 1}
			b := Person{Name: "hoge", Age: 1}

			gt.Expect(t, &a).Not().ToBe(b)
			// Output: Failed at [failing_test.go]:line 81: actual:[examples.Person{Name:"hoge", Age:1}] IS expected:[examples.Person{Name:"hoge", Age:1}]
		})

		t.It("fail with ToMatchRegex", func() {
			a := "hoge"
			gt.Expect(t, &a).ToMatchRegex("^foo$")
			// Output: Failed at [failure_behavior_test.go]:line 88: actual:["hoge"] does NOT match with regex expected:["^foo$"]
		})

		t.It("fail with ToMatchRegex NOT", func() {
			a := "hoge"
			gt.Expect(t, &a).Not().ToMatchRegex("^hoge$")
			// Output: Failed at [failure_behavior_test.go]:line 94: actual:["hoge"] DOES match with regex expected:["^hoge$"]
		})

		t.It("fail with ToContainString", func() {
			a := "hello world"
			gt.Expect(t, &a).ToContainString("bar")
			// Output: Failed at [failure_behavior_test.go]:line 100: actual:["hello world"] does NOT contain expected:["bar"]
		})

		t.It("fail with ToContainString NOT", func() {
			a := "hello world"
			gt.Expect(t, &a).Not().ToContainString("world")
			// Output: Failed at [failure_behavior_test.go]:line 106: actual:["hello world"] DOES contain expected:["world"]
		})

		t.It("fail with ExpectPanic", func() {
			gt.ExpectPanic(t).ToHappen(func() {
				// no panic
			})
			// Output: Panic: [failing_test.go]:line 86: Panic did NOT happen
		})

		t.It("fail with ExpectPanic Not", func() {
			gt.ExpectPanic(t).Not().ToHappen(func() {
				panic("panic!")
			})
			// Output: Panic: [failing_test.go]:line 93: Panic DID happen
		})

		t.It("LogWhenFail", func() {
			var a int = 1
			var b int = 2

			gt.LogWhenFail[int](t, "show this message when fail %#v, %#v").Expect(&a).ToBe(b)
			// Output: Failed at [failing_test.go]:line 103: show this message when fail 1, 2
		})

	})

	t2 := gt.CreateTest(testingT)
	t2.Describe("Tests for failure behaviors with ToBe_", func() {
		t2.It("fails with ToBe_", func() {
			intVal1 := 10
			gt.Expect(t2, &intVal1).ToBe_(gt.GreaterThan, 11)
			// Output: Failed at [failure_behavior_test.go]:line 139: compared actual:[10] and expected:[11]
		})
	})

	t3 := gt.CreateTest(testingT)
	t3.Describe("Tests for failure behaviors with Test method", func() {
		// Testing `Test` method. It should work exactly the same as `It` method.
		t3.Test("fails with Test method", func() {
			v := false
			gt.Expect(t3, &v).ToBe(true)
			// Output: Failed at [failure_behavior_test.go]:line 154: actual:[false] is NOT expected:[true]
		})
	})

	t4 := gt.CreateTest(testingT)
	t4.Describe("type assertion failure", func() {
		t4.It("fails with ToBeType", func() {
			valInt := 1
			gt.Expect(t4, &valInt).ToBeType(gt.OfString)
			// Output: Failed at [failure_behavior_test.go]:line 163: actual:[int] is NOT expected type
		})

		t4.It("fails with ToBeType with Not", func() {
			valStr := "abc"
			gt.Expect(t4, &valStr).Not().ToBeType(gt.OfString)
			// Output: Failed at [failure_behavior_test.go]:line 169: actual:[string] IS expected type
		})
	})

}
