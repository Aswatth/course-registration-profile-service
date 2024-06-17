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

func init_student_profile() (models.Login, StudentProfileController, services.AdminProfileService) {

	load_env()

	admin_profile_service := new(services.AdminProfileService)
	admin_profile_service.Init(sql_database)

	student_profile_service := new(services.StudentProfileService)
	student_profile_service.Init(sql_database)

	student_profile_controller := new(StudentProfileController)
	student_profile_controller.Init(student_profile_service)

	//Creating mock profile
	email_id := "test_student@univ.edu"
	mock_student_login := models.Login{Email_id: email_id, Password: "12345", User_type: "STUDENT"}

	mock_student_profile := models.StudentProfile{Email_id: email_id, First_name: "test", Last_name: "student", Program_enrolled: "test_program"}

	admin_profile_service.CreateStudentProfile(mock_student_login, mock_student_profile)

	return mock_student_login, *student_profile_controller, *admin_profile_service
}

func TestStudentUpdatePassword(t *testing.T) {
	server := setup_test_router()

	mock_student_login, student_profile_controller, admin_profile_service := init_student_profile()

	server.PUT("students/password/:email_id", student_profile_controller.UpdatePassword)

	w := httptest.NewRecorder()

	new_password := map[string]string{"new_password": "1234"}
	new_password_data, _ := json.Marshal(new_password)

	req, _ := http.NewRequest("PUT", "/students/password/"+mock_student_login.Email_id, strings.NewReader(string(new_password_data)))

	server.ServeHTTP(w, req)

	login_service := new(services.LoginService)
	login_service.Init(sql_database)

	result := login_service.Validate(models.Login{Email_id: mock_student_login.Email_id, Password: "1234"})

	if result != "STUDENT" {
		t.Error("Failed to update password")
	}

	admin_profile_service.DeleteProfile(mock_student_login.Email_id)
}

func TestStudentUpdatePasswordFail(t *testing.T) {
	server := setup_test_router()

	mock_student_login, student_profile_controller, admin_profile_service := init_student_profile()

	server.PUT("students/password/:email_id", student_profile_controller.UpdatePassword)

	w := httptest.NewRecorder()

	new_password := map[string]string{"new_password": "12345"}
	new_password_data, _ := json.Marshal(new_password)

	req, _ := http.NewRequest("PUT", "/students/password/"+mock_student_login.Email_id, strings.NewReader(string(new_password_data)))

	server.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Error("Should fail to update password")
	}

	admin_profile_service.DeleteProfile(mock_student_login.Email_id)
}
