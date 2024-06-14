package controllers

import (
	"course-registration-system/profile-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfessorProfileController struct {
	service services.ProfessorProfileService
}

func (obj *ProfessorProfileController) Init(service *services.ProfessorProfileService) {
	obj.service = *service
}

func (obj *ProfessorProfileController) FetchProfessorProfile(context *gin.Context) {
	email_id := context.Query("email_id")

	//Fetch from DB
	fetched_professor_profile := obj.service.FetchProfessorProfile(email_id)

	context.JSON(http.StatusOK, fetched_professor_profile)
}

func (obj *ProfessorProfileController) UpdatePassword(context *gin.Context) {

	email_id := context.Param("email_id")

	type new_password struct {
		New_password string
	}

	var new_password_data new_password

	if err := context.ShouldBindJSON(&new_password_data); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	err := obj.service.UpdatePassword(email_id, new_password_data.New_password)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	context.Status(http.StatusOK)
}

func (obj *ProfessorProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	professor_profile_routes := rg.Group("/professors")

	professor_profile_routes.GET("", obj.FetchProfessorProfile)
	professor_profile_routes.PUT("/password/:email_id", obj.UpdatePassword)
}
