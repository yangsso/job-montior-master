package main

import (
	"github.com/gin-gonic/gin"
	"job-master/config"
	"job-master/data"
	"job-master/routes"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
var router = gin.Default()

func main() {
	config.DB, _ = gorm.Open(sqlite.Open("./test.db"), &gorm.Config{});

	if !config.DB.Migrator().HasTable(&data.Agent{}) {
		config.DB.AutoMigrate(data.Agent{})
	}

	if !config.DB.Migrator().HasTable(&data.AgentJob{}) {
		config.DB.AutoMigrate(data.AgentJob{})
	}

	routes.GetRoutes(router)
	router.Use(gin.Logger())
	router.Run(":8080")
}