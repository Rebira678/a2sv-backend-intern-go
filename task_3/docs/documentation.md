# Library Management System

## Overview
This is a console-based library management application built in Go. It allows users to manage books, members, and library transactions like borrowing and returning books.

## Project Structure
- **models/**: Contains the basic data structures (`Book` and `Member`).
- **services/**: Contains the core business logic (`Library` struct), implementing the `LibraryManager` interface. Maps are used here for fast lookup of books and members by ID.
- **controllers/**: Handles standard input/output from the console, mapping user commands to service actions.
- **main.go**: The entry point that initializes the dependencies and starts the controller loop.

## How to Run
Ensure you have Go installed. Open a terminal in the root directory and run:
`go run main.go`