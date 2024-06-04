package controllers

import (
	"course-registration-system/profile-service/models"
	"course-registration-system/profile-service/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentProfileController struct {
	service services.StudentProfileService
}

func (obj *StudentProfileController) Init(service *services.StudentProfileService) {
	obj.service = *service
}

func (obj *StudentProfileController) CreateProfile(context *gin.Context) {

	var student_profile models.StudentProfile

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&student_profile); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	fmt.Print(student_profile)

	//Store to DB
	obj.service.CreateProfile(student_profile)
}

func (obj *StudentProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	student_profile_routes := rg.Group("/student")

	student_profile_routes.POST("/create", obj.CreateProfile)
}
