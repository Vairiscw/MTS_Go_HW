package main

import (
	"MTS_Go_HW/library"
	"fmt"
)

func tests() {
	var bookSlice []library.Book
	bookSlice = append(bookSlice, library.Book{Name: "Harry Potter", Author: "J. K. Rowling", Year: 1997})
	bookSlice = append(bookSlice, library.Book{Name: "Курс теоретической физики", Author: "Л. Д. Ландау", Year: 1960})
	bookSlice = append(bookSlice, library.Book{Name: "Grokking Algorithms", Author: "Aditya Bhargava", Year: 2016})

	lib := library.Library{
		BooksSlice:  make([]library.Book, 0),
		IdGenerator: library.IdGeneratorOne,
		BooksMap:    make(map[string]library.Book),
	}

	var ids []string
	for _, book := range bookSlice {
		ids = append(ids, lib.AddBook(book))
	}

	// Тестики
	if lib.GetBookById(ids[0]).Name != "Harry Potter" {
		fmt.Println("TEST1: ERROR")
	} else {
		fmt.Println("TEST1: OK")
	}
	if lib.GetBookById(ids[2]).Name != "Grokking Algorithms" {
		fmt.Println("TEST2: ERROR")
	} else {
		fmt.Println("TEST2: OK")
	}
	if lib.GetBookByName("Курс теоретической физики").Name != "Курс теоретической физики" {
		fmt.Println("TEST3: ERROR")
	} else {
		fmt.Println("TEST3: OK")
	}
	if lib.GetBookByName("Grokking mainstream").Name != "" {
		fmt.Println("TEST4: ERROR")
	} else {
		fmt.Println("TEST4: OK")
	}

	lib.IdGenerator = library.IdGeneratorTwo
	lib.BooksSlice = make([]library.Book, 0)
	lib.BooksMap = make(map[string]library.Book)

	ids = make([]string, 0)
	for _, book := range bookSlice {
		ids = append(ids, lib.AddBook(book))
	}
	if lib.GetBookById(ids[0]).Name != "Harry Potter" {
		fmt.Println("TEST5: ERROR")
	} else {
		fmt.Println("TEST5: OK")
	}
	if lib.GetBookById(ids[2]).Name != "Grokking Algorithms" {
		fmt.Println("TEST6: ERROR")
	} else {
		fmt.Println("TEST6: OK")
	}
}

func main() {
	tests()
}
