package main

import (
	"log"
	"net/http"

	"github.com/HanawuZ/gin-be-basics/configs"
	"github.com/gin-gonic/gin"

	"github.com/HanawuZ/gin-be-basics/repositories"
)

func main() {
	db, err := configs.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world!",
		})
	})

	userRepo := &repositories.UserRepository{}
	userRepo.DB = db

	// userUseCase := usecases.NewUserUseCase(userRepo)
	// controller := controllers.UserController{UserUsecase: userUseCase}

	/*
			coinRepo := repository.NewCoinRepository(DB)
		coinUseCase := usecases.NewCoinUseCase(coinRepo)
		coinController := controllers.NewCoinController(coinUseCase)

		r.GET("coin", coinController.GetAllCoins)
		r.GET("coin/:id", coinController.GetACoin)
		r.POST("coin", coinController.CreateACoin)
		r.PATCH("coin/:id", coinController.UpdateACoin)
	*/

	r.Run()
}
