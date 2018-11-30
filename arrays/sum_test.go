package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("when using an arbitrary sized collection", func(t *testing.T) {
		arr := []int{1, 2, 3}

		got := Sum(arr)
		want := 6

		if got != want {
			t.Errorf("wanted '%d' got '%d' when summing %v", want, got, arr)
		}
	})
}
