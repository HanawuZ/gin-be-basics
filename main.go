package main

import (
	"net/http"

	"github.com/HanawuZ/gin-be-basics/configs"
	"github.com/gin-gonic/gin"

	// "github.com/HanawuZ/gin-be-basics/repositories"
	"github.com/HanawuZ/gin-be-basics/routes"
)

//? Source: https://amitshekhar.me/blog/go-backend-clean-architecture

func main() {
	configs.DatabaseConnection()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	router := gin.Default()
	db := configs.DB()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world!",
		})
	})
	routes.Setup(router, db)
	router.Run()
}
