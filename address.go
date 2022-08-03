package model

import "time"

// AddressModel 用户地址管理
type AddressModel struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// UserId 用户ID (登录dash平台的人)
	AdminId string `json:"admin_id"`
	// MerchantId 商户ID (例如: 黄李记作为一个商户存在)
	MerchantId string `json:"merchant_id"`
	// Status 状态
	Status string `gorm:"default:effected"`
	// CreatedAt 创建时间
	CreatedAt time.Time
	// UpdatedAt 修改时间
	UpdatedAt time.Time

	// gorm.Model
	UserId int    `json:"user_id" gorm:"user_id"`
	Name   string `json:"name" gorm:"name" `
	Phone  string `json:"phone" gorm:"phone"`
	Gender string `json:"gender" gorm:"gender"`
	Addr   string `json:"addr" gorm:"addr"`
	Tag    string `json:"tag" gorm:"tag"`
}

