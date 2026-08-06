package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	Service *services.Library // The controller needs access to the service (brain)
}

// Run starts the console menu
func (c *LibraryController) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Library Management System ---")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove an existing book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List all available books")
		fmt.Println("6. List all borrowed books by a member")
		fmt.Println("7. Add a new member (Required to borrow books)")
		fmt.Println("8. Exit")
		fmt.Print("Choose an option: ")

		optionStr, _ := reader.ReadString('\n')
		optionStr = strings.TrimSpace(optionStr)
		option, err := strconv.Atoi(optionStr)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch option {
		case 1:
			fmt.Print("Enter Book ID: ")
			id := readInt(reader)
			fmt.Print("Enter Book Title: ")
			title := readString(reader)
			fmt.Print("Enter Book Author: ")
			author := readString(reader)

			book := models.Book{ID: id, Title: title, Author: author}
			c.Service.AddBook(book)
			fmt.Println("Book added successfully!")

		case 2:
			fmt.Print("Enter Book ID to remove: ")
			id := readInt(reader)
			c.Service.RemoveBook(id)
			fmt.Println("Book removed successfully!")

		case 3:
			fmt.Print("Enter Book ID to borrow: ")
			bookID := readInt(reader)
			fmt.Print("Enter Member ID: ")
			memberID := readInt(reader)
			err := c.Service.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed successfully!")
			}

		case 4:
			fmt.Print("Enter Book ID to return: ")
			bookID := readInt(reader)
			fmt.Print("Enter Member ID: ")
			memberID := readInt(reader)
			err := c.Service.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned successfully!")
			}

		case 5:
			books := c.Service.ListAvailableBooks()
			fmt.Println("\n--- Available Books ---")
			if len(books) == 0 {
				fmt.Println("No books available.")
			}
			for _, b := range books {
				fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
			}

		case 6:
			fmt.Print("Enter Member ID: ")
			memberID := readInt(reader)
			books := c.Service.ListBorrowedBooks(memberID)
			fmt.Printf("\n--- Books Borrowed by Member %d ---\n", memberID)
			if len(books) == 0 {
				fmt.Println("No books borrowed.")
			}
			for _, b := range books {
				fmt.Printf("ID: %d | Title: %s\n", b.ID, b.Title)
			}

		case 7:
			fmt.Print("Enter Member ID: ")
			id := readInt(reader)
			fmt.Print("Enter Member Name: ")
			_ = readString(reader)

			member := models.Member{ID: id}
			c.Service.AddMember(member)
			fmt.Println("Member added successfully!")

		case 8:
			fmt.Println("Exiting system. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice, please select 1-8.")
		}
	}
}

// Helper function to easily read strings with spaces (like Book Titles)
func readString(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Helper function to easily read numbers
func readInt(reader *bufio.Reader) int {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)
	return num
}
