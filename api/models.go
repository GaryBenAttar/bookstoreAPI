package api

import (
	"errors"
	"sync"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type BookStore struct {
	books map[string]Book
	mutex sync.RWMutex
}

func NewBookStore() *BookStore {
	return &BookStore{
		books: make(map[string]Book),
	}
}

func (bs *BookStore) GetBooks() []Book {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()

	books := make([]Book, 0, len(bs.books))
	for _, book := range bs.books {
		books = append(books, book)
	}

	return books
}

func (bs *BookStore) GetBook(id string) (Book, error) {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()

	book, exists := bs.books[id]
	if !exists {
		return Book{}, errors.New("book not found")
	}

	return book, nil
}

func (bs *BookStore) CreateBook(book Book) error {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	if _, exists := bs.books[book.ID]; exists {
		return errors.New("book already exists")
	}

	bs.books[book.ID] = book
	return nil
}

func (bs *BookStore) UpdateBook(id string, book Book) error {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	if _, exists := bs.books[id]; !exists {
		return errors.New("book not found")
	}

	bs.books[id] = book
	return nil
}

func (bs *BookStore) DeleteBook(id string) error {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	if _, exists := bs.books[id]; !exists {
		return errors.New("book not found")
	}

	delete(bs.books, id)
	return nil
}

