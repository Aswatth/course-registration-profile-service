package services

import (
	"course-registration-system/profile-service/models"
	"fmt"
	"log"
)

type StudentProfileService struct {
	sql_database MySqlDatabase
}

func (obj *StudentProfileService) Init(db MySqlDatabase) {
	obj.sql_database = db
	obj.sql_database.db.AutoMigrate(&models.StudentProfile{})
}

func (obj *StudentProfileService) CreateProfile(student_profile models.StudentProfile) {
	result := obj.sql_database.db.Create(&student_profile)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Println("New profile created successfully")
	}
}

func (obj *StudentProfileService) FetchStudentProfile(email_id string) models.StudentProfile {
	var student_profile models.StudentProfile

	obj.sql_database.db.First(&student_profile, "email_id = ?", email_id)

	return student_profile
}

func (obj *StudentProfileService) UpdateStudentProfile(student_profile models.StudentProfile) {
	obj.sql_database.db.Save(&student_profile)
}

func (obj *StudentProfileService) DeleteStudentProfile(email_id string) {
	var student_profile models.StudentProfile

	obj.sql_database.db.Delete(&student_profile, "email_id = ?", email_id)
}
