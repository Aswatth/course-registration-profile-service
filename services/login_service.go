package services

import (
	"course-registration-system/profile-service/models"
	"errors"
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
			log.Fatal("Unable to create admin credentials")
		}

		obj.sql_database.db.Create(&models.Login{Email_id: "admin@univ.edu", Password: string(hashed_password), User_type: "ADMIN"})
	}

}

func (obj *LoginService) Validate(login_details models.Login) (string, error) {

	var fetched_login models.Login

	result := obj.sql_database.db.First(&fetched_login, "email_id = ?", login_details.Email_id)

	if result.RowsAffected == 0 {
		return "", errors.New("INVALID_CREDENTIALS")
	}

	err := bcrypt.CompareHashAndPassword([]byte(fetched_login.Password), []byte(login_details.Password))

	if err != nil {
		return "", errors.New("INVALID_CREDENTIALS")
	} else {
		return fetched_login.User_type, nil
	}
}

func (obj *LoginService) UpdatePassword(email_id string, new_password string, current_user_type string) error {

	var existing_login models.Login

	result := obj.sql_database.db.First(&existing_login, "email_id = ?", email_id)

	if result.Error != nil {
		return errors.New("record not found")
	}

	if existing_login.User_type != current_user_type {
		return errors.New("record not found")
	}

	//Check if new password is same as old password
	err := bcrypt.CompareHashAndPassword([]byte(existing_login.Password), []byte(new_password))

	if err == nil {
		return errors.New("new password cannot be same as old password")
	}

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(new_password), bcrypt.DefaultCost)

	if err != nil {
		return errors.New("unable to hash new password")
	}

	result = obj.sql_database.db.Exec(`UPDATE logins SET password = ? WHERE email_id = ?`, hashed_password, email_id)

	return result.Error
}
