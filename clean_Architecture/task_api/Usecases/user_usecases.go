package usecases

import (
	"errors"

	"task-api/Domain"
)

type userUsecase struct {
	userRepository  domain.UserRepository
	passwordService domain.PasswordService
	jwtService      domain.JwtService
}

func NewUserUsecase(userRepo domain.UserRepository, passService domain.PasswordService, jwtSvc domain.JwtService) domain.UserUsecase {
	return &userUsecase{
		userRepository:  userRepo,
		passwordService: passService,
		jwtService:      jwtSvc,
	}
}

func (u *userUsecase) Register(user *domain.User) error {
	existingUser, _ := u.userRepository.GetByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := u.passwordService.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return u.userRepository.Create(user)
}

func (u *userUsecase) Login(username, password string) (string, error) {
	user, err := u.userRepository.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = u.passwordService.ComparePassword(user.Password, password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := u.jwtService.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
