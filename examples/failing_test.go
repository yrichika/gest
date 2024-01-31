package examples

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
)

type MockObject struct {
	Name string
	Age  int
}

// gest test for failure
func TestFailBehavior(testingT *testing.T) {

	t := gt.CreateTest(testingT)

	t.Describe("Gest test for failure", func() {
		// 	t.It("should fail and show fail messages", func() {
		// 		v := false
		// 		gt.Expect(t, &v).ToBe(true)
		// 	})
		// })

		// t.It("should fail when nil", func() {
		// 	var nilValue *int = nil
		// 	Expect(t, nilValue).Not().ToBeNil()
		t.It("should fail when two objects are NOT equal", func() {
			a := MockObject{Name: "hoge", Age: 1}
			b := MockObject{Name: "hog", Age: 1}

			gt.Expect(t, &a).ToBe(b)
		})
		t.It("pointer equality", func() {
			p := MockObject{Name: "hoge", Age: 1}
			o := p
			gt.Expect(t, &p).ToBeSamePointerAs(&o)
		})
		t.It("toBe messaging test", func() {
			var a = true
			gt.Expect(t, &a).Not().ToBe(true)

		})

		t.It("panic assertion", func() {
			gt.ExpectPanic(t).Not().ToHappen(func() {
				panic("panic!")
			})
		})
	})

}
