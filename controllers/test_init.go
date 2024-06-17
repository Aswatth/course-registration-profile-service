package controllers

import (
	"course-registration-system/profile-service/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var sql_database services.MySqlDatabase

func setup_test_router() *gin.Engine {
	r := gin.Default()
	return r
}

func load_env() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sql_database = *new(services.MySqlDatabase)
	sql_database.Connect(os.Getenv("MYSQL_CONNECTION_STRING"))
}
