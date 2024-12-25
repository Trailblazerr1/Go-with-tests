package dictionary

import (
	"errors"
)

type Dict map[string]string

var errorKeyNotFound = errors.New("Key not found")
var errorKeyAlreadyExists = errors.New("Key already exists")

func (d Dict) Search(word string) (string, error) {
	//map returns two value, first is the actual value and second is bool, whether val is found
	//or not
	value, found := d[word]
	if !found {
		return "", errorKeyNotFound
	}
	return value, nil
}

func (d Dict) Add(key, value string) (bool, error) {
	_, err := d.Search(key)
	if err == nil {
		return false, errorKeyAlreadyExists
	}
	d[key] = value
	return true, nil
}

func (d Dict) Update(key, value string) error {
	_, err := d.Search(key)
	if err != nil {
		return errorKeyNotFound
	}
	d[key] = value
	return nil
}

func (d Dict) Delete(key string) error {
	_, err := d.Search(key)
	if err != nil {
		return errorKeyNotFound
	}
	delete(d, key)
	return nil
}
