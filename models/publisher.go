package models

/*
 Book * ------- 1 Publisher

 Publisher
    PublisherID (Primary Key)
    Publisher Name
    Address
    Phone Number
    Email
*/
type Publisher struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Books       []Book `gorm:"foreignKey:PublisherId"`
}
