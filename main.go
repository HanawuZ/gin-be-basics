package main

import (
	// "net/http"

	"github.com/HanawuZ/gin-be-basics/configs"
	// "github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	// "github.com/HanawuZ/gin-be-basics/repositories"
	"github.com/HanawuZ/gin-be-basics/routes"
)

//? Source: https://amitshekhar.me/blog/go-backend-clean-architecture

func main() {
	configs.DatabaseConnection()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db := configs.DB()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "hello world!",
	// 	})
	// })
	var router = app
	routes.Setup(router, db)
	// router.Run()

	app.Listen(":8080")

}
