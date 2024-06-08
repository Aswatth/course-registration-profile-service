package controllers

import (
	"course-registration-system/profile-service/models"
	"course-registration-system/profile-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminProfileController struct {
	service services.AdminProfileService
}

func (obj *AdminProfileController) Init(service *services.AdminProfileService) {
	obj.service = *service
}

func (obj *AdminProfileController) Login(context *gin.Context) {
	var admin_profile models.AdminProfile

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&admin_profile); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	result := obj.service.Validate(admin_profile)

	if result {
		context.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		context.JSON(http.StatusOK, gin.H{"message": "Invalid credentials"})
	}
}

func (obj *AdminProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	admin_profile_routes := rg.Group("/admin")

	admin_profile_routes.GET("/login", obj.Login)
}
