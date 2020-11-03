package maps2

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"my term": "found!"}

	got, err := dict.Search("my term")
	want := "found!"

	if err != nil {
		t.Fatal("Error not expected")
	}

	if got != want {
		t.Errorf("got: %s; want: %s", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("New Word", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Add("New", "Value")

		if err != nil {
			t.Fatal("Error not expected")
		}

		got, _ := dict.Search("New")
		want := "Value"

		if got != want {
			t.Errorf("got: %s; want: %s", got, want)
		}
	})

	t.Run("Exiting Word", func(t *testing.T) {
		dict := Dictionary{"New": "Value"}

		err := dict.Add("New", "Value")
		want := ErrTermExists

		if err != want {
			t.Errorf("got: %v; want: %v", err, want)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update existing word", func(t *testing.T) {
		dict := Dictionary{"New": "Value"}

		err := dict.Update("New", "Updated Value")

		if err != nil {
			t.Fatal("Error not expected")
		}

		got, err := dict.Search("New")
		want := "Updated Value"

		if got != want {
			t.Errorf("got: %s; want: %s", got, want)
		}
	})

	t.Run("Update word not present", func(t *testing.T) {
		dict := Dictionary{"New": "Value"}

		err := dict.Update("Old", "Updated Value")
		want := ErrWordToUpdateNotPresent

		if err != want {
			t.Errorf("got: %v; want: %v", err, want)
		}
	})
}

func TestDelete(t *testing.T) {
	dict := Dictionary{"New": "Value"}

	dict.Delete("New")

	_, err := dict.Search("New")

	if err != ErrTermNotFound {
		t.Errorf("got: %v; want: %v", err, ErrTermNotFound)
	}
}
