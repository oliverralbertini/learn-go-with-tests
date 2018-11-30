package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("when using an arbitrary sized collection", func(t *testing.T) {
		slc := []int{1, 2, 3}

		got := Sum(slc)
		want := 6

		if got != want {
			t.Errorf("wanted '%d' got '%d' when summing %v", want, got, slc)
		}
	})
}

func TestSumAll(t *testing.T) {
	slc1 := []int{1, 2, 3}
	slc2 := []int{4, 5, 6}

	got := SumAll(slc1, slc2)
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted '%v' got '%v' when summing %v and %v", want, got, slc1, slc2)
	}
}

func compareSlices(t *testing.T, got, want, slc1, slc2 []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted '%v' got '%v' with: %v and %v", want, got, slc1, slc2)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum the tails of all slices", func(t *testing.T) {
		slc1 := []int{1, 2, 3}
		slc2 := []int{4, 5, 6}

		got := SumAllTails(slc1, slc2)
		want := []int{5, 11}

		compareSlices(t, got, want, slc1, slc2)
	})

	t.Run("Safely sum tails of empty slices", func(t *testing.T) {
		slc1 := []int{}
		slc2 := []int{1}

		got := SumAllTails(slc1, slc2)
		want := []int{0, 0}

		compareSlices(t, got, want, slc1, slc2)
	})
}

func TestSumAllHeads(t *testing.T) {
	t.Run("Sum the heads of all slices", func(t *testing.T) {
		slc1 := []int{1, 2, 3}
		slc2 := []int{4, 5, 6}

		got := SumAllHeads(slc1, slc2)
		want := []int{3, 9}

		compareSlices(t, got, want, slc1, slc2)
	})

	t.Run("Safely sum heads of empty slices", func(t *testing.T) {
		slc1 := []int{}
		slc2 := []int{6}

		got := SumAllHeads(slc1, slc2)
		want := []int{0, 0}

		compareSlices(t, got, want, slc1, slc2)
	})
}

func TestSumAllIns(t *testing.T) {
	t.Run("Sum the heads of all slices", func(t *testing.T) {
		slc1 := []int{1, 2, 3}
		slc2 := []int{4, 5, 9, 6}
		slc3 := []int{4, 5}

		got := SumAllIns(slc1, slc2, slc3)
		want := []int{2, 14, 0}

		compareSlices(t, got, want, slc1, slc2)
	})

	t.Run("Safely sum heads of empty slices", func(t *testing.T) {
		slc1 := []int{}
		slc2 := []int{4}

		got := SumAllIns(slc1, slc2)
		want := []int{0, 0}

		compareSlices(t, got, want, slc1, slc2)
	})
}
