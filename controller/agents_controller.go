package controller

import (
	"github.com/gin-gonic/gin"
	"job-master/data"
	"job-master/model"
	"job-master/service"
	"net/http"
)

// 등록
func Register(c *gin.Context)  {
	//agent server name
	var agentRegisterModel model.AgentRegisterModel
	if err := c.ShouldBindJSON(&agentRegisterModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	agent := data.Agent{Name: agentRegisterModel.Name, Host: agentRegisterModel.Host}
	service.AddAgent(&agent)
	//return enroll generate key (timeout)
}

// 조회
func Find(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "data",
	})
}

