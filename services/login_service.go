package services

import (
	"course-registration-system/profile-service/models"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	sql_database MySqlDatabase
}

func (obj *LoginService) Init(db MySqlDatabase) {
	obj.sql_database = db
	obj.sql_database.db.AutoMigrate(&models.Login{})
}

func (obj *LoginService) Validate(login_details models.Login) bool {

	var fetched_login models.Login

	obj.sql_database.db.First(&fetched_login, "username = ?", login_details.Email_id)

	err := bcrypt.CompareHashAndPassword([]byte(login_details.Password), []byte(fetched_login.Password))

	if err != nil {
		return false
	} else {
		return true
	}
}
