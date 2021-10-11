package controller

import "github.com/gin-gonic/gin"

// 등록
func register()  {
	//agent server name
	//return enroll generate key (timeout)
}

// 조회
func Find(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "data",
	})
}

