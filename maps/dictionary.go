package main

import "errors"

var ErrNotFound = errors.New("couldn't find the word you were looking for")
var ErrWordExists = errors.New("word already exist")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	defination, found := d[word]

	if !found {
		return "", ErrNotFound
	}

	return defination, nil
}

func (d Dictionary) Add(word, defination string) error {
	_, found := d[word]
	if found {
		return ErrWordExists
	}
	d[word] = defination
	return nil
}

func (d Dictionary) Update(word, defination string) {
	d[word] = defination
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
