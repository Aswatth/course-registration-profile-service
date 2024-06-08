package controllers

import (
	"course-registration-system/profile-service/models"
	"course-registration-system/profile-service/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfessorProfileController struct {
	service services.ProfessorProfileService
}

func (obj *ProfessorProfileController) Init(service *services.ProfessorProfileService) {
	obj.service = *service
}

func (obj *ProfessorProfileController) CreateProfile(context *gin.Context) {

	var professor_profile models.ProfessorProfile

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&professor_profile); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	fmt.Print(professor_profile)

	//Store to DB
	obj.service.CreateProfile(professor_profile)
}

func (obj *ProfessorProfileController) FetchProfessorProfile(context *gin.Context) {
	email_id := context.Query("email_id")

	//Fetch from DB
	fetched_professor_profile := obj.service.FetchProfessorProfile(email_id)

	context.JSON(http.StatusOK, fetched_professor_profile)
}

func (obj *ProfessorProfileController) UpdateProfessorProfile(context *gin.Context) {
	email_id := context.Query("email_id")

	//Fetch from DB
	fetched_professor_profile := obj.service.FetchProfessorProfile(email_id)

	var professor_profile models.ProfessorProfile
	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&professor_profile); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	professor_profile.Email_id = fetched_professor_profile.Email_id
	professor_profile.Password = fetched_professor_profile.Password

	obj.service.UpdateProfessorProfile(professor_profile)

	context.JSON(http.StatusOK, gin.H{"message": "Successfully updated!"})
}

func (obj *ProfessorProfileController) UpdateStudentPassword(context *gin.Context) {

	email_id := context.Query("email_id")

	//Fetch from DB
	fetched_professor_profile := obj.service.FetchProfessorProfile(email_id)

	type NewPasswordData struct {
		New_password string
	}

	var new_password_data NewPasswordData

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&new_password_data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	fetched_professor_profile.Password = new_password_data.New_password

	obj.service.UpdateProfessorProfile(fetched_professor_profile)

	context.JSON(http.StatusOK, gin.H{"message": "Successfully updated password!"})
}

func (obj *ProfessorProfileController) DeleteProfessorProfile(context *gin.Context) {
	email_id := context.Query("email_id")

	obj.service.DeleteProfessorProfile(email_id)

	context.JSON(http.StatusOK, gin.H{"message": "Successfully deleted!"})
}

func (obj *ProfessorProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	professor_profile_routes := rg.Group("/professor")

	professor_profile_routes.POST("/create", obj.CreateProfile)
	professor_profile_routes.GET("/fetch", obj.FetchProfessorProfile)
	professor_profile_routes.PUT("/update", obj.UpdateProfessorProfile)
	professor_profile_routes.PUT("/update_password", obj.UpdateStudentPassword)
	professor_profile_routes.DELETE("/delete", obj.DeleteProfessorProfile)
}
