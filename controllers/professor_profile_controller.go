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
	email_id := context.Param("email_id")

	//Fetch from DB
	fetched_professor_profile, err := obj.service.FetchProfessorProfile(email_id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	} else {
		context.JSON(http.StatusOK, fetched_professor_profile)
	}

}

func (obj *ProfessorProfileController) UpdatePassword(context *gin.Context) {

	email_id := context.Param("email_id")

	new_password := make(map[string]string)

	if err := context.ShouldBindJSON(&new_password); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		err := obj.service.UpdatePassword(email_id, new_password["new_password"])

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		} else {
			context.Status(http.StatusOK)
		}
	}
}

func (obj *ProfessorProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	professor_profile_routes := rg.Group("/professors")

	professor_profile_routes.GET("/:email_id", obj.FetchProfessorProfile)
	professor_profile_routes.PUT("/password/:email_id", obj.UpdatePassword)
}
