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

func init_student_profile() (models.Login, models.StudentProfile, services.AdminProfileService, StudentProfileController) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sql_database := new(services.MySqlDatabase)
	sql_database.Connect(os.Getenv("MYSQL_CONNECTION_STRING"))

	admin_profile_service := new(services.AdminProfileService)
	admin_profile_service.Init(*sql_database)

	student_profile_service := new(services.StudentProfileService)
	student_profile_service.Init(*sql_database)

	student_profile_controller := new(StudentProfileController)
	student_profile_controller.Init(student_profile_service)

	//Creating mock profile
	email_id := "test_student@univ.edu"
	mock_student_login := models.Login{Email_id: email_id, Password: "12345", User_type: "STUDENT"}

	mock_student_profile := models.StudentProfile{Email_id: email_id, First_name: "test", Last_name: "student", Program_enrolled: "test_program"}

	return mock_student_login, mock_student_profile, *admin_profile_service, *student_profile_controller
}

func TestStudentUpdatePassword(t *testing.T) {
	server := setup_test_router()

	mock_student_login, mock_student_profile, admin_profile_service, student_profile_controller := init_student_profile()

	err := admin_profile_service.CreateStudentProfile(mock_student_login, mock_student_profile)

	if err == nil {
		server.PUT("students/password/:email_id", student_profile_controller.UpdatePassword)

		w := httptest.NewRecorder()

		new_password := map[string]string{"new_password": "1234"}
		new_password_data, _ := json.Marshal(new_password)

		req, _ := http.NewRequest("PUT", "/students/password/"+mock_student_login.Email_id, strings.NewReader(string(new_password_data)))

		server.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Error("Failed to update password")
		}
	}

	admin_profile_service.DeleteProfile(mock_student_login.Email_id)
}

func TestStudentUpdatePasswordFail(t *testing.T) {
	server := setup_test_router()

	mock_student_login, mock_student_profile, admin_profile_service, student_profile_controller := init_student_profile()

	err := admin_profile_service.CreateStudentProfile(mock_student_login, mock_student_profile)

	if err == nil {
		server.PUT("students/password/:email_id", student_profile_controller.UpdatePassword)

		w := httptest.NewRecorder()

		new_password := map[string]string{"new_password": "12345"}
		new_password_data, _ := json.Marshal(new_password)

		req, _ := http.NewRequest("PUT", "/students/password/"+mock_student_login.Email_id, strings.NewReader(string(new_password_data)))

		server.ServeHTTP(w, req)

		if w.Code != 400 {
			t.Error("Should fail to update password")
		}
	}

	admin_profile_service.DeleteProfile(mock_student_login.Email_id)
}
