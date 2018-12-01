package dictionary

import "errors"

var ErrNotFound = errors.New("could not find that word")

type Dictionary map[string]string

func (d Dictionary) Search(search string) (string, error) {
	definition, ok := d[search]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
