package gt

import (
	"testing"
	"time"
)

func TestSuiteGest(testingT *testing.T) {
	t := CreateTest(testingT)
	mockTrue := true

	t.BeforeAll(func() {
		Expect(t, &mockTrue).ToBe(true)
	})
	t.BeforeEach(func() {
		Expect(t, &mockTrue).ToBe(true)
	})
	t.AfterEach(func() {
		Expect(t, &mockTrue).ToBe(true)
	})
	t.AfterAll(func() {
		Expect(t, &mockTrue).ToBe(true)
	})

	t.Skip().Describe("Testing Gest Describe", func() {

		t.It("also should work with Describe", func() {
			Expect(t, &mockTrue).ToBe(true)
		})
		t.Async().It("should be async", func() {
			time.Sleep(3 * time.Second)
		})
		t.Todo("Todo function should do nothing but print todo message")

		t.Skip().It("should be skipped", func() {
			//
		})
	})
	// TODO: 他のgestの関数・メソッドのテストも書く
}
