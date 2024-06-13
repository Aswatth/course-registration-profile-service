package models

type Login struct {
	Email_id  string `gorm:"primaryKey"`
	Password  string
	User_type string
}
