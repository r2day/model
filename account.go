package model

import (
	"time"
)

// AccountInfo 账号信息
type AccountInfo struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// UserId 用户ID (登录dash平台的人)
	UserId string `json:"user_id"`
	// MerchantId 商户ID (例如: 黄李记作为一个商户存在)
	MerchantId string `json:"merchant_id"`
	// Status 状态
	Status string `gorm:"default:effected"`
	// CreatedAt 创建时间
	CreatedAt time.Time
	// UpdatedAt 修改时间
	UpdatedAt time.Time

	// 客户id
	CustomerId string `json:"customer_id"`
	// Avatar 头像
	Avatar string `json:"avatar"`
	// NickName 昵称
	NickName string `json:"nick_name"` // db func
	// 性别
	Gender string `json:"gender"`
	// 账号类型
	TypeId string `json:"type_id"`
	// 手机号
	Phone string `json:"phone"`
	// 出生日期
	BirthDay string `json:"birth_day"`
}

// Save 保存实例
func (m AccountInfo) Save() {
	DataHandler.Create(&m)
}

// UpdateStatus 更新状态
func (m AccountInfo) UpdateStatus(id int, status string) {
	DataHandler.UpdateColumn("status", status).
		Where("id = ?", id)
}

// Delete 删除
func (m AccountInfo) Delete(id int) AccountInfo {
	// TODO 以后放到一个事务中执行
	// 删除前先查出记录
	DataHandler.
		Where("id = ?", id).First(&m)
	// 执行删除
	DataHandler.Where("id = ?", id).Delete(&m)
	return m
}

// All 保存实例
func (m AccountInfo) All(merchantId string) []AccountInfo {
	instance := make([]AccountInfo, 0)
	DataHandler.Where("merchant_id = ?", merchantId).Find(&instance)
	return instance
}
