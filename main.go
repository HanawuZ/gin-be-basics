package main

import (
	"log"
	"net/http"

	"github.com/HanawuZ/gin-be-basics/configs"
	"github.com/gin-gonic/gin"

	// "github.com/HanawuZ/gin-be-basics/repositories"
	"github.com/HanawuZ/gin-be-basics/routes"
)

func main() {
	db, err := configs.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world!",
		})
	})
	routes.Setup(router, db)
	router.Run()
}
