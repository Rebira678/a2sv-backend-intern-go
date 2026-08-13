package data

import (
	"errors"
	"task_manager/models"
	"golang.org/x/crypto/bcrypt"
)

var users = make(map[string]models.User)

func CreateUser(user models.User) error {
	if _, exists := users[user.Username]; exists {
		return errors.New("username already exists")
	}
	
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	
	user.Password = string(hashedPassword)
	users[user.Username] = user
	return nil
}

func AuthenticateUser(username, password string) (*models.User, error) {
	user, exists := users[username]
	if !exists {
		return nil, errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
