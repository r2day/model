package model

import "time"

// CustomerReviews 客户评价
type CustomerReviews struct {
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

	StoreId      string `json:"store_id"`
	CustomerId   string `json:"customer_id"`
	Rating       int    `json:"rating"`
	Comment      string `json:"comment"`
	ReviewStatus string `json:"review_status"`
}

// ListAll 获取所有数据
// 店铺信息
func (m CustomerReviews) ListAll() ([]CustomerReviews, error) {
	instance := make([]CustomerReviews, 0)
	err := DataHandler.Table("customer_reviews").
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
func (m CustomerReviews) GetOne() (CustomerReviews, error) {
	instance := CustomerReviews{}
	err := DataHandler.Table("customer_reviews").
		Where("status = ? and merchant_id = ? and id = ?", m.Status, m.MerchantId, m.Id).
		First(&instance).Error
	if err != nil {
		return instance, err
	} else {
		// 保存成功可以进行消息通知操作
		return instance, nil
	}
}
