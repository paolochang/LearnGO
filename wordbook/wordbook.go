package wordbook

import (
	"errors"
)

// Wordbook type
type Wordbook map[string]string

var (
	errNotFound = errors.New("not found")
	errCantUpdate = errors.New("cannot update none-existing word")
	errCantDelete = errors.New("cannot delete noen-existing word")
	errWordExists = errors.New("word already exists in the dictionary")
)

// Search for a word
func (w Wordbook) Search(word string) (string, error) {
	value, exists := w[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the Wordbook
func (w Wordbook) Add(word, def string) error {
	_, err := w.Search(word)
	switch err {
	case errNotFound:
		w[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word in the Wordbook
func (w Wordbook) Update(word, definition string) error {
	_, err := w.Search(word)
	switch err {
	case nil:
		w[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word in the Wordbook
func (w Wordbook) Delete(word string) error {
	_, err := w.Search(word)
	switch err {
	case nil:
		delete(w, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}