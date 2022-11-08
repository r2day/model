package model

import "time"

// BaseModel 基本数据表模型
type BaseModel struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`

	// MerchantId 商户ID (例如: 黄李记作为一个商户存在)
	MerchantId string `json:"merchant_id" gorm:"index:idx_merchant"`

	// Status 状态
	Status string `gorm:"default:effected" gorm:"index:idx_status"`

	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt 修改时间
	UpdatedAt time.Time `json:"updated_at"`
}

// Save 保存实例
func (m BaseModel) save(instance interface{}) {
	DataHandler.Create(&instance)
}

// all 获取所有数据
func (m BaseModel) all(table string, instance interface{}) error {
	//instance := make([]BaseModel, 0)
	err := DataHandler.Table(table).
		Where("status = ? and merchant_id = ?", m.Status, m.MerchantId).
		Find(instance).Error
	if err != nil {
		return err
	}
	return nil
}
