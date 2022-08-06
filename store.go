package model

import "time"

// StoreModel 店铺信息
type StoreModel struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// UserId 用户ID (登录dash平台的人)
	UserId string `json:"user_id"`
	// MerchantId 商户ID (例如: 黄李记作为一个商户存在)
	MerchantId string `json:"merchant_id"`
	// StoreId store id
	StoreId string `json:"store_id"`
	// Status 状态
	Status string `gorm:"default:effected"`
	// CreatedAt 创建时间
	CreatedAt time.Time
	// UpdatedAt 修改时间
	UpdatedAt time.Time

	Name     string `json:"name"`
	Position string `json:"position"`
	Addr     string `json:"addr"`
	Phone    string `json:"phone"`
	Pic      string `json:"pic"`
}

// Save 保存实例
func (m StoreModel) Save() {
	DataHandler.Create(&m)
}

func (m StoreModel) GetStore(merchantId string) []StoreModel {
	instance := make([]StoreModel, 0)
	DataHandler.Where("merchant_id = ?", merchantId).Find(&instance)
	return instance
}
