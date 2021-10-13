package data

type Agent struct {
	Id		uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Name	string `json:"name"`
	Host	string `json:"host"`
	BaseData
}
