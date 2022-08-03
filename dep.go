package model

import "time"

// Department 部门
type Department struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// UserId 用户ID (登录dash平台的人)
	UserId string `json:"user_id"`
	// MerchantId 商户ID (例如: 黄李记作为一个商户存在)
	MerchantId string `json:"merchant_id"`
	// Status 状态
	Status string `json:"status" gorm:"default:effected"`
	// CreatedAt 创建时间
	CreatedAt time.Time
	// UpdatedAt 修改时间
	UpdatedAt time.Time

	// Name 业务名称
	ItemName string `json:"item_name"`
	// ItemId 项目编号(也是唯一值)
	ItemId string `json:"item_id"` // db func

	TypeId    string `json:"type_id"`
	TypeName  string `json:"type_name"`
	BrandName string `json:"brand_name"`
	BrandId   string `json:"brand_id"`
	Desc      string `json:"desc"`
	From      string `json:"from"`
}

// Save 保存实例
func (m Department) Save() {
	DataHandler.Create(&m)
}

// GetByItemId 获取对象
func (m Department) GetByItemId(itemId string) *Department {
	DataHandler.Where("item_id = ?", itemId).First(&m)
	return &m
}
