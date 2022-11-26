package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jgcaceres97/goly/app/server/routes"
	"github.com/jgcaceres97/goly/app/settings"
)

func SetupAndListen() fiber.App {
	router := fiber.New()

	router.Use(logger.New())
	router.Use(cors.New(
		cors.Config{AllowMethods: "Accept, Content-Type"},
	))

	router.Get("/r/:redirect", routes.Redirect)

	goly := router.Group("/goly")

	goly.Get("/", routes.GetAllRedirects)
	goly.Get("/:id", routes.GetGoly)
	goly.Post("/", routes.CreateGoly)
	goly.Put("/", routes.UpdateGoly)
	goly.Delete("/:id", routes.DeleteGoly)

	go func() {
		addr := fmt.Sprintf(":%s", *settings.Port)

		if err := router.Listen(addr); err != nil {
			panic(err)
		}
	}()

	return *router
}
