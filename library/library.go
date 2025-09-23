package library

type Library struct {
	IdGenerator idGenerator
	BooksSlice  []Book
	BooksMap    map[string]Book
}

func (l *Library) AddBook(book Book) string {
	id := l.IdGenerator(book.Name)
	l.BooksMap[id] = book
	l.BooksSlice = append(l.BooksSlice, book)
	return id
}

func (l *Library) GetBookById(id string) Book {
	book, ok := l.BooksMap[id]
	if ok {
		return book
	} else {
		return Book{}
	}
}

func (l *Library) GetBookByName(name string) Book {
	for _, book := range l.BooksSlice {
		if book.Name == name {
			return book
		}
	}
	return Book{}
}
