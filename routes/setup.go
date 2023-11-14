package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(router *gin.Engine, db *gorm.DB) {
	UserRoute(router, db)
}
