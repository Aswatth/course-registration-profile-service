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

func (obj *AdminProfileController) CreateStudentProfile(context *gin.Context) {
	type NewStudentData struct {
		Email_id         string
		Password         string
		First_name       string
		Last_name        string
		Program_enrolled string
	}

	var new_student NewStudentData
	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&new_student); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		//separate profile and login data
		login_data := models.Login{Email_id: new_student.Email_id, Password: new_student.Password, User_type: "STUDENT"}
		student_profile_data := models.StudentProfile{Email_id: new_student.Email_id, First_name: new_student.First_name, Last_name: new_student.Last_name, Program_enrolled: new_student.Program_enrolled}

		//Store to DB
		err := obj.service.CreateStudentProfile(login_data, student_profile_data)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		} else {
			context.Status(http.StatusOK)
		}
	}
}

func (obj *AdminProfileController) GetStudentProfile(context *gin.Context) {
	email_id := context.Query("email_id")

	student_profile, err := obj.service.GetStudentProfile(email_id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
	} else {
		context.JSON(http.StatusOK, student_profile)
	}
}

func (obj *AdminProfileController) GetAllStudentProfiles(context *gin.Context) {

	if(context.Query("email_id") != "") {
		obj.GetStudentProfile(context)
		return
	}

	student_profile_list, err := obj.service.GetAllStudentProfiles()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
	} else {
		context.JSON(http.StatusOK, student_profile_list)
	}
}

func (obj *AdminProfileController) UpdateStudentProfile(context *gin.Context) {
	email_id := context.Param("email_id")

	var student_profile models.StudentProfile
	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&student_profile); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		err := obj.service.UpdateStudentProfile(email_id, student_profile)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
		} else {
			context.AbortWithStatus(http.StatusOK)
		}
	}
}

func (obj *AdminProfileController) DeleteStudentProfile(context *gin.Context) {
	email_id := context.Param("email_id")

	err := obj.service.DeleteProfile(email_id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
	} else {
		context.Status(http.StatusOK)
	}
}

func (obj *AdminProfileController) CreateProfessorProfile(context *gin.Context) {

	type NewProfessorData struct {
		Email_id    string
		Password    string
		First_name  string
		Last_name   string
		Designation string
		Department  string
	}

	var new_professor_data NewProfessorData

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&new_professor_data); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		//separate profile and login data
		login_data := models.Login{Email_id: new_professor_data.Email_id, Password: new_professor_data.Password, User_type: "PROFESSOR"}
		professor_profile_data := models.ProfessorProfile{Email_id: new_professor_data.Email_id, First_name: new_professor_data.First_name, Last_name: new_professor_data.Last_name, Department: new_professor_data.Department, Designation: new_professor_data.Designation}

		//Store to DB
		err := obj.service.CreateProfessorProfile(login_data, professor_profile_data)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
		} else {
			context.Status(http.StatusOK)
		}
	}

}

func (obj *AdminProfileController) GetProfessorProfile(context *gin.Context) {
	email_id := context.Query("email_id")

	professor_profile, err := obj.service.GetProfessorProfile(email_id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
	} else {
		context.JSON(http.StatusOK, professor_profile)
	}
}

func (obj *AdminProfileController) GetAllProfessorProfiles(context *gin.Context) {

	if(context.Query("email_id") != "") {
		obj.GetProfessorProfile(context)
		return
	}

	professor_profile_list, err := obj.service.GetAllProfessorProfiles()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
	} else {
		context.JSON(http.StatusOK, professor_profile_list)
	}
}

func (obj *AdminProfileController) UpdateProfessorProfile(context *gin.Context) {
	email_id := context.Param("email_id")

	var professor_profile models.ProfessorProfile
	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&professor_profile); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		err := obj.service.UpdateProfessorProfile(email_id, professor_profile)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
		} else {
			context.AbortWithStatus(http.StatusOK)
		}
	}

}

func (obj *AdminProfileController) DeleteProfessorProfile(context *gin.Context) {
	email_id := context.Param("email_id")

	err := obj.service.DeleteProfile(email_id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
	} else {
		context.Status(http.StatusOK)
	}
}

func (obj *AdminProfileController) UpdatePassword(context *gin.Context) {

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

func (obj *AdminProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	admin_profile_routes := rg.Group("/admin")

	admin_profile_routes.PUT("/password/:email_id", obj.UpdatePassword)
	//Student routes
	admin_profile_routes.POST("/students", obj.CreateStudentProfile)
	admin_profile_routes.GET("/students", obj.GetAllStudentProfiles)
	admin_profile_routes.PUT("/students/:email_id", obj.UpdateStudentProfile)
	admin_profile_routes.DELETE("/students/:email_id", obj.DeleteStudentProfile)

	//Professor routes
	admin_profile_routes.POST("/professors", obj.CreateProfessorProfile)
	admin_profile_routes.GET("/professors", obj.GetAllProfessorProfiles)
	admin_profile_routes.PUT("/professors/:email_id", obj.UpdateProfessorProfile)
	admin_profile_routes.DELETE("/professors/:email_id", obj.DeleteProfessorProfile)
}
