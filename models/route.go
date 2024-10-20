package models

import "github.com/gin-gonic/gin"

type DefaultRoutes struct {
	UnauthenticatedRoute *gin.RouterGroup
	AuthenticatedRoute   *gin.RouterGroup
	InternalRoute        *gin.RouterGroup
}
