package services

import (
	"errors"
	"library_management/models"
)

// library management outlines the rules our library must follow
type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(BookID int)
	BorrowBook(BookID, memberID int)
	ReturnBook(BookID, memberID int)
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

// Library is the actual implementation of LibraryManager
type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

// NewLibrary is the helper function to create a new library and initialize the maps
func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

// AddBook adds a book to the library map
func (l *Library) AddBook(book models.Book) {
	book.Status = "Available" //Make sure new books are available
	l.Books[book.ID] = book
}

// RemoveBook remove books using their id
func (l *Library) RemoveBook(BookID int) {
	delete(l.Books, BookID)
}

// BorrowBook lets a member take a book if it is available
func (l *Library) BorrowBook(BookID, memberID int) error {
	book, bookExists := l.Books[BookID]
	if !bookExists {
		return errors.New("Book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("Book is already Borrowed")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return errors.New("Member not found")
	}

	//check the book status and update the map
	book.Status = "Borrowed"
	l.Books[BookID] = book

	//Add book to member's borrowed list
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member

	return nil
}

// ReturnBook lets a member give a book back
func (l *Library) ReturnBook(BookID, memberID int) error {
	book, bookExists := l.Books[BookID]
	if !bookExists {
		return errors.New("Book not found")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return errors.New("Member not Found")
	}

	//check if the member actual has this book
	foundIndex := -1
	for i, b := range member.BorrowedBooks {
		if b.ID == BookID {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return errors.New("this member did not borrow this book")
	}

	// Remove the book from the member's slice
	member.BorrowedBooks = append(member.BorrowedBooks[:foundIndex], member.BorrowedBooks[foundIndex+1:]...)
	l.Members[memberID] = member

	// Update the book's status
	book.Status = "Available"
	l.Books[BookID] = book

	return nil
}

// ListAvailableBooks returns all books with "Available" status
func (l *Library) ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, book := range l.Books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}

// ListBorrowedBooks returns books held by a specific member
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.Members[memberID]
	if !exists {
		return []models.Book{}
	}
	return member.BorrowedBooks
}

// Helper method added to let us register members through the console easily
func (l *Library) AddMember(member models.Member) {
	l.Members[member.ID] = member
}
