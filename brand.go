package model

import (
	"time"
)

// Brand 品牌
type Brand struct {
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
	// logo 图标地址
	BrandLogo string `json:"brand_logo"`
	// BrandSince 诞生时间
	BrandSince string `json:"brand_since"`
	// BrandDesc 品牌描述
	BrandDesc string `json:"brand_desc"`
}

// GetByItemId 获取对象
func (m Brand) GetByItemId(itemId string) *Brand {
	DataHandler.Where("item_id = ?", itemId).First(&m)
	return &m
}
