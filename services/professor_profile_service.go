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
	obj.sql_database.db.Exec(`ALTER TABLE professor_profiles ADD CONSTRAINT fk_professor_profiles FOREIGN KEY (email_id) REFERENCES logins(email_id) ON UPDATE CASCADE ON DELETE CASCADE;`)
}

func (obj *ProfessorProfileService) FetchProfessorProfile(email_id string) models.ProfessorProfile {
	var professor_profile models.ProfessorProfile

	obj.sql_database.db.First(&professor_profile, "email_id = ?", email_id)

	return professor_profile
}
