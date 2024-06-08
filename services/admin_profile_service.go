package services

import "course-registration-system/profile-service/models"

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
