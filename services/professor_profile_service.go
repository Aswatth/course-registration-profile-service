package services

import (
	"course-registration-system/profile-service/models"
	"fmt"
	"log"
)

type ProfessorProfileService struct {
	sql_database MySqlDatabase
}

func (obj *ProfessorProfileService) Init(db MySqlDatabase) {
	obj.sql_database = db
	obj.sql_database.db.AutoMigrate(&models.ProfessorProfile{})
}

func (obj *ProfessorProfileService) CreateProfile(professor_profile models.ProfessorProfile) {
	result := obj.sql_database.db.Create(&professor_profile)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	if result.RowsAffected > 0 {
		fmt.Println("New profile created successfully")
	}
}

func (obj *ProfessorProfileService) FetchProfessorProfile(email_id string) models.ProfessorProfile {
	var professor_profile models.ProfessorProfile

	obj.sql_database.db.First(&professor_profile, "email_id = ?", email_id)

	return professor_profile
}

func (obj *ProfessorProfileService) UpdateProfessorProfile(professor_profile models.ProfessorProfile) {
	obj.sql_database.db.Save(&professor_profile)
}

func (obj *ProfessorProfileService) DeleteProfessorProfile(email_id string) {
	var professor_profile models.ProfessorProfile

	obj.sql_database.db.Delete(&professor_profile, "email_id = ?", email_id)
}
