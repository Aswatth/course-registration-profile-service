package controllers

import "github.com/gin-gonic/gin"

func setup_test_router() *gin.Engine {
	r := gin.Default()
	return r
}
