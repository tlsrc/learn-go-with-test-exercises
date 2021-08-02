package dictionary

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryErr("word not found")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exists")
)

func (d Dictionary) Search(word string) (string, error) {
	def, found := d[word]

	if !found {
		return "", ErrNotFound
	}

	return def, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	if err == nil {
		return ErrWordExists
	}

	d.set(word, definition)
	return nil
}

func (d Dictionary) set(word string, definition string) {
	d[word] = definition
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	if err != nil {
		return ErrWordDoesNotExist
	}

	d.set(word, definition)
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
