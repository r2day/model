package model

import "time"

// Unit 规格
type Unit struct {
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

	// 自动增加(被引用次数)
	Ref int `json:"ref" gorm:"default:0"`
}

// Save 保存实例
func (m Unit) Save() {
	DataHandler.Create(&m)
}

// All 保存实例
func (m Unit) All(userId string) []Unit {
	instance := make([]Unit, 0)
	DataHandler.Where("user_id = ?", userId).Find(&instance)
	return instance
}

func (m Unit) Delete(id int) {
	DataHandler.Where("id = ?", id).Delete(&m)
}

// GetByItemId 获取对象
func (m Unit) GetByItemId(itemId string) *Unit {
	DataHandler.Where("item_id = ?", itemId).First(&m)
	return &m
}
