package services

import (
	"course-registration-system/profile-service/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	sql_database MySqlDatabase
}

func (obj *LoginService) Init(db MySqlDatabase) {
	obj.sql_database = db

	obj.sql_database.db.AutoMigrate(&models.Login{})

	//Check if already exists
	var admin_login models.Login
	obj.sql_database.db.First(&admin_login, "email_id = ?", "admin@univ.edu")

	//Create if not already exists
	if admin_login.Email_id == "" {
		hashed_password, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)

		if err != nil {
			log.Fatal("Unable to create adming credentials")
		}

		obj.sql_database.db.Create(&models.Login{Email_id: "admin@univ.edu", Password: string(hashed_password)})
	}

}

func (obj *LoginService) Validate(login_details models.Login) bool {

	var fetched_login models.Login

	obj.sql_database.db.First(&fetched_login, "email_id = ?", login_details.Email_id)

	err := bcrypt.CompareHashAndPassword([]byte(fetched_login.Password), []byte(login_details.Password))

	if err != nil {
		return false
	} else {
		return true
	}
}
