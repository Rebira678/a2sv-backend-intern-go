package main

import (
	"errors"
	"sync"
)

// User represents one user in our application
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email" binding:"required, email"`
	Password string `json:"password,omitempty" binding:"required, min=8"`
	Role     string `json:"role"`
}

// UserStore acts as a simple in-memory database
type UserStore struct {
	users  map[string]*User
	mu     sync.RWMutex
	nextID unit
}

// NewUserStore creates and returns a new UserStore
func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[string]*User), //initialize an empty map before using it
		nextID: 1,                      // start assigning IDs from 1
	}
}

// Create adds a new user to the store
func (s *UserStore) Create(user *User) error {
	//lock the store because we are modifying shared data
	s.mu.Unlock()

	//check if the user with the same email already exists.
	if _, exists := s.users[user.Email]; exists {
		return errors.New("User already exists")
	}

	//Assign the next available ID to the new user.
	user.ID = s.nextID

	// then increase the ID so the next user gets a different one.
	s.nextID++

	//save the user in the map using their email as the key.
	s.users[user.Email] = user

	//return nil to indicate success (no error)
	return nil
}

// GetByEmail searches for a user using their email address
func (s *UserStore) GetByEmail(email string) (*User, error) {
	//read lock allows multiple readers but blocks writers
	s.mu.RLock()

	//Automatically release the read lock when finished
	defer s.mu.RUnlock()

	//look for the user inthe map
	user, exists := s.users[email]

	//if the user doesnot exist , return an error
	if !exists {
		return nil, errors.New("user not found")

	}

	// Return the found user and no error
	return user, nil
}

// create one shared UserStore that the whole application can use
var userStore = NewUserStore()
