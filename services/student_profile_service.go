package services

import (
	"course-registration-system/profile-service/models"
)

type StudentProfileService struct {
	sql_database MySqlDatabase
}

func (obj *StudentProfileService) Init(db MySqlDatabase) {
	obj.sql_database = db
	obj.sql_database.db.AutoMigrate(&models.StudentProfile{})
	obj.sql_database.db.Exec(`ALTER TABLE student_profiles ADD CONSTRAINT fk_student_profiles FOREIGN KEY (email_id) REFERENCES logins(email_id) ON UPDATE CASCADE ON DELETE CASCADE;`)
}

func (obj *StudentProfileService) FetchStudentProfile(email_id string) models.StudentProfile {
	var student_profile models.StudentProfile

	obj.sql_database.db.First(&student_profile, "email_id = ?", email_id)

	return student_profile
}
