package controllers

import (
	"course-registration-system/profile-service/models"
	"course-registration-system/profile-service/services"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func init_professor_profile() (models.Login, models.ProfessorProfile, services.AdminProfileService, ProfessorProfileController) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sql_database := new(services.MySqlDatabase)
	sql_database.Connect(os.Getenv("MYSQL_CONNECTION_STRING"))

	admin_profile_service := new(services.AdminProfileService)
	admin_profile_service.Init(*sql_database)

	professor_profile_service := new(services.ProfessorProfileService)
	professor_profile_service.Init(*sql_database)

	professor_profile_controller := new(ProfessorProfileController)
	professor_profile_controller.Init(professor_profile_service)

	//Creating mock profile
	email_id := "test_professor@univ.edu"
	mock_professor_login := models.Login{Email_id: email_id, Password: "12345", User_type: "STUDENT"}

	mock_professor_profile := models.ProfessorProfile{Email_id: email_id, First_name: "test", Last_name: "professor", Designation: "Assistant professor", Department: "CS"}

	return mock_professor_login, mock_professor_profile, *admin_profile_service, *professor_profile_controller
}

func TestProfessorUpdatePassword(t *testing.T) {
	server := setup_test_router()

	mock_professor_login, mock_professor_profile, admin_profile_service, professor_profile_controller := init_professor_profile()

	err := admin_profile_service.CreateProfessorProfile(mock_professor_login, mock_professor_profile)

	if err == nil {
		server.PUT("professors/password/:email_id", professor_profile_controller.UpdatePassword)

		w := httptest.NewRecorder()

		new_password := map[string]string{"new_password": "1234"}
		new_password_data, _ := json.Marshal(new_password)

		req, _ := http.NewRequest("PUT", "/professors/password/"+mock_professor_login.Email_id, strings.NewReader(string(new_password_data)))

		server.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Error("Failed to update password")
		}
	}

	admin_profile_service.DeleteProfile(mock_professor_login.Email_id)
}

func TestProfessorUpdatePasswordFail(t *testing.T) {
	server := setup_test_router()

	mock_professor_login, mock_professor_profile, admin_profile_service, professor_profile_controller := init_professor_profile()

	err := admin_profile_service.CreateProfessorProfile(mock_professor_login, mock_professor_profile)

	if err == nil {
		server.PUT("professors/password/:email_id", professor_profile_controller.UpdatePassword)

		w := httptest.NewRecorder()

		new_password := map[string]string{"new_password": "12345"}
		new_password_data, _ := json.Marshal(new_password)

		req, _ := http.NewRequest("PUT", "/professors/password/"+mock_professor_login.Email_id, strings.NewReader(string(new_password_data)))

		server.ServeHTTP(w, req)

		if w.Code != 400 {
			t.Error("Should fail to update password")
		}
	}

	admin_profile_service.DeleteProfile(mock_professor_login.Email_id)
}
