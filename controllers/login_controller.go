package controllers

import (
	"course-registration-system/profile-service/models"
	"course-registration-system/profile-service/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	service services.LoginService
}

func (obj *LoginController) Init(service *services.LoginService) {
	obj.service = *service
}

func (obj *LoginController) Login(context *gin.Context) {

	var login models.Login

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&login); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	//Check if given credentials are correct. If yes, then corresponding user_type is returned else INVALID_CREDENTIALS is returned
	user_type := obj.service.Validate(login)

	if user_type == "INVALID_CREDENTIALS" {
		context.AbortWithError(http.StatusBadRequest, errors.New(user_type))
	} else {
		context.JSON(http.StatusOK, gin.H{"user_type": user_type})
	}
}

func (obj *LoginController) RegisterRoutes(rg *gin.RouterGroup) {
	login_routes := rg.Group("")

	login_routes.POST("/login", obj.Login)
}
