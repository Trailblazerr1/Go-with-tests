package dictionary

import (
	"errors"
)

type Dict map[string]string

var errorKeyNotFound = errors.New("Key not found")

func (d Dict) Search(word string) (string, error) {
	//map returns two value, first is the actual value and second is bool, whether val is found
	//or not
	value, found := d[word]
	if !found {
		return "", errorKeyNotFound
	}
	return value, nil
}
