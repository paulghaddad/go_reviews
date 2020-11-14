package maps3

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Unknown word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("Unknown")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "This is just a test"

		err := dictionary.Add("test", definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "This is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "This is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "Updated definition"

		err := dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
		assertError(t, err, nil)
	})

	t.Run("Word doesn't exist", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		definition := "Updated definition"

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "This is just a test"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)

	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Got error %q; want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("Error not expected:", err)
	}
	if got != definition {
		t.Errorf("got %q; want %q", got, definition)
	}
}
