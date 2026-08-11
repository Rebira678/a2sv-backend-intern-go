package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Represents the JSON body expected from the client.
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`    // Must be a valid email.
	Password string `json:"password" binding:"required,min=8"` // At least 8 characters.
}

// Handles POST /register requests.
func registerHandler(c *gin.Context) {

	// Create an empty request object.
	var req RegisterRequest

	// Read JSON from the request body and validate it.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Convert the plain-text password into a secure bcrypt hash.
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		12,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to hash password",
		})
		return
	}

	// Build a new user object.
	user := &User{
		Email:    req.Email,
		Password: string(hashedPassword), // Store the hash, never the original password.
		Role:     "user",
	}

	// Save the user using the application's storage layer.
	if err := userStore.Create(user); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Remove the password hash before sending the response.
	user.Password = ""

	c.JSON(http.StatusCreated, gin.H{
		"message": "registration successful",
		"user":    user,
	})
}

// Read the JWT secret from the environment.
// If it doesn't exist, use a development fallback.
// In production, always set JWT_SECRET.
var jwtSecret = []byte(getEnv("JWT_SECRET", "development-secret"))

// Custom JWT claims carried inside the token.
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`

	// Standard JWT fields such as exp, iat, iss.
	jwt.RegisteredClaims
}

// Returns an environment variable if present,
// otherwise returns the supplied default value.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// Generates a signed JWT for an authenticated user.
func generateToken(user *User) (string, error) {

	// Token expires after 24 hours.
	expiration := time.Now().Add(24 * time.Hour)

	claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-auth-service",
		},
	}

	// Create a token using the HS256 signing algorithm.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Convert the token into a signed string.
	return token.SignedString(jwtSecret)
}

// Represents the JSON body expected during login.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Handles POST /login.
func loginHandler(c *gin.Context) {

	var req LoginRequest

	// Read and validate the JSON request.
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Find the user by email.
	user, err := userStore.GetByEmail(req.Email)

	if err != nil {

		// Do not reveal whether the email exists.
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})

		return
	}

	// Compare the stored bcrypt hash with the password
	// supplied by the client.
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})

		return
	}

	// Create a JWT after successful authentication.
	token, err := generateToken(user)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})

		return
	}

	// Return the JWT to the client.
	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"type":       "Bearer",
		"expires_in": 86400,
	})
}
