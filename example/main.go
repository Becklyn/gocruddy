package main

import (
	"github.com/Becklyn/gocruddy"
	"github.com/Becklyn/gocruddy/example/app"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	c := app.NewContainer([]gocruddy.CrudConfig{
		&app.UserCrud{},
	})

	router := fiber.New()
	router.Use(logger.New(logger.Config{
		Output: c.GetLogger(),
		Format: "${method} ${path} [${status} - ${latency}]",
	}))

	gocruddy.RegisterCrudRoutes(router, c)
	err := router.Listen(":3000")

	if err != nil {
		c.GetLogger().ErrFatal(err)
	}
}
