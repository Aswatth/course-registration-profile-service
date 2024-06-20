package models

type StudentProfile struct {
	Email_id         string `json:"email_id" gorm:"primaryKey"`
	First_name       string `json:"first_name"`
	Last_name        string `json:"last_name"`
	Program_enrolled string `json:"program_enrolled"`
}
