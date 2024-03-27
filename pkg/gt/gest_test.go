package gt

import (
	"testing"
	"time"
)

func TestSuiteGestRunner(testingT *testing.T) {
	t := CreateTest(testingT)
	mockTrue := true

	t.BeforeAll(func() {
		Expect(t, &mockTrue).ToBe(true)
		println("BeforeAll")
	})
	t.BeforeEach(func() {
		Expect(t, &mockTrue).ToBe(true)
		println("BeforeEach")
	})
	t.AfterEach(func() {
		Expect(t, &mockTrue).ToBe(true)
		println("AfterEach")
	})
	t.AfterAll(func() {
		Expect(t, &mockTrue).ToBe(true)
		println("AfterAll")
	})

	t.Describe("Testing Gest Describe", func() {

		t.It("also should work with Describe", func() {
			Expect(t, &mockTrue).ToBe(true)
		})

		t.Test("Test function should work as It does", func() {
			Expect(t, &mockTrue).ToBe(true)
		})

		t.Parallel().It("should be Parallel", func() {
			// actually not sure how to Parallel functionality,
			// but it should be working
			time.Sleep(300 * time.Millisecond)
		})

		t.Todo("Todo function should do nothing but print todo message")

		t.Skip().It("should be skipped", func() {
			Expect(t, &mockTrue).ToBe(false) // this should not be executed
		})
	})

	t2 := CreateTest(testingT)
	t2.Skip().Describe("this test should be skipped", func() {
		Expect(t2, &mockTrue).ToBe(false) // this should not be executed
	})
}
