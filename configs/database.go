package configs

import (
	"fmt"

	"github.com/HanawuZ/gin-be-basics/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func DB() *gorm.DB {
	return database
}

func DatabaseConnection() {
	dsn := "root:123456@tcp(localhost:3306)/springboot?parseTime=true"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial)
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	database = db
	fmt.Println("Successfully connected to database....")
	fmt.Println(db)

	err = db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		fmt.Println("Error auto migrating models")
		panic(err)
	}

	// return db, err
}
