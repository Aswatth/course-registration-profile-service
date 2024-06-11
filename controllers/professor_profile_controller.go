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

func (obj *ProfessorProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	professor_profile_routes := rg.Group("/professor")

	professor_profile_routes.GET("", obj.FetchProfessorProfile)
}
