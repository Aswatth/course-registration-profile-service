package models

type ProfessorProfile struct {
	Email_id    string `json:"email_id" gorm:"primaryKey"`
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Designation string `json:"designation"`
	Department  string `json:"department"`
}

func (obj *ProfessorProfile) CreateProfile(email_id string, password string, first_name string, last_name string, designation string, department string) {
	obj.First_name = first_name
	obj.Last_name = last_name
	obj.Designation = designation
	obj.Department = department
}
