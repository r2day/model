package model

import "time"

type CartModel struct {
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
	
	UserId        int    `json:"user_id" gorm:"user_id"`
	ProductId     int    `json:"product_id" gorm:"product_id" `
	ProductName   string `json:"product_name" gorm:"product_name"`
	ProductNumber int    `json:"product_number" gorm:"product_number"`
	TotalPrice    int    `json:"total_price" gorm:"total_price"`
	UnitPrice     int    `json:"unit_price" gorm:"unit_price"`
	Pic           string `json:"pic" gorm:"pic"`
	// 特性
	Characteristic string `json:"characteristic" gorm:"characteristic"`
}
