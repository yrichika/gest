package gt

import (
	"testing"
)

func TestGestCommon(testingT *testing.T) {
	t := CreateTest(testingT)

	t.Describe("IsInSlice", func() {
		t.It("should return true if specified value exists in array", func() {
			intResult := IsInSlice(3, []int{1, 2, 3, 4, 5})
			Expect(t, &intResult).ToBe(true)

			strResult := IsInSlice("hoge", []string{"hoge", "fuga", "piyo"})
			Expect(t, &strResult).ToBe(true)
		})

		t.It("should return false if specified value does NOT exist in array", func() {
			intResult := IsInSlice(6, []int{1, 2, 3, 4, 5})
			Expect(t, &intResult).ToBe(false)

			strResult := IsInSlice("foo", []string{"hoge", "fuga", "piyo"})
			Expect(t, &strResult).ToBe(false)
		})
	})

	t2 := CreateTest(testingT)
	t2.Describe("GetAllTestFileDirectories", func() {
		t2.It("should return all test file directories", func() {
			// TODO:
			// 実行を仮に別のディレクトリに変更できないか?
			// もしくは、TempDirを使ってテストディレクトリを作成してテストを実行するか?
			expectedDirectories := []string{"."}

			result := GetAllTestFileDirectories(false)
			Expect(t2, &result).ToBe(expectedDirectories)
		})
	})
}
