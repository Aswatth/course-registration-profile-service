package controllers

import (
	"course-registration-system/profile-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentProfileController struct {
	service services.StudentProfileService
}

func (obj *StudentProfileController) Init(service *services.StudentProfileService) {
	obj.service = *service
}

func (obj *StudentProfileController) FetchStudentProfile(context *gin.Context) {
	email_id := context.Query("email_id")

	//Fetch from DB
	fetched_student_profile := obj.service.FetchStudentProfile(email_id)

	context.JSON(http.StatusOK, fetched_student_profile)
}

func (obj *StudentProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	student_profile_routes := rg.Group("/student")

	student_profile_routes.GET("", obj.FetchStudentProfile)
}
