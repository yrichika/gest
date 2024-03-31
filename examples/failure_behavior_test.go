package examples

import (
	"errors"
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
			// Output: failure_behavior_test.go:26: actual:[false] is NOT expected:[true]
		})

		t.It("fails with Not.ToBe", func() {
			v := true
			gt.Expect(t, &v).Not().ToBe(true)
			// Output: failure_behavior_test.go:32: actual:[true] IS expected:[true]
		})

		t.It("fails with ToBeNil", func() {
			val := 1
			gt.Expect(t, &val).ToBeNil()
			// Output: failure_behavior_test.go:38: value is NOT nil: [1]
		})

		t.It("fails with Not.ToBeNil", func() {
			var val *int = nil
			gt.Expect(t, val).Not().ToBeNil()
			// Output: failure_behavior_test.go:44: value IS nil: [(*int)(nil)]
		})

		t.It("fails with ToBeNilInterface", func() {
			var err error = errors.New("error message")
			gt.Expect(t, &err).ToBeNilInterface()
			// Output: failure_behavior_test.go:50: the interface is NOT nil: [&errors.errorString{s:"error message"}]
		})

		t.It("fails with Not.ToBeNilInterface", func() {
			var err error
			gt.Expect(t, &err).Not().ToBeNilInterface()
			// Output: failure_behavior_test.go:56: interface IS nil
		})

		t.It("fails with ToBeSamePointerAs", func() {
			p := Person{Name: "hoge", Age: 1}
			o := p // not same pointer, copying

			gt.Expect(t, &p).ToBeSamePointerAs(&o)
			// Output: failure_behavior_test.go:64: [&examples.Person{Name:"hoge", Age:1}] is NOT the same. Expected: [&examples.Person{Name:"hoge", Age:1}]
		})

		t.It("fails with Not.ToBeSamePointerAs", func() {
			p := Person{Name: "hoge", Age: 1}
			o := &p

			gt.Expect(t, &p).Not().ToBeSamePointerAs(o)
			// Output: failure_behavior_test.go:72: [&examples.Person{Name:"hoge", Age:1}] IS the same. Expected: [&examples.Person{Name:"hoge", Age:1}]
		})

		t.It("fails ToBe with string", func() {
			str1 := "hoge"
			str2 := "hog"
			gt.Expect(t, &str1).ToBe(str2)
			// Output: failure_behavior_test.go:79: actual:["hoge"] is NOT expected:["hog"]
		})

		t.It("fails ToBe with string", func() {
			str1 := "hoge"
			str2 := "hoge"
			gt.Expect(t, &str1).Not().ToBe(str2)
			// Output: failure_behavior_test.go:86: actual:["hoge"] IS expected:["hoge"]
		})

		t.It("fails ToBe with time.Duration", func() {
			duration1 := 1 * time.Second
			duration2 := 2 * time.Second
			gt.Expect(t, &duration1).ToBe(duration2)
			// Output: failure_behavior_test.go:93: actual:[1s] is NOT expected:[2s]
		})

		t.It("fails Not.ToBe with time.Duration", func() {
			duration1 := 1 * time.Second
			duration2 := 1 * time.Second
			gt.Expect(t, &duration1).Not().ToBe(duration2)
			// Output: failure_behavior_test.go:100: actual:[1s] IS expected:[1s]
		})

		t.It("fails with ToBe with time.Time", func() {
			time1 := time.Now()
			time2, _ := time.Parse("2006-01-02", "2021-01-01")
			gt.Expect(t, &time1).ToBe(time2)
			// Output: failure_behavior_test.go:107: actual:[time.Date(2024, time.March, 27, 18, 14, 23, 838646000, time.Local)] is NOT expected:[time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)]
		})

		t.It("fails with Not.ToBe with time.Time", func() {
			time1 := time.Now()
			time2 := time1
			gt.Expect(t, &time1).Not().ToBe(time2)
			// Output: failure_behavior_test.go:114: actual:[time.Date(2024, time.March, 27, 18, 14, 23, 838810000, time.Local)] IS expected:[time.Date(2024, time.March, 27, 18, 14, 23, 838810000, time.Local)]
		})

		t.It("fails with ToBe with struct type", func() {
			a := Person{Name: "hoge", Age: 1}
			b := Person{Name: "hog", Age: 1}

			gt.Expect(t, &a).ToBe(b)
			// Output: failure_behavior_test.go:122: actual:[examples.Person{Name:"hoge", Age:1}] is NOT expected:[examples.Person{Name:"hog", Age:1}]
		})

		t.It("fails with Not.ToBe with struct type", func() {
			a := Person{Name: "hoge", Age: 1}
			b := Person{Name: "hoge", Age: 1}

			gt.Expect(t, &a).Not().ToBe(b)
			// Output: failure_behavior_test.go:130: actual:[examples.Person{Name:"hoge", Age:1}] IS expected:[examples.Person{Name:"hoge", Age:1}]
		})

		t.It("fail with ToMatchRegex", func() {
			a := "hoge"
			gt.Expect(t, &a).ToMatchRegex("^foo$")
			// Output: failure_behavior_test.go:136: actual:["hoge"] does NOT match with regex expected:["^foo$"]
		})

		t.It("fail with ToMatchRegex NOT", func() {
			a := "hoge"
			gt.Expect(t, &a).Not().ToMatchRegex("^hoge$")
			// Output: failure_behavior_test.go:142: actual:["hoge"] DOES match with regex expected:["^hoge$"]
		})

		t.It("fail with ToContainString", func() {
			a := "hello world"
			gt.Expect(t, &a).ToContainString("bar")
			// Output: failure_behavior_test.go:148: actual:["hello world"] does NOT contain expected:["bar"]
		})

		t.It("fail with ToContainString NOT", func() {
			a := "hello world"
			gt.Expect(t, &a).Not().ToContainString("world")
			// Output: failure_behavior_test.go:154: actual:["hello world"] DOES contain expected:["world"]
		})

		t.It("fail with ExpectPanic", func() {
			gt.ExpectPanic(t).ToHappen(func() {
				// no panic
			})
			// Output: failure_behavior_test.go:159: Panic did NOT happen
		})

		t.It("fail with ExpectPanic Not", func() {
			gt.ExpectPanic(t).Not().ToHappen(func() {
				panic("panic!")
			})
			// Output: failure_behavior_test.go:167: Panic DID happen
		})

		t.It("LogWhenFail", func() {
			var a int = 1
			var b int = 2

			gt.LogWhenFail[int](t, "show this message when fail %#v, %#v").Expect(&a).ToBe(b)
			// Output: failure_behavior_test.go:176: show this message when fail 1, 2
		})

	})

	t2 := gt.CreateTest(testingT)
	t2.Describe("Tests for failure behaviors with ToBe_", func() {
		t2.It("fails with ToBe_", func() {
			intVal1 := 10
			gt.Expect(t2, &intVal1).ToBe_(gt.GreaterThan(11))
			// Output: failure_behavior_test.go:186: compared values: actual:[10]
		})
	})

	t3 := gt.CreateTest(testingT)
	t3.Describe("Tests for failure behaviors with Test method", func() {
		// Testing `Test` method. It should work exactly the same as `It` method.
		t3.Test("fails with Test method", func() {
			v := false
			gt.Expect(t3, &v).ToBe(true)
			// Output: failure_behavior_test.go:196: actual:[false] is NOT expected:[true]
		})
	})

	t4 := gt.CreateTest(testingT)
	t4.Describe("type assertion failure", func() {
		t4.It("fails with ToBeType", func() {
			valInt := 1
			gt.Expect(t4, &valInt).ToBeType(gt.OfString)
			// Output: failure_behavior_test.go:205: actual:[*int] is NOT expected type
		})

		t4.It("fails with ToBeType with Not", func() {
			valStr := "abc"
			gt.Expect(t4, &valStr).Not().ToBeType(gt.OfString)
			// Output: failure_behavior_test.go:211: actual:[*string] IS expected type
		})
	})

	t5 := gt.CreateTest(testingT)
	t5.Describe("Tests for failure behaviors with ToBeIn", func() {
		t5.It("fails when value is in []int", func() {
			intSlice := []int{1, 2, 3, 4, 5}
			val := 6
			gt.Expect(t5, &val).ToBeIn(intSlice)
			// Output: failure_behavior_test.go:221: actual:[6] is NOT in expected:[[1 2 3 4 5]]
		})
		t5.It("fails when value is NOT in []int", func() {
			intSlice := []int{1, 2, 3, 4, 5}
			val := 3
			gt.Expect(t5, &val).Not().ToBeIn(intSlice)
			// Output: failure_behavior_test.go:227: actual:[3] IS in expected:[[1 2 3 4 5]]
		})

		t5.It("fails when value is in []bool", func() {
			boolSlice := []bool{false, false, false}
			val := true
			gt.Expect(t5, &val).ToBeIn(boolSlice)
			// Output: failure_behavior_test.go:234: actual:[true] is NOT in expected:[[false false false]]
		})

		t5.It("fails when value is NOT in []bool", func() {
			boolSlice := []bool{false, false, false}
			val := false
			gt.Expect(t5, &val).Not().ToBeIn(boolSlice)
			// Output: failure_behavior_test.go:241: actual:[false] IS in expected:[[false false false]]
		})

		t5.It("fails when value is in []string", func() {
			strSlice := []string{"foo", "bar", "baz"}
			val := "ban"
			gt.Expect(t5, &val).ToBeIn(strSlice)
			// Output: failure_behavior_test.go:248: actual:["ban"] is NOT in expected:[[foo bar baz]]
		})

		t5.It("fails when value is NOT in []string", func() {
			strSlice := []string{"foo", "bar", "baz"}
			val := "foo"
			gt.Expect(t5, &val).Not().ToBeIn(strSlice)
			// Output: failure_behavior_test.go:255: actual:["foo"] IS in expected:[[foo bar baz]]
		})

		t5.It("fails when value is in []time.Duration", func() {
			durationSlice := []time.Duration{
				1 * time.Second,
				2 * time.Second,
				3 * time.Second,
			}
			val := 4 * time.Second
			gt.Expect(t5, &val).ToBeIn(durationSlice)
			// Output: failure_behavior_test.go:266: actual:[4s] is NOT in expected:[[1s 2s 3s]]
		})

		t5.It("fails when value is NOT in []time.Duration", func() {
			durationSlice := []time.Duration{
				1 * time.Second,
				2 * time.Second,
				3 * time.Second,
			}
			val := 1 * time.Second
			gt.Expect(t5, &val).Not().ToBeIn(durationSlice)
			// Output: failure_behavior_test.go:277: actual:[1s] IS in expected:[[1s 2s 3s]]
		})

		t5.It("fails when value is in []time.Time", func() {
			now := time.Now()
			timeSlice := []time.Time{
				now,
				now.Add(1 * time.Hour),
				now.Add(2 * time.Hour),
			}
			val := now.Add(3 * time.Hour)
			gt.Expect(t5, &val).ToBeIn(timeSlice)
			// Output: failure_behavior_test.go:289: actual:[time.Date(2024, time.March, 27, 21, 14, 23, 839812000, time.Local)] is NOT in expected:[[]time.Time{time.Date(2024, time.March, 27, 18, 14, 23, 839812000, time.Local), time.Date(2024, time.March, 27, 19, 14, 23, 839812000, time.Local), time.Date(2024, time.March, 27, 20, 14, 23, 839812000, time.Local)}]
		})

		t5.It("fails when value is NOT in []time.Time", func() {
			now := time.Now()
			timeSlice := []time.Time{
				now,
				now.Add(1 * time.Hour),
				now.Add(2 * time.Hour),
			}
			val := now.Add(2 * time.Hour)
			gt.Expect(t5, &val).Not().ToBeIn(timeSlice)
			// Output: failure_behavior_test.go:301: actual:[time.Date(2024, time.March, 27, 20, 14, 23, 839858000, time.Local)] IS in expected:[[]time.Time{time.Date(2024, time.March, 27, 18, 14, 23, 839858000, time.Local), time.Date(2024, time.March, 27, 19, 14, 23, 839858000, time.Local), time.Date(2024, time.March, 27, 20, 14, 23, 839858000, time.Local)}]
		})

		t5.It("fails when value is in custom struct slice", func() {
			personSlice := []Person{
				{Name: "foo", Age: 20},
				{Name: "bar", Age: 30},
				{Name: "baz", Age: 40},
			}
			val := Person{Name: "ban", Age: 50}
			gt.Expect(t5, &val).ToBeIn(personSlice)
			// Output: failure_behavior_test.go:312: actual:[examples.Person{Name:"ban", Age:50}] is NOT in expected:[[]interface {}{examples.Person{Name:"foo", Age:20}, examples.Person{Name:"bar", Age:30}, examples.Person{Name:"baz", Age:40}}]
		})

		t5.It("fails when value is NOT in custom struct slice", func() {
			personSlice := []Person{
				{Name: "foo", Age: 20},
				{Name: "bar", Age: 30},
				{Name: "baz", Age: 40},
			}
			val := Person{Name: "foo", Age: 20}
			gt.Expect(t5, &val).Not().ToBeIn(personSlice)
			// Output: failure_behavior_test.go:323: actual:[examples.Person{Name:"foo", Age:20}] IS in expected:[[]interface {}{examples.Person{Name:"foo", Age:20}, examples.Person{Name:"bar", Age:30}, examples.Person{Name:"baz", Age:40}}]
		})
	})

}
