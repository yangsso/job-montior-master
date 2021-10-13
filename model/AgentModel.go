package model

type AgentRegisterModel struct {
	Name string `json:"name" binding:"required"`
	Host string `json:"host" binding:"required"`
}
