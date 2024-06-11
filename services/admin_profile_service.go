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
	obj.sql_database.db.AutoMigrate(&models.AdminProfile{})
}

func (obj *AdminProfileService) Validate(admin_profile models.AdminProfile) bool {

	var fetched_admin_profile models.AdminProfile

	obj.sql_database.db.First(&fetched_admin_profile, "username = ?", admin_profile.Username)

	return fetched_admin_profile.Password == admin_profile.Password
}

func (obj *AdminProfileService) CreateStudentProfile(student_profile models.StudentProfile) error {
	result := obj.sql_database.db.Create(&student_profile)

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

func (obj *AdminProfileService) DeleteStudentProfile(email_id string) error {
	var student_profile models.StudentProfile

	result := obj.sql_database.db.Delete(&student_profile, "email_id = ?", email_id)

	return result.Error
}

func (obj *AdminProfileService) CreateProfessorProfile(professor_profile models.ProfessorProfile) error {
	result := obj.sql_database.db.Create(&professor_profile)

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

func (obj *AdminProfileService) DeleteProfessorProfile(email_id string) error {
	var professor_profile models.ProfessorProfile

	result := obj.sql_database.db.Delete(&professor_profile, "email_id = ?", email_id)

	return result.Error
}