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
		return c.Status(fiber.StatusNotFound).JSON(AppResponse(nil))
	})
}

// GoDotEnvVariable at godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {

	// // load .env file
	env := os.Getenv("APPLICATION_ENV")
	if env == "" {
		env = "localdev"
	}

	err := godotenv.Load("./environments/.env." + env)
	if err != nil {
		log.Fatalf("Error loading .env file" + env)
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}

// AppResponse is for response config show to Frontend side
func AppResponse(data interface{}) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
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

// ErrorResponse is for response error
func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}
