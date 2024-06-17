package controllers

import (
	"course-registration-system/profile-service/models"
	"course-registration-system/profile-service/services"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateStudentProfile(t *testing.T) {
	load_env()

	server := setup_test_router()

	admin_profile_service := new(services.AdminProfileService)
	admin_profile_service.Init(sql_database)

	admin_profile_controller := new(AdminProfileController)
	admin_profile_controller.Init(admin_profile_service)

	new_student_profile := map[string]string{"email_id": "student@univ.edu", "password": "12345", "first_name": "test", "last_name": "student", "program_enrolled": "test"}
	new_student_profile_json, _ := json.Marshal(new_student_profile)

	server.POST("/admin/students", admin_profile_controller.CreateStudentProfile)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/admin/students", strings.NewReader(string(new_student_profile_json)))

	server.ServeHTTP(w, req)

	if w.Code == 200 {
		login_service := new(services.LoginService)
		login_service.Init(sql_database)

		result := login_service.Validate(models.Login{Email_id: "student@univ.edu", Password: "12345"})

		if result != "STUDENT" {
			t.Errorf("unable to login after successfully creating student profile")
		}

	}

	admin_profile_service.DeleteProfile("student@univ.edu")
}

func TestCreateProfessorProfile(t *testing.T) {
	load_env()

	server := setup_test_router()

	admin_profile_service := new(services.AdminProfileService)
	admin_profile_service.Init(sql_database)

	admin_profile_controller := new(AdminProfileController)
	admin_profile_controller.Init(admin_profile_service)

	new_professor_profile := map[string]string{"email_id": "professor@univ.edu", "password": "12345", "first_name": "test", "last_name": "professor", "department": "test", "designation": "professor"}
	new_professor_profile_json, _ := json.Marshal(new_professor_profile)

	server.POST("/admin/professors", admin_profile_controller.CreateProfessorProfile)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/admin/professors", strings.NewReader(string(new_professor_profile_json)))

	server.ServeHTTP(w, req)

	if w.Code == 200 {
		login_service := new(services.LoginService)
		login_service.Init(sql_database)

		result := login_service.Validate(models.Login{Email_id: "professor@univ.edu", Password: "12345"})

		if result != "PROFESSOR" {
			t.Errorf("unable to login after successfully creating professor profile")
		}

	}

	admin_profile_service.DeleteProfile("professor@univ.edu")
}

func TestDeleteStudentProfile(t *testing.T) {
	load_env()

	server := setup_test_router()

	admin_profile_service := new(services.AdminProfileService)
	admin_profile_service.Init(sql_database)

	admin_profile_controller := new(AdminProfileController)
	admin_profile_controller.Init(admin_profile_service)

	mock_login := models.Login{Email_id: "student@univ.edu", Password: "12345", User_type: "STUDENT"}
	mock_profile := models.StudentProfile{Email_id: "student@univ.edu", First_name: "test", Last_name: "student", Program_enrolled: "test"}

	admin_profile_service.CreateStudentProfile(mock_login, mock_profile)

	server.DELETE("/admin/students/:email_id", admin_profile_controller.DeleteStudentProfile)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/admin/students/student@univ.edu", nil)

	server.ServeHTTP(w, req)

	if w.Code == 200 {
		login_service := new(services.LoginService)
		login_service.Init(sql_database)

		result := login_service.Validate(models.Login{Email_id: "student@univ.edu", Password: "12345"})

		if result == "STUDENT" {
			t.Errorf("able to login after successful deletion")
		}

	}
}

func TestDeleteProfessorProfile(t *testing.T) {
	load_env()

	server := setup_test_router()

	admin_profile_service := new(services.AdminProfileService)
	admin_profile_service.Init(sql_database)

	admin_profile_controller := new(AdminProfileController)
	admin_profile_controller.Init(admin_profile_service)

	mock_login := models.Login{Email_id: "professor@univ.edu", Password: "12345", User_type: "PROFESSOR"}
	mock_profile := models.StudentProfile{Email_id: "professor@univ.edu", First_name: "test", Last_name: "professor", Program_enrolled: "test"}

	admin_profile_service.CreateStudentProfile(mock_login, mock_profile)

	server.DELETE("/admin/professors/:email_id", admin_profile_controller.DeleteProfessorProfile)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/admin/professors/professor@univ.edu", nil)

	server.ServeHTTP(w, req)

	if w.Code == 200 {
		login_service := new(services.LoginService)
		login_service.Init(sql_database)

		result := login_service.Validate(models.Login{Email_id: "professor@univ.edu", Password: "12345"})

		if result == "PROFESSOR" {
			t.Errorf("able to login after successful deletion")
		}

	}
}
