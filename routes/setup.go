package routes

import (
	// "github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// func Setup(router *gin.Engine, db *gorm.DB) {
// 	UserRoute(router, db)
// 	AuthorRoutes(router, db)
// 	BookRoutes(router, db)
// }

func Setup(router *fiber.App, db *gorm.DB) {
	UserRoute(router, db)
	AuthorRoutes(router, db)
	BookRoutes(router, db)
}
