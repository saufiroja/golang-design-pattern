package main

import "fmt"

// Generic Iterator interface
type Iterator[T any] interface {
	HasNext() bool
	Next() *T
}

// Generic Collection interface
type Collection[T any] interface {
	CreateIterator() Iterator[T]
}

// Concrete iterator for books
type BookIterator struct {
	books []string
	index int
}

func (bi *BookIterator) HasNext() bool {
	return bi.index < len(bi.books)
}

func (bi *BookIterator) Next() *string {
	if bi.HasNext() {
		book := bi.books[bi.index]
		bi.index++
		return &book
	}
	return nil
}

// Concrete collection for books
type BookCollection struct {
	books []string
}

func (bc *BookCollection) CreateIterator() Iterator[string] {
	return &BookIterator{
		books: bc.books,
		index: 0,
	}
}

// Concrete map for category books
type CategoryBooks map[string][]string

// Concrete collection for category books
type CategoryBookCollection struct {
	categoryBooks CategoryBooks
	category      string
}

func (cbc *CategoryBookCollection) CreateIterator() Iterator[string] {
	return &BookIterator{
		books: cbc.categoryBooks[cbc.category],
		index: 0,
	}
}

func main() {
	// Example books
	books := []string{"book1", "book2", "book3", "book4", "book5"}
	bookCollection := BookCollection{books: books}
	iterator := bookCollection.CreateIterator()

	fmt.Println("Books in BookCollection:")
	for iterator.HasNext() {
		book := iterator.Next()
		fmt.Println(*book)
	}

	// Example category books
	categoryBooks := CategoryBooks{
		"science": {"bookA", "bookB", "bookC"},
		"history": {"bookD", "bookE"},
	}

	// Create iterator for science category
	categoryCollection := CategoryBookCollection{
		categoryBooks: categoryBooks,
		category:      "science",
	}
	categoryIterator := categoryCollection.CreateIterator()

	fmt.Println("\nBooks in Science Category:")
	for categoryIterator.HasNext() {
		book := categoryIterator.Next()
		fmt.Println(*book)
	}

	// Create iterator for history category
	categoryCollection.category = "history"
	categoryIterator = categoryCollection.CreateIterator()

	fmt.Println("\nBooks in History Category:")
	for categoryIterator.HasNext() {
		book := categoryIterator.Next()
		fmt.Println(*book)
	}
}
