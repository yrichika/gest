package gt

import (
	"testing"
)

func TestGestCommon(testingT *testing.T) {
	t := CreateTest(testingT)

	t.Describe("InArray", func() {
		t.It("should return true if specified value exists in array", func() {
			intResult := InArray(3, []int{1, 2, 3, 4, 5})
			Expect(t, &intResult).ToBe(true)

			strResult := InArray("hoge", []string{"hoge", "fuga", "piyo"})
			Expect(t, &strResult).ToBe(true)
		})

		t.It("should return false if specified value does NOT exist in array", func() {
			intResult := InArray(6, []int{1, 2, 3, 4, 5})
			Expect(t, &intResult).ToBe(false)

			strResult := InArray("foo", []string{"hoge", "fuga", "piyo"})
			Expect(t, &strResult).ToBe(false)
		})
	})

	t.Describe("GetAllTestFileDirectories", func() {
		t.It("should return all test file directories", func() {
			// TODO:
			// 実行を仮に別のディレクトリに変更できないか?
			// もしくは、TempDirを使ってテストディレクトリを作成してテストを実行するか?
			expectedDirectories := []string{"."}

			result := GetAllTestFileDirectories(false)
			Expect(t, &result).ToDeepEqual(expectedDirectories)
		})
	})
}
