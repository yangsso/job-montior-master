package data

import (
	"time"
)
type AgentJob struct {
	Id			uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	AgentId 	uint64 `gorm:"index;not null" json:"agent_id"`
	Name		string `gorm:"index;not null" json:"name"`
	LastExecutionKey uint64 `json:"last_execution_key"`
	LastExecutionTime time.Time	`json:"last_execution_time"`
	BaseData
}
