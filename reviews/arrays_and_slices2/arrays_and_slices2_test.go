package arrays_and_slices2

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3, 4, 4})
	want := 14

	if got != want {
		t.Errorf("got: %d; want: %d", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %d; want: %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %d; want: %d", got, want)
	}
}
