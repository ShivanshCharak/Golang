package main

const (
	ErrorNotFound    = DictionaryErr("String not found")
	ErrorNoWordFound = DictionaryErr("No value for the key found")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	defintion, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return defintion, nil
}

func (d Dictionary) Update(word, val string) (string, error) {
	_, ok := d[word]
	if !ok {
		return "", ErrorNoWordFound
	}
	d[word] = val
	return d[word], nil
}

func (d Dictionary) Delete(word, val string) (string, error) {
	_, ok := d[word]
	if !ok {
		return "", ErrorNoWordFound
	}
	delete(d, word)
}
