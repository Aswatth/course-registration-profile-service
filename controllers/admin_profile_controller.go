package controllers

import (
	"course-registration-system/profile-service/middlewares"
	"course-registration-system/profile-service/models"
	"course-registration-system/profile-service/services"
	"course-registration-system/profile-service/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		context.AbortWithError(http.StatusBadRequest, err)
	}

	result := obj.service.Validate(admin_profile)

	if result {
		token, err := utils.GenerateToken("admin")

		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
		}

		context.SetCookie("Authorization", token, 3600*24, "", "", true, true)
		context.AbortWithStatus(http.StatusOK)
	} else {
		context.AbortWithError(http.StatusBadRequest, errors.New("invalid credentials"))
	}
}

func (obj *AdminProfileController) CreateStudentProfile(context *gin.Context) {
	var student_profile models.StudentProfile

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&student_profile); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	//hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(student_profile.Password), bcrypt.DefaultCost)

	student_profile.Password = string(hash)

	//Store to DB
	err := obj.service.CreateStudentProfile(student_profile)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	}

	context.AbortWithStatus(http.StatusOK)
}

func (obj *AdminProfileController) UpdateStudentProfile(context *gin.Context) {
	email_id := context.Param("email_id")

	var student_profile models.StudentProfile
	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&student_profile); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	err := obj.service.UpdateStudentProfile(email_id, student_profile)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	}

	context.AbortWithStatus(http.StatusOK)
}

func (obj *AdminProfileController) DeleteStudentProfile(context *gin.Context) {
	email_id := context.Param("email_id")

	err := obj.service.DeleteStudentProfile(email_id)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	} else {
		context.AbortWithStatus(http.StatusOK)
	}
}

func (obj *AdminProfileController) CreateProfessorProfile(context *gin.Context) {

	var professor_profile models.ProfessorProfile

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&professor_profile); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	//hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(professor_profile.Password), bcrypt.DefaultCost)

	professor_profile.Password = string(hash)

	//Store to DB
	err := obj.service.CreateProfessorProfile(professor_profile)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	} else {
		context.AbortWithStatus(http.StatusOK)
	}
}

func (obj *AdminProfileController) UpdateProfessorProfile(context *gin.Context) {
	email_id := context.Param("email_id")

	var professor_profile models.ProfessorProfile
	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&professor_profile); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	err := obj.service.UpdateProfessorProfile(email_id, professor_profile)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	} else {
		context.AbortWithStatus(http.StatusOK)
	}
}

func (obj *AdminProfileController) DeleteProfessorProfile(context *gin.Context) {
	email_id := context.Param("email_id")

	err := obj.service.DeleteProfessorProfile(email_id)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	} else {
		context.AbortWithStatus(http.StatusOK)
	}
}

func (obj *AdminProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	admin_profile_routes := rg.Group("/admin")

	admin_profile_routes.GET("/login", obj.Login)

	//Student routes
	admin_profile_routes.Use(middlewares.ValidateAuthorization).POST("/students", obj.CreateStudentProfile)
	admin_profile_routes.Use(middlewares.ValidateAuthorization).PUT("/students/:email_id", obj.UpdateStudentProfile)
	admin_profile_routes.Use(middlewares.ValidateAuthorization).DELETE("/students/:email_id", obj.DeleteStudentProfile)

	//Professor routes
	admin_profile_routes.Use(middlewares.ValidateAuthorization).POST("/professors", obj.CreateProfessorProfile)
	admin_profile_routes.Use(middlewares.ValidateAuthorization).PUT("/professors/:email_id", obj.UpdateProfessorProfile)
	admin_profile_routes.Use(middlewares.ValidateAuthorization).DELETE("/professors/:email_id", obj.DeleteProfessorProfile)
}
