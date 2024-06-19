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
	fetched_student_profile, err := obj.service.FetchStudentProfile(email_id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		context.JSON(http.StatusOK, fetched_student_profile)
	}

}

func (obj *StudentProfileController) UpdatePassword(context *gin.Context) {

	email_id := context.Param("email_id")

	new_password := make(map[string]string)

	if err := context.ShouldBindJSON(&new_password); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	err := obj.service.UpdatePassword(email_id, new_password["new_password"])

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		context.Status(http.StatusOK)
	}

}

func (obj *StudentProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	student_profile_routes := rg.Group("/students")

	student_profile_routes.GET("", obj.FetchStudentProfile)

	student_profile_routes.PUT("/password/:email_id", obj.UpdatePassword)
}
