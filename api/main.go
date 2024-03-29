package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	_ "github.com/sachin-404/ZipLink/docs"
	"github.com/sachin-404/ZipLink/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault)

	app.Get("/:url", routes.ResolveURL)
	app.Post("api/v1", routes.ShortenURL)
}

//	@title		ZipLink API
//	@version	1.0
//	@description.markdown

// @BasePath	/
// @schemes	http https
func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
