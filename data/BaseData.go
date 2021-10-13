package data

import "time"

type BaseData struct {
	CreatedAt	time.Time `json:"created_at" gorm:"autoCreateTime"`
	CreatedBy	string	`json:"created_by"`
	ModifiedAt 	time.Time `json:"modified_at" gorm:"autoUpdateTime"`
	ModifiedBy	string	`json:"modified_by"`
}
