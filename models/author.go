package models

import "time"

// Languages []Language `gorm:"many2many:user_languages;"`

type Author struct {
	Id    uint      `json:"id" gorm:"primaryKey"`
	Name  string    `json:"name"`
	Dob   time.Time `json:"dob"`
	Books []Book    `json:"-" gorm:"many2many:book_authors;"`
}
