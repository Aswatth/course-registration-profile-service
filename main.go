package main

import (
	"course-registration-system/profile-service/controllers"
	"course-registration-system/profile-service/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sql_database := new(services.MySqlDatabase)
	sql_database.Connect(os.Getenv("MYSQL_CONNECTION_STRING"))

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

	server.Run(":" + os.Getenv("PORT"))
}
