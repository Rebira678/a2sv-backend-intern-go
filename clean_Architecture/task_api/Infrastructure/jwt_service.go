package infrastructure

import (
	"errors"
	"time"

	"task-api/Domain"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService(secretKey string, issuer string) domain.JwtService {
	return &jwtService{
		secretKey: secretKey,
		issuer:    issuer,
	}
}

func (j *jwtService) GenerateToken(user *domain.User) (string, error) {
	claims := &jwtCustomClaims{
		user.ID,
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			Issuer:    j.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (j *jwtService) ValidateToken(tokenString string) (*domain.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtCustomClaims)
	if ok && token.Valid {
		return &domain.User{
			ID:       claims.ID,
			Username: claims.Username,
		}, nil
	}
	return nil, errors.New("invalid token")
}
