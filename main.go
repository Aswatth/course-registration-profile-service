package main

import (
	"course-registration-system/profile-service/controllers"
	"course-registration-system/profile-service/services"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	PORT                 int
	DB_TYPE              string
	DB_NAME              string
	DB_CONNECTION_STRING string
}

func (config *Config) LoadConfig(file_name string) {

	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal("Error occured while reading config file.", err)
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error occured while reading config file.", err)
	}
}

func main() {
	//Load config
	config := new(Config)
	config.LoadConfig("config.json")

	sql_database := new(services.MySqlDatabase)
	sql_database.Connect(config.DB_CONNECTION_STRING)

	student_profile_service := new(services.StudentProfileService)
	student_profile_service.Init(*sql_database)

	student_profile_controller := new(controllers.StudentProfileController)
	student_profile_controller.Init(student_profile_service)

	professor_profile_service := new(services.ProfessorProfileService)
	professor_profile_service.Init(*sql_database)

	professor_profile_controller := new(controllers.ProfessorProfileController)
	professor_profile_controller.Init(professor_profile_service)

	admin_profile_service := new(services.AdminProfileService)
	admin_profile_service.Init(*sql_database)

	admin_profile_controller := new(controllers.AdminProfileController)
	admin_profile_controller.Init(admin_profile_service)

	server := gin.Default()

	base_path := server.Group("")
	student_profile_controller.RegisterRoutes(base_path)
	professor_profile_controller.RegisterRoutes(base_path)
	admin_profile_controller.RegisterRoutes(base_path)

	server.Run(":" + fmt.Sprint(config.PORT))
}
