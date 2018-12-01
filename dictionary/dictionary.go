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

// Add doesn't need to be a pointer since maps are 'reference types'
func (d Dictionary) Add(word, def string) {
	d[word] = def
}
