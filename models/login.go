package models

type Login struct {
	Email_id  string `json:"email_id" gorm:"primaryKey"`
	Password  string `json:"password"`
	User_type string `json:"user_type"`
}
