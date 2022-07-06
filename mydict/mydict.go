package mydict

import "errors"

type Dictionary map[string]string

var errorNotFound = errors.New("Not Found")
var errorWordExists = errors.New("Already Exists")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errorNotFound:
		d[word] = def
	case nil:
		return errorWordExists
	}
	return nil
}
