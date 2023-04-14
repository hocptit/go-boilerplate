package _type

import "github.com/gin-gonic/gin"

type V1Route struct {
	V1PublicRoute    *gin.RouterGroup
	V1ProtectedRoute *gin.RouterGroup
}
