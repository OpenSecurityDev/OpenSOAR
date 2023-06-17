package auth

import (
	"crypto/rand"
	"encoding/base64"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var secretJWT = ""

func GenerateKey(length int) error {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return err
	}
	secretJWT = base64.URLEncoding.EncodeToString(key)

	return nil
}

// Middleware JWT function
func MiddlewareVerifyJWT() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secretJWT)},
	})
}
