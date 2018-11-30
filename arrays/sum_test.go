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

func TestSumAllTails(t *testing.T) {
	slc1 := []int{1, 2, 3}
	slc2 := []int{4, 5, 6}

	got := SumAllTails(slc1, slc2)
	want := []int{5, 11}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted '%v' got '%v' when summing tails of %v and %v", want, got, slc1, slc2)
	}
}

func TestSumAllHeads(t *testing.T) {
	slc1 := []int{1, 2, 3}
	slc2 := []int{4, 5, 6}

	got := SumAllHeads(slc1, slc2)
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted '%v' got '%v' when summing heads of %v and %v", want, got, slc1, slc2)
	}
}

func TestSumAllIns(t *testing.T) {
	slc1 := []int{1, 2, 3}
	slc2 := []int{4, 5, 9, 6}
	slc3 := []int{4, 5}

	got := SumAllIns(slc1, slc2, slc3)
	want := []int{2, 14, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted '%v' got '%v' when summing insides of %v, %v, and %v", want, got, slc1, slc2, slc3)
	}
}
