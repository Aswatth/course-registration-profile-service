package services

import (
	"course-registration-system/profile-service/models"
)

type ProfessorProfileService struct {
	sql_database MySqlDatabase
}

func (obj *ProfessorProfileService) Init(db MySqlDatabase) {
	obj.sql_database = db
	obj.sql_database.db.AutoMigrate(&models.ProfessorProfile{})
}

func (obj *ProfessorProfileService) FetchProfessorProfile(email_id string) models.ProfessorProfile {
	var professor_profile models.ProfessorProfile

	obj.sql_database.db.First(&professor_profile, "email_id = ?", email_id)

	return professor_profile
}
