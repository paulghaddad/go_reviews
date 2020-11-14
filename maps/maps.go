package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("Could not find the word")
	ErrWordExists       = DictionaryErr("The word already exists")
	ErrWordDoesNotExist = DictionaryErr("Cannot update because word does not exist")
)

func (d Dictionary) Search(term string) (string, error) {
	result, ok := d[term]

	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}

func (d Dictionary) Add(term, value string) error {
	_, err := d.Search(term)

	switch err {
	case ErrNotFound:
		d[term] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	d[term] = value

	return nil
}

func (d Dictionary) Update(term, value string) error {
	_, err := d.Search(term)

	if err == ErrNotFound {
		return ErrWordDoesNotExist
	}

	d[term] = value
	return nil
}

func (d Dictionary) Delete(term string) {
	delete(d, term)
}

func (e DictionaryErr) Error() string {
	return string(e)
}
