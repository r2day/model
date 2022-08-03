package model

import (
	"github.com/r2day/base/util"
	"time"
)

// CustomerInfo 客户账号信息
type CustomerInfo struct {
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

	UserId    string `json:"user_id" gorm:"user_id"`
	WxOpenId  string `json:"wx_open_id" gorm:"wx_open_id" `
	WxUnionId string `json:"wx_union_id" gorm:"wx_union_id"`
	Username  string `json:"username" gorm:"username"`
	Avatar    string `json:"avatar" gorm:"avatar"`
	Gender    string `json:"gender" gorm:"gender"`
	UserType  int    `json:"user_type" gorm:"user_type"`
}

// Save 保存实例
func (m CustomerInfo) Save() {
	DataHandler.Create(&m)
}

// Update 保存实例
func (m CustomerInfo) Update(userId string) {
	// https://stackoverflow.com/questions/39333102/how-to-create-or-update-a-record-with-gorm
	if DataHandler.Model(m).Where("user_id = ?", userId).
		Updates(m).RowsAffected == 0 {
		uniqueId := util.GetAccountId()
		m.UserId = uniqueId
		DataHandler.Debug().Create(&m)
	}
}
