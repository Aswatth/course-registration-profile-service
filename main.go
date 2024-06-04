package main

import (
	// "course-registration-system/profile-service/models"
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

	server := gin.Default()

	base_path := server.Group("")
	student_profile_controller.RegisterRoutes(base_path)

	server.Run(":" + fmt.Sprint(config.PORT))
}
