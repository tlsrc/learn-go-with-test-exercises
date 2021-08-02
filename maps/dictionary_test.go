package dictionary

import "testing"

func TestSeach(t *testing.T) {
	dic := Dictionary{"test": "This is just a test"}

	t.Run("Search for existing word", func(t *testing.T) {
		want := "This is just a test"
		word := "test"
		assertDefinition(t, dic, word, want)
	})

	t.Run("Search for missing word", func(t *testing.T) {
		_, err := dic.Search("unknown")
		want := ErrNotFound
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	dic := Dictionary{}
	word := "test"
	definition := "this is just a test"

	t.Run("new definition", func(t *testing.T) {
		err := dic.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dic, word, definition)
	})

	t.Run("existing definition", func(t *testing.T) {
		dic.Add(word, definition)

		err := dic.Add("test", "a new definition for test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dic, word, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dic := Dictionary{word: definition}

		newDefinition := "this is test but new"
		dic.Update(word, newDefinition)

		assertDefinition(t, dic, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		dic := Dictionary{}

		word := "test"
		newDefinition := "this is test but new"
		err := dic.Update(word, newDefinition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dic := Dictionary{word: definition}

		dic.Delete(word)
		word, err := dic.Search(word)

		assertError(t, err, ErrNotFound)
	})
}

func assertDefinition(t *testing.T, dic Dictionary, word, definition string) {
	got, err := dic.Search(word)
	if err != nil {
		t.Fatal("Definition should exist")
	}

	if got != definition {
		t.Errorf("Wanted definition %q but got %q", definition, got)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %q but got %q", want, got)
	}
}
