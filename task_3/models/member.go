package models

// Member represent a person who can borrow books
type Member struct {
	ID            int
	name          string
	BorrowedBooks []Book
}
