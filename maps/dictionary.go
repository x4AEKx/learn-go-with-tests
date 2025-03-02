package maps

import "errors"

var (
	ErrNotFound         = errors.New("could not find the word you were looking for")
	ErrWordExists       = errors.New("value is existing")
	ErrWordDoesNotExist = errors.New("value dosnt exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, key)
	default:
		return nil
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return nil
	}

	return nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
