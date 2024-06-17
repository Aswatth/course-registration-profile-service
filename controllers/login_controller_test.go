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

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init_login() *gin.Engine {
	r := setup_test_router()

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sql_database := new(services.MySqlDatabase)
	sql_database.Connect(os.Getenv("MYSQL_CONNECTION_STRING"))

	login_service := new(services.LoginService)
	login_service.Init(*sql_database)

	login_controller := new(LoginController)
	login_controller.Init(login_service)

	r.POST("/login", login_controller.Login)

	return r
}

func init_admin_profile_service() services.AdminProfileService {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sql_database := new(services.MySqlDatabase)
	sql_database.Connect(os.Getenv("MYSQL_CONNECTION_STRING"))

	admin_profile_service := new(services.AdminProfileService)
	admin_profile_service.Init(*sql_database)

	return *admin_profile_service
}

func create_professor_profile(admin_profile_service services.AdminProfileService) models.Login {
	email_id := "professor@univ.edu"
	mock_login := &models.Login{Email_id: email_id, Password: "12345", User_type: "PROFESSOR"}
	mock_professor_profile := &models.ProfessorProfile{Email_id: email_id, First_name: "TEST", Last_name: "test", Designation: "Professor", Department: "TEST"}

	admin_profile_service.CreateProfessorProfile(*mock_login, *mock_professor_profile)

	return *mock_login
}

func create_student_profile(admin_profile_service services.AdminProfileService) models.Login {
	email_id := "student@univ.edu"
	mock_login := &models.Login{Email_id: email_id, Password: "12345", User_type: "STUDENT"}
	mock_student_profile := &models.StudentProfile{Email_id: email_id, First_name: "TEST", Last_name: "test", Program_enrolled: "TEST"}

	admin_profile_service.CreateStudentProfile(*mock_login, *mock_student_profile)

	return *mock_login
}

func run_and_get_user_type(login_data []byte) string {

	router := init_login()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(string(login_data)))

	router.ServeHTTP(w, req)

	type user_type struct {
		User_type string
	}

	var user_type_data user_type

	json.Unmarshal(w.Body.Bytes(), &user_type_data)

	return user_type_data.User_type
}

func TestLogin(t *testing.T) {

	admin_profile_service := init_admin_profile_service()

	t.Run("Admin login", func(t *testing.T) {
		login := map[string]string{"email_id": "admin@univ.edu", "password": "admin"}

		login_data, _ := json.Marshal(login)

		actual_result := run_and_get_user_type(login_data)

		if actual_result != "ADMIN" {
			t.Errorf("\nExpected:\t%s \nActual:\t%s", "ADMIN", actual_result)
		}
	})

	t.Run("Professor login", func(t *testing.T) {
		mock_login := create_professor_profile(admin_profile_service)

		login_data, _ := json.Marshal(mock_login)

		actual_result := run_and_get_user_type(login_data)

		if actual_result != "PROFESSOR" {
			t.Errorf("\nExpected:\t%s \nActual:\t%s", "PROFESSOR", actual_result)
		}

		admin_profile_service.DeleteProfile(mock_login.Email_id)
	})

	t.Run("Student login", func(t *testing.T) {
		mock_login := create_student_profile(admin_profile_service)

		login_data, _ := json.Marshal(mock_login)

		actual_result := run_and_get_user_type(login_data)

		if actual_result != "STUDENT" {
			t.Errorf("\nExpected:\t%s \nActual:\t%s", "STUDENT", actual_result)
		}

		admin_profile_service.DeleteProfile(mock_login.Email_id)
	})
}
