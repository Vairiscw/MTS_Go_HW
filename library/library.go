package library

import "errors"

type Storage interface {
	GetBookByID(Id string) Book
	GetBookByName(name string) Book
}

type MapBookStorage struct {
	BooksMap    map[string]Book
	IDGenerator IDGenerator
}

func (ms *MapBookStorage) GetBookByID(Id string) Book {
	book, ok := ms.BooksMap[Id]
	if ok {
		return book
	}
	return Book{}
}

func (ms *MapBookStorage) GetBookByName(name string) Book {
	for _, book := range ms.BooksMap {
		if book.Name == name {
			return book
		}
	}
	return Book{}
}

func (ms *MapBookStorage) AddBook(book Book) (string, error) {
	if book.ID != "" {
		if _, ok := ms.BooksMap[book.ID]; ok {
			return "", errors.New("book already exists")
		}
		ms.BooksMap[book.ID] = book
		return book.ID, nil
	}

	id := ms.IDGenerator(book.Name)
	book.ID = id
	ms.BooksMap[id] = book
	return id, nil
}

type ArrayBookStorage struct {
	BooksArray  []Book
	IDGenerator IDGenerator
}

func (ms *ArrayBookStorage) AddBook(book Book) (string, error) {
	if book.ID != "" {
		for _, curBook := range ms.BooksArray {
			if curBook.ID == book.ID {
				return "", errors.New("book already exists")
			}
		}
		ms.BooksArray = append(ms.BooksArray, book)
		return book.ID, nil
	}
	id := ms.IDGenerator(book.Name)
	book.ID = id
	ms.BooksArray = append(ms.BooksArray, book)
	return id, nil
}

func (ms *ArrayBookStorage) GetBookByID(Id string) Book {
	for _, book := range ms.BooksArray {
		if book.ID == Id {
			return book
		}
	}
	return Book{}
}

func (ms *ArrayBookStorage) GetBookByName(name string) Book {
	for _, book := range ms.BooksArray {
		if book.Name == name {
			return book
		}
	}
	return Book{}
}
