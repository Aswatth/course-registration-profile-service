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

func init_professor_profile() (models.Login, ProfessorProfileController, services.AdminProfileService) {

	load_env()

	admin_profile_service := new(services.AdminProfileService)
	admin_profile_service.Init(sql_database)

	professor_profile_service := new(services.ProfessorProfileService)
	professor_profile_service.Init(sql_database)

	professor_profile_controller := new(ProfessorProfileController)
	professor_profile_controller.Init(professor_profile_service)

	//Creating mock profile
	email_id := "test_professor@univ.edu"
	mock_professor_login := models.Login{Email_id: email_id, Password: "12345", User_type: "PROFESSOR"}

	mock_professor_profile := models.ProfessorProfile{Email_id: email_id, First_name: "test", Last_name: "professor", Designation: "Assistant professor", Department: "CS"}

	admin_profile_service.CreateProfessorProfile(mock_professor_login, mock_professor_profile)

	return mock_professor_login, *professor_profile_controller, *admin_profile_service
}

func TestProfessorUpdatePassword(t *testing.T) {
	server := setup_test_router()

	mock_professor_login, professor_profile_controller, admin_profile_service := init_professor_profile()

	server.PUT("professors/password/:email_id", professor_profile_controller.UpdatePassword)

	w := httptest.NewRecorder()

	new_password := map[string]string{"new_password": "1234"}
	new_password_data, _ := json.Marshal(new_password)

	req, _ := http.NewRequest("PUT", "/professors/password/"+mock_professor_login.Email_id, strings.NewReader(string(new_password_data)))

	server.ServeHTTP(w, req)

	login_service := new(services.LoginService)
	login_service.Init(sql_database)

	result := login_service.Validate(models.Login{Email_id: mock_professor_login.Email_id, Password: "1234"})

	if result != "PROFESSOR" {
		t.Error("Failed to update password")
	}

	admin_profile_service.DeleteProfile(mock_professor_login.Email_id)
}

func TestProfessorUpdatePasswordFail(t *testing.T) {
	server := setup_test_router()

	mock_professor_login, professor_profile_controller, admin_profile_service := init_professor_profile()

	server.PUT("professors/password/:email_id", professor_profile_controller.UpdatePassword)

	w := httptest.NewRecorder()

	new_password := map[string]string{"new_password": mock_professor_login.Password}
	new_password_data, _ := json.Marshal(new_password)

	req, _ := http.NewRequest("PUT", "/professors/password/"+mock_professor_login.Email_id, strings.NewReader(string(new_password_data)))

	server.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Error("Should fail to update password")
	}

	admin_profile_service.DeleteProfile(mock_professor_login.Email_id)
}
