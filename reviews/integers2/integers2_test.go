package integers2

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(8, 8)
	want := 16

	if got != want {
		t.Errorf("got: %d; want: %d", got, want)
	}
}

func ExampleAdd() {
	fmt.Println(Add(50, 50))
	// Output:
	// 100
}
