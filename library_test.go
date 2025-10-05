package main

import (
	"MTS_Go_HW/library"
	"testing"
)

func TestArrayStorage(t *testing.T) {
	arrayStorage := library.ArrayBookStorage{
		BooksArray:  make([]library.Book, 0),
		IdGenerator: library.IdGeneratorAddRandomText,
	}

	var bookSlice []library.Book
	bookSlice = append(bookSlice, library.Book{Name: "Harry Potter", Author: "J. K. Rowling", Year: 1997})
	bookSlice = append(bookSlice, library.Book{Name: "Курс теоретической физики", Author: "Л. Д. Ландау", Year: 1960})
	bookSlice = append(bookSlice, library.Book{Name: "Grokking Algorithms", Author: "Aditya Bhargava", Year: 2016})

	var ids []string
	for _, book := range bookSlice {
		id, err := arrayStorage.AddBook(book)
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, id)
	}

	if arrayStorage.GetBookByName("Harry Potter").Name != "Harry Potter" {
		t.Errorf("Doesn't find book")
	}
	if arrayStorage.GetBookByName("Grokking Streaming").Name != "" {
		t.Errorf("Find something that doesn't exist")
	}
	if id := arrayStorage.GetBookById(ids[1]).Id; id != ids[1] {
		t.Errorf("Find something that doesn't exist. Expected: %s Find: %s", ids[1], id)
	}
	if id := arrayStorage.GetBookById(ids[0]).Id; id != ids[0] {
		t.Errorf("Find something that doesn't exist Expected: %s Find: %s", ids[0], id)
	}

	arrayStorage.IdGenerator = library.IdGeneratorAddPrefix
	arrayStorage.BooksArray = make([]library.Book, 0)

	ids = make([]string, 0)
	for _, book := range bookSlice {
		id, err := arrayStorage.AddBook(book)
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, id)
	}

	if arrayStorage.GetBookById(ids[0]).Name != "Harry Potter" {
		t.Errorf("Doesn't find book")
	}
	if arrayStorage.GetBookById(ids[2]).Name != "Grokking Algorithms" {
		t.Errorf("Doesn't find book")
	}
}

func TestMapStorage(t *testing.T) {
	mapStorage := library.MapBookStorage{
		BooksMap:    make(map[string]library.Book),
		IdGenerator: library.IdGeneratorAddRandomText,
	}

	var bookSlice []library.Book
	bookSlice = append(bookSlice, library.Book{Name: "Harry Potter", Author: "J. K. Rowling", Year: 1997})
	bookSlice = append(bookSlice, library.Book{Name: "Курс теоретической физики", Author: "Л. Д. Ландау", Year: 1960})
	bookSlice = append(bookSlice, library.Book{Name: "Grokking Algorithms", Author: "Aditya Bhargava", Year: 2016})

	var ids []string
	for _, book := range bookSlice {
		id, err := mapStorage.AddBook(book)
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, id)
	}

	if mapStorage.GetBookByName("Harry Potter").Name != "Harry Potter" {
		t.Errorf("Doesn't find book")
	}
	if mapStorage.GetBookByName("Grokking Streaming").Name != "" {
		t.Errorf("Find something that doesn't exist")
	}
	if id := mapStorage.GetBookById(ids[1]).Id; id != ids[1] {
		t.Errorf("Find something that doesn't exist. Expected: %s Find: %s", ids[1], id)
	}
	if id := mapStorage.GetBookById(ids[0]).Id; id != ids[0] {
		t.Errorf("Find something that doesn't exist Expected: %s Find: %s", ids[0], id)
	}

	mapStorage.IdGenerator = library.IdGeneratorAddPrefix
	mapStorage.BooksMap = make(map[string]library.Book)

	ids = make([]string, 0)
	for _, book := range bookSlice {
		id, err := mapStorage.AddBook(book)
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, id)
	}

	if mapStorage.GetBookById(ids[0]).Name != "Harry Potter" {
		t.Errorf("Doesn't find book")
	}
	if mapStorage.GetBookById(ids[2]).Name != "Grokking Algorithms" {
		t.Errorf("Doesn't find book")
	}
}
