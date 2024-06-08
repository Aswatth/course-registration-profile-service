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

func (obj *StudentProfileController) FetchStudentProfile(context *gin.Context) {
	email_id := context.Query("email_id")

	//Fetch from DB
	fetched_student_profile := obj.service.FetchStudentProfile(email_id)

	context.JSON(http.StatusOK, fetched_student_profile)
}

func (obj *StudentProfileController) UpdateStudentProfile(context *gin.Context) {
	email_id := context.Query("email_id")

	//Fetch from DB
	fetched_student_profile := obj.service.FetchStudentProfile(email_id)

	var student_profile models.StudentProfile
	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&student_profile); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	student_profile.Email_id = fetched_student_profile.Email_id
	student_profile.Password = fetched_student_profile.Password

	obj.service.UpdateStudentProfile(student_profile)

	context.JSON(http.StatusOK, gin.H{"message": "Successfully updated!"})
}

func (obj *StudentProfileController) UpdateStudentPassword(context *gin.Context) {

	email_id := context.Query("email_id")

	//Fetch from DB
	fetched_student_profile := obj.service.FetchStudentProfile(email_id)

	type NewPasswordData struct {
		New_password string
	}

	var new_password_data NewPasswordData

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&new_password_data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	fetched_student_profile.Password = new_password_data.New_password

	obj.service.UpdateStudentProfile(fetched_student_profile)

	context.JSON(http.StatusOK, gin.H{"message": "Successfully updated password!"})
}

func (obj *StudentProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	student_profile_routes := rg.Group("/student")

	student_profile_routes.POST("/create", obj.CreateProfile)
	student_profile_routes.GET("/fetch", obj.FetchStudentProfile)
	student_profile_routes.PUT("/update", obj.UpdateStudentProfile)
	student_profile_routes.PUT("/update_password", obj.UpdateStudentPassword)
}
