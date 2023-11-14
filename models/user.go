package models

type User struct {
	username  string `json:"username"`
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
	email     string `json:"email"`
	password  string `json:"password"`
}

// TableName use to specific table
func (b *User) TableName() string {
	return "users"
}
