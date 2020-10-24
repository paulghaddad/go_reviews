package hello2

import "testing"

func TestHello(t *testing.T) {
	t.Run("English", func(t *testing.T) {
		got := Hello("Paul", "English")
		want := "Hello, Paul!"

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("Default name", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World!"

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("Spanish", func(t *testing.T) {
		got := Hello("Pablo", "Spanish")
		want := "Hola, Pablo!"

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("French", func(t *testing.T) {
		got := Hello("Paul", "French")
		want := "Bonjour, Paul!"

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}
