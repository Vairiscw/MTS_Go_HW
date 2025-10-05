package library

import "errors"

type Storage interface {
	GetBookById(Id string) Book
	GetBookByName(name string) Book
}

type MapBookStorage struct {
	BooksMap    map[string]Book
	IdGenerator IdGenerator
}

func (ms *MapBookStorage) GetBookById(Id string) Book {
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
	if book.Id != "" {
		if _, ok := ms.BooksMap[book.Id]; ok {
			return "", errors.New("book already exists")
		}
		ms.BooksMap[book.Id] = book
		return book.Id, nil
	}

	id := ms.IdGenerator(book.Name)
	book.Id = id
	ms.BooksMap[id] = book
	return id, nil
}

type ArrayBookStorage struct {
	BooksArray  []Book
	IdGenerator IdGenerator
}

func (ms *ArrayBookStorage) AddBook(book Book) (string, error) {
	if book.Id != "" {
		for _, curBook := range ms.BooksArray {
			if curBook.Id == book.Id {
				return "", errors.New("book already exists")
			}
		}
		ms.BooksArray = append(ms.BooksArray, book)
		return book.Id, nil
	}
	id := ms.IdGenerator(book.Name)
	book.Id = id
	ms.BooksArray = append(ms.BooksArray, book)
	return id, nil
}

func (ms *ArrayBookStorage) GetBookById(Id string) Book {
	for _, book := range ms.BooksArray {
		if book.Id == Id {
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
