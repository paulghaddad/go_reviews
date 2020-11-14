package maps3

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("Term not found")
var ErrWordExists = errors.New("Term already exists")
var ErrWordDoesNotExist = errors.New("Word does not exist")

func (d Dictionary) Search(term string) (string, error) {

	value, ok := d[term]

	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(term, value string) error {
	_, ok := d[term]
	if ok {
		return ErrWordExists
	}

	d[term] = value
	return nil
}

func (d Dictionary) Update(term, value string) error {
	_, ok := d[term]
	if !ok {
		return ErrWordDoesNotExist
	}

	d[term] = value
	return nil
}

func (d Dictionary) Delete(term string) {
	delete(d, term)
}
