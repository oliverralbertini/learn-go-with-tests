package dictionary

const (
	ErrNotFound   = DictionaryErr("could not find that word")
	ErrWordExists = DictionaryErr("that word is already in the dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(search string) (string, error) {
	definition, ok := d[search]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add doesn't need to be a pointer since maps are 'reference types'
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = def
		return nil
	case nil:
		return ErrWordExists
	}
	return err
}
