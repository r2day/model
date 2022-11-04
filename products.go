package model

import "time"

// Products 产品信息
type Products struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// MerchantId 商户ID (例如: 黄李记作为一个商户存在)
	MerchantId string `json:"merchant_id" gorm:"index:idx_merchant"`
	// Status 状态
	Status string `gorm:"default:effected" gorm:"index:idx_status"`
	// CreatedAt 创建时间
	CreatedAt time.Time
	// UpdatedAt 修改时间
	UpdatedAt time.Time

	CategoryId  int     `json:"category_id"`
	Reference   string  `json:"reference"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	Price       float64 `json:"price"`
	Thumbnail   string  `json:"thumbnail"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Stock       string  `json:"stock"`
}

// ListAll 获取所有数据
// 店铺信息
func (m Products) ListAll() ([]Products, error) {
	instance := make([]Products, 0)
	err := DataHandler.Table("products").
		Where("status = ? and merchant_id = ?", m.Status, m.MerchantId).
		Find(&instance).Error
	if err != nil {
		return nil, err
	} else {
		// 保存成功可以进行消息通知操作
		return instance, nil
	}
}

// GetOne 获取单个详情
func (m Products) GetOne() (Products, error) {
	instance := Products{}
	err := DataHandler.Table("products").
		Where("status = ? and merchant_id = ? and id = ?", m.Status, m.MerchantId, m.Id).
		First(&instance).Error
	if err != nil {
		return instance, err
	} else {
		// 保存成功可以进行消息通知操作
		return instance, nil
	}
}
