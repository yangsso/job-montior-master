package routes

import (
	"github.com/gin-gonic/gin"

	"job-master/controller"
)

func addAgentsRoutes(rg *gin.RouterGroup)  {
	agents := rg.Group("/agents")
	agents.GET("", controller.Find)
}
