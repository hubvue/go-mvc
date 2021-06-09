package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"server-test/src/config"
	"server-test/src/database"
	"server-test/src/middleware"
	"server-test/src/router"
)

func main() {
	defer database.CloseMysql()
	config.NewConfig("dev")
	app := fiber.New()

	middleware.Init(app)
	router.Init(app)

	log.Fatal(app.Listen(config.Conf.App.Host + ":" + config.Conf.App.Port))
}
