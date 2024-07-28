package main

import (
        "log"

        "github.com/gofiber/fiber/v2"
        "github.com/gofiber/fiber/v2/middleware/logger"
        "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
        app := fiber.New(fiber.Config{
                Prefork:       true,
                CaseSensitive: true,
        })

        // Middleware
        app.Use(recover.New())
        app.Use(logger.New())

        // Static files
        app.Static("/", "./public")

        // Routes
        app.Get("/", func(c *fiber.Ctx) error {
                return c.Render("index", fiber.Map{})
        })

        app.Get("/about", func(c *fiber.Ctx) error {
                return c.Render("about", fiber.Map{})
        })

        log.Fatal(app.Listen(":3000"))
}
