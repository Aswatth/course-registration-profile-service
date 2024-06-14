package services

import (
	"course-registration-system/profile-service/models"
	"log"
)

type AdminProfileService struct {
	sql_database MySqlDatabase
}

func (obj *AdminProfileService) Init(db MySqlDatabase) {
	obj.sql_database = db
	// obj.sql_database.db.AutoMigrate(&models.AdminProfile{})
}

func (obj *AdminProfileService) CreateStudentProfile(login_data models.Login, student_profile models.StudentProfile) error {
	result := obj.sql_database.db.Create(&login_data)

	if result.Error == nil {
		result = obj.sql_database.db.Create(&student_profile)
	}

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return result.Error
}

func (obj *AdminProfileService) FetchStudentProfile(email_id string) (models.StudentProfile, error) {
	var student_profile models.StudentProfile

	result := obj.sql_database.db.First(&student_profile, "email_id = ?", email_id)

	return student_profile, result.Error
}

func (obj *AdminProfileService) UpdateStudentProfile(email_id string, student_profile models.StudentProfile) error {
	result := obj.sql_database.db.Model(&models.StudentProfile{}).Where("email_id = ?", email_id).Updates(student_profile)

	return result.Error
}

func (obj *AdminProfileService) CreateProfessorProfile(login_data models.Login, professor_profile models.ProfessorProfile) error {
	result := obj.sql_database.db.Create(&login_data)

	if result.Error == nil {
		result = obj.sql_database.db.Create(&professor_profile)
	}

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return result.Error
}

func (obj *AdminProfileService) FetchProfessorProfile(email_id string) (models.ProfessorProfile, error) {
	var professor_profile models.ProfessorProfile

	result := obj.sql_database.db.First(&professor_profile, "email_id = ?", email_id)

	return professor_profile, result.Error
}

func (obj *AdminProfileService) UpdateProfessorProfile(email_id string, professor_profile models.ProfessorProfile) error {

	result := obj.sql_database.db.Model(&models.ProfessorProfile{}).Where("email_id = ?", email_id).Updates(professor_profile)

	return result.Error
}

func (obj *AdminProfileService) DeleteProfile(email_id string) error {
	result := obj.sql_database.db.Delete(models.Login{}, "email_id = ?", email_id)

	return result.Error
}

func (obj *AdminProfileService) UpdatePassword(email_id string, new_password string) error {

	login_service := new(LoginService)
	login_service.Init(obj.sql_database)

	err := login_service.UpdatePassword(email_id, new_password)

	return err
}
