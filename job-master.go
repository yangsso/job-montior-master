package main

import (
	"github.com/gin-gonic/gin"
	"job-master/routes"
)
var router = gin.Default()

func main() {
	routes.GetRoutes(router)
	router.Use(gin.Logger())
	router.Run(":8080")
}