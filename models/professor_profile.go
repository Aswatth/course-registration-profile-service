package models

type ProfessorProfile struct {
	Email_id    string `gorm:"primaryKey"`
	First_name  string
	Last_name   string
	Designation string
	Department  string
}

func (obj *ProfessorProfile) CreateProfile(email_id string, password string, first_name string, last_name string, designation string, department string) {
	obj.First_name = first_name
	obj.Last_name = last_name
	obj.Designation = designation
	obj.Department = department
}
