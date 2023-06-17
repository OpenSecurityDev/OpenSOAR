package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"main/auth"
	"main/configuration"
	"os"
)

func main() {
	configFilePath := flag.String("config", "config.yaml", "The file path of the config")
	flag.Parse()

	//Load configuration
	config, err := configuration.Load(*configFilePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1000)
	}

	//Start server
	app := fiber.New()

	//Global Middleware
	jwt := auth.MiddlewareVerifyJWT()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Create a Login route
	app.Post("/login", auth.Login)

	app.Get("/protected", jwt, auth.Protected)

	err = app.Listen(
		fmt.Sprintf("%s:%d",
			config.Configuration.Server.Address,
			config.Configuration.Server.Port))
	if err != nil {
		fmt.Printf("Failed to start server. Error: %v", err)
		os.Exit(1001)
	}

}
