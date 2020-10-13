package integers1

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2, 4)
	want := 6

	if got != want {
		t.Errorf("got: %d; want: %d", got, want)
	}
}

func ExampleAdd() {
	fmt.Println(Add(4, 4))
	// Output:
	// 8
}
