package model

import (
	"time"

	"github.com/r2day/enum"
)

// Finance 账号信息
type Finance struct {
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

	// AccountId 账号id
	AccountId string `json:"account_id" gorm:"account_id"`
	// FKind 金融类型, 例如: 积分，余额，优惠券
	Kind enum.Fkind `json:"kind" gorm:"kind"`
	// Fid 金融编号
	FId string `json:"item_id" gorm:"item_id"`
	// Balance 余额
	Balance float64 `json:"balance" gorm:"balance"`
	// Currency 币种
	Currency string `json:"currency" gorm:"currency"`
}
