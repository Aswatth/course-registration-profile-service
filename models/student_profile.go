package models

type StudentProfile struct {
	Email_id         string `gorm:"primaryKey"`
	First_name       string
	Last_name        string
	Program_enrolled string
}
