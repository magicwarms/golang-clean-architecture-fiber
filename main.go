package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"startup-backend/apps/books"
	Books "startup-backend/apps/books/route"
	"startup-backend/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	numOfCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfCores)
	// Print current process
	if fiber.IsChild() {
		fmt.Printf("[%d] CHILD\n", os.Getppid())
	} else {
		fmt.Printf("[%d] MASTER\n", os.Getppid())
	}
	enablePrefork := false
	if config.GoDotEnvVariable("APPLICATION_ENV") == "production" {
		enablePrefork = true
	}
	app := fiber.New(fiber.Config{
		Prefork: enablePrefork,
		// Enables the Server HTTP header with the given value.
		ServerHeader: "STARTUP-V1",
		// Override default error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(config.AppResponse(fiber.StatusInternalServerError, "Internal Server Error - "+err.Error(), nil))
			}
			// Return from handler
			return nil
		},
	})
	// will compress the response using gzip, deflate and brotli compression depending on the Accept-Encoding header.
	app.Use(compress.New())
	// to enable Cross-Origin Resource Sharing with various options.
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	// Logger middleware for Fiber that logs HTTP request/response details.
	file, err := os.OpenFile("./logs/app-logging.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Output:     file,
		Format:     "${pid} [${time}] | [${host} - ${ip}] | ${status} - ${latency} - ${method} | ${path} || ${error}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	// To recover from a panic thrown by any handler in the stack
	app.Use(recover.New())
	// for Fiber to let's caches be more efficient and save bandwidth,
	// as a web server does not need to resend a full response if the content has not changed.
	app.Use(etag.New())
	// Custom Timer middleware
	app.Use(config.Timer())
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")
		// Go to next middleware:
		return c.Next()
	})
	// setup initial routes
	apiV1 := app.Group("/api/v1")
	apiV1.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(config.AppResponse(http.StatusOK, "THE API IS RUNNING NOW", nil))
	})
	apiV1.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})
	DBConnection := config.InitDatabase()

	bookRepo := books.NewRepo(DBConnection)
	bookService := books.NewService(bookRepo)
	Books.BookRouter(apiV1, bookService)

	// setup not found 404 response
	config.NotFoundConfig(app)

	// running the app
	fmt.Println("⚡️ [" + config.GoDotEnvVariable("APPLICATION_ENV") + "] - " + config.GoDotEnvVariable("APP_NAME") + " IS RUNNING ON PORT - " + config.GoDotEnvVariable("APP_PORT"))
	// start listen app
	log.Fatal(app.Listen(":" + config.GoDotEnvVariable("APP_PORT")))

}
