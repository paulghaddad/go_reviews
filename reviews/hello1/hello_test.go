package hello1

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	var tests = []struct {
		name, language string
		want           string
	}{
		{"Paul", "English", "Hello, Paul!"},
		{"", "English", "Hello, World!"},
		{"Pablo", "Spanish", "Hola, Pablo!"},
		{"Paul", "French", "Bonjour, Paul!"},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%s; %s", test.name, test.want)
		t.Run(testname, func(t *testing.T) {
			got := Hello(test.name, test.language)
			want := test.want

			if got != want {
				t.Errorf("got: %s; want: %s", got, want)
			}
		})
	}
}
