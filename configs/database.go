package configs

import (
	"fmt"
	"time"

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
		&models.Author{},
		&models.Book{},
		&models.Publisher{},
	)
	if err != nil {
		fmt.Println("Error auto migrating models")
		panic(err)
	}

	// AddSampleData()
}

// Add sample data to database
func AddSampleData() {

	var authors = []*models.Author{
		{Name: "John Smith", Dob: time.Date(1980, 5, 15, 0, 0, 0, 0, time.UTC)},
		{Name: "Alice Johnson", Dob: time.Date(1975, 12, 10, 0, 0, 0, 0, time.UTC)},
		{Name: "Michael Davis", Dob: time.Date(1992, 8, 22, 0, 0, 0, 0, time.UTC)},
		{Name: "Emily Wilson", Dob: time.Date(1988, 4, 30, 0, 0, 0, 0, time.UTC)},
		{Name: "David Roberts", Dob: time.Date(1970, 11, 5, 0, 0, 0, 0, time.UTC)},
		{Name: "Sarah Brown", Dob: time.Date(1995, 2, 18, 0, 0, 0, 0, time.UTC)},
		{Name: "James Miller", Dob: time.Date(1982, 9, 25, 0, 0, 0, 0, time.UTC)},
		{Name: "Olivia Taylor", Dob: time.Date(1978, 7, 12, 0, 0, 0, 0, time.UTC)},
		{Name: "Daniel Anderson", Dob: time.Date(1990, 3, 3, 0, 0, 0, 0, time.UTC)},
		{Name: "Sophia White", Dob: time.Date(1986, 6, 28, 0, 0, 0, 0, time.UTC)},
	}
	database.Create(&authors)

	var publishers = []*models.Publisher{
		{Name: "Penguin Random House", Address: "1745 Broadway, New York, NY 10019", PhoneNumber: "212-782-9000", Email: "penguin.d@gmail.com"},
		{Name: "HarperCollins", Address: "195 Broadway, New York, NY 10007", PhoneNumber: "212-207-7000", Email: "CollinsBook@gmail.com"},
	}
	database.Create(&publishers)

	var book1 = models.Book{
		Isbn:            "9781234567890",
		Title:           "Book Title 1",
		Genre:           "Fiction",
		PublicationYear: time.Date(2021, 1, 15, 0, 0, 0, 0, time.UTC),
		CopiesAvailable: 50,
		Price:           19.99,
		Publisher:       *publishers[0],
		Authors:         []models.Author{*authors[0], *authors[1]},
	}

	var book2 = models.Book{
		Isbn:            "9782345678901",
		Title:           "Book Title 2",
		Genre:           "Mystery",
		PublicationYear: time.Date(2019, 7, 20, 0, 0, 0, 0, time.UTC),
		CopiesAvailable: 30,
		Price:           15.99,
		Publisher:       *publishers[1],
		Authors:         []models.Author{*authors[2]},
	}

	var book3 = models.Book{
		Isbn:            "9783456789012",
		Title:           "Book Title 3",
		Genre:           "Thriller",
		PublicationYear: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
		CopiesAvailable: 40,
		Price:           17.99,
		Publisher:       *publishers[0],
		Authors:         []models.Author{*authors[3], *authors[4]},
	}

	var books = []*models.Book{&book1, &book2, &book3}

	database.Create(books)
}
