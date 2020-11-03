package maps2

import "errors"

type Dictionary map[string]string

var (
	ErrTermNotFound           = errors.New("Result not found")
	ErrTermExists             = errors.New("Term already present")
	ErrWordToUpdateNotPresent = errors.New("Word not present")
)

func (d Dictionary) Search(term string) (string, error) {
	result, ok := d[term]
	if !ok {
		return "", ErrTermNotFound
	}

	return result, nil
}

func (d Dictionary) Add(term, value string) error {
	_, ok := d[term]
	if ok {
		return ErrTermExists
	}

	d[term] = value
	return nil
}

func (d Dictionary) Update(term, value string) error {
	_, ok := d[term]
	if !ok {
		return ErrWordToUpdateNotPresent
	}

	d[term] = value
	return nil
}

func (d Dictionary) Delete(term string) {
	delete(d, term)
}
