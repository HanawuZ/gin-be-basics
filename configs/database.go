package configs

import (
	"github.com/HanawuZ/gin-be-basics/models"
	"gorm.io/gorm"

	// Import mysql
	"fmt"

	"gorm.io/driver/mysql"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

// type DBConfig struct {
// 	user     string
// 	password string
// 	host     string
// 	name     string
// 	port     string
// }

func DatabaseConnection() (*gorm.DB, error) {
	dsn := "root:123456@tcp(localhost:3306)/springboot?parseTime=true"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial)
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	fmt.Println("Successfully connected to database....")
	fmt.Println(db)

	err = db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		fmt.Println("Error auto migrating models")
		panic(err)
	}

	return db, err
}
