package routes

import "github.com/gin-gonic/gin"

func GetRoutes(router *gin.Engine)  {
	v1 := router.Group("/api")
	addAgentsRoutes(v1)
}
