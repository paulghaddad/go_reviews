package arrays_and_slices1

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{10, 20, 30, 40, 50}

	got := Sum(nums)
	want := 150

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestSumSlices(t *testing.T) {
	got := SumSlices([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
