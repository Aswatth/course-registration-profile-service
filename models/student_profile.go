package models

type StudentProfile struct {
	Email_id         string `gorm:"primaryKey"`
	Password         string
	First_name       string
	Last_name        string
	Program_enrolled string
}

func (obj *StudentProfile) CreateProfile(email_id string, password string, first_name string, last_name string, program_enrolled string) {
	obj.Email_id = email_id
	obj.Password = password
	obj.First_name = first_name
	obj.Last_name = last_name
	obj.Program_enrolled = program_enrolled
}
