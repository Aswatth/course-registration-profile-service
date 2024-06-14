package controllers

import (
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

func setup_router() *gin.Engine {
	r := gin.Default()

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

func run_and_get_user_type(login_data []byte) string {

	router := setup_router()

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

	t.Run("Admin login", func(t *testing.T) {
		login := map[string]string{"email_id": "admin@univ.edu", "password": "admin"}

		login_data, _ := json.Marshal(login)

		actual_result := run_and_get_user_type(login_data)

		if actual_result != "ADMIN" {
			t.Errorf("\nExpected:\t%s \nActual:\t%s", "ADMIN", actual_result)
		}
	})

	t.Run("professor login", func(t *testing.T) {
		login := map[string]string{"email_id": "abc@univ.edu", "password": "1234"}

		login_data, _ := json.Marshal(login)

		actual_result := run_and_get_user_type(login_data)

		if actual_result != "PROFESSOR" {
			t.Errorf("\nExpected:\t%s \nActual:\t%s", "PROFESSOR", actual_result)
		}
	})

	t.Run("student login", func(t *testing.T) {
		login := map[string]string{"email_id": "abc.pqr@univ.edu", "password": "1234"}

		login_data, _ := json.Marshal(login)

		actual_result := run_and_get_user_type(login_data)

		if actual_result != "STUDENT" {
			t.Errorf("\nExpected:\t%s \nActual:\t%s", "STUDENT", actual_result)
		}
	})
}
