package model

// CustomerGroups 客户分组
type CustomerGroups struct {
	BaseModel

	Name string `json:"name" gorm:"index:name"`
}
