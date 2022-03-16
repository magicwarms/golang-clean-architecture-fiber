package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// NotFoundConfig is to handle route 404 not found exception
func NotFoundConfig(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(AppResponse{
			Code:    fiber.StatusNotFound,
			Message: "Route not found",
			Data:    nil,
		})
	})
}

// GoDotEnvVariable at godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("../environments/.env.localdev")

	if err != nil {
		log.Fatalf("Error loading .env file")
		panic("Error loading .env file")
	}

	return os.Getenv(key)
}

// AppResponse is for response config show to Frontend side
type AppResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Timer will measure how long it takes before a response is returned
func Timer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// start timer
		start := time.Now()
		// next routes
		err := c.Next()
		// stop timer
		stop := time.Now()
		// Do something with response
		c.Append("Server-Timing", fmt.Sprintf("response-duration=%v", stop.Sub(start).String()))
		// return stack error if exist
		return err
	}
}

//PrettyPrint is make easier to print data result to console after querying data to db
func PrettyPrint(i interface{}) string {
	results, _ := json.MarshalIndent(i, "", "\t")
	return string(results)
}
