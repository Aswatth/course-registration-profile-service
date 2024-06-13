package controllers

import (
	"course-registration-system/profile-service/models"
	"course-registration-system/profile-service/services"
	"course-registration-system/profile-service/utils"
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
	result := obj.service.Validate(login)

	if result == "INVALID_CREDENTIALS" {
		context.AbortWithError(http.StatusBadRequest, errors.New(result))
	} else {
		token, err := utils.GenerateToken(result)

		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, errors.New("unable to generate jwt token"))
		}

		context.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func (obj *LoginController) RegisterRoutes(rg *gin.RouterGroup) {
	login_routes := rg.Group("")

	login_routes.POST("/login", obj.Login)
}
