package examples

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
)

// gest test for failure
func TestFailBehavior(testingT *testing.T) {

	t := gt.CreateTest(testingT)

	t.Describe("Gest test for failure", func() {
		t.It("should fail and show fail messages", func() {
			v := false
			gt.Expect(t, &v).ToBeTrue()
		})

	})

}
