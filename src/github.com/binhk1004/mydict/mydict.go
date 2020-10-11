package mydict

import "errors"

//Dictionary type
type Dictionary map[string] string
var (
	errNotFound = errors.New("검색결과가 없습니다.")
	errWordExists = errors.New("해당 단어는 이미 등록 되어 있습니다.")
	errCantUpdate = errors.New("존재하지 않는 단어는 업데이트 할 수 없습니다.")
)

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}
//Add a word to dictionary
func (d Dictionary)  Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists

	}
	return nil
}

//Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate

	}
	return nil
}

//Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}