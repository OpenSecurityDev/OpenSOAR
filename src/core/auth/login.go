package auth

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v5"
	"time"
)

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	ID       int
	Username string
	AuthKey  int
}

// Temp Function to verify creds
func verifyCredentials(email, password string) (*User, error) {
	if email == "admin@admin.com" && password == "pass" {
		return &User{
			ID:       1,
			Username: "admin",
			AuthKey:  420,
		}, nil
	}
	return nil, errors.New("user not found")
}

func Login(c *fiber.Ctx) error {
	// Extract the credentials from the request body
	loginRequest := new(LoginBody)
	err := c.BodyParser(loginRequest)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Verify user login
	user, err := verifyCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"ID":       user.ID,
		"username": user.Username,
		"AuthKey":  user.AuthKey,
		"exp":      time.Now().Add(day * 1).Unix(),
	}

	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secretJWT))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return the token
	return c.JSON(LoginResponse{
		Token: t,
	})
}

// Protected route
func Protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	tokenString := c.Locals("user").(*jtoken.Token)
	claims := &User{}
	token, _, err := new(jtoken.Parser).ParseUnverified(tokenString, claims)

	username := claims["username"].(string)
	authKey := claims["AuthKey"].(int)
	return c.SendString(fmt.Sprintf("Hello, %s. Your permissions are %d", username, authKey))
}
