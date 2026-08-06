package models

// Book represent a book in the library
type Book struct {
	ID     int
	Title  string
	Author string
	Status string // "Available" or "Borrowed"
}
