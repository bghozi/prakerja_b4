package models

type User struct {
	Id       int    `gorm:"primaryKey autoIncrement" json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
