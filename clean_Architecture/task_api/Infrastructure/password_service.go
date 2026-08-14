package infrastructure

import (
	"task-api/Domain"

	"golang.org/x/crypto/bcrypt"
)

type passwordService struct{}

func NewPasswordService() domain.PasswordService {
	return &passwordService{}
}

func (p *passwordService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (p *passwordService) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
