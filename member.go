package model

import (
	"time"
)

// MemberInfo 会员信息
type MemberInfo struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// MerchantId 商户ID (例如: 黄李记作为一个商户存在)
	MerchantId string `json:"merchant_id"`
	// Status 状态
	Status string `gorm:"default:effected"`
	// CreatedAt 创建时间
	CreatedAt time.Time
	// UpdatedAt 修改时间
	UpdatedAt time.Time

	// CardId 卡号
	CardId string `json:"card_id"`
	// CustomerId 客户编号
	CustomerId string `json:"customer_id"`
	// 手机号
	Phone string `json:"phone" gorm:"index:idx_phone,unique"`
	// 姓名
	Name string `json:"name"`
	// 性别
	Gender string `json:"gender"`
	// 出生日期
	BirthDay string `json:"birth_day"`
	// 卡类型
	CardType string `json:"card_type"`
	// 卡状态
	CardStatus string `json:"card_status"`
	// 卡等级
	CardLevel string `json:"card_level"`
	// 开卡店铺
	CardFrom string `json:"card_from"`

	// ====  资金信息 ===
	// 卡余额
	Balance float64 `json:"balance"`
	// 现金卡值
	CashCharge float64 `json:"cash_charge"`
	// 冻结卡值
	Freezing float64 `json:"freezing"`
	// 赠送卡值
	Gift float64 `json:"gift"`
	// 积分余额
	Integral uint64 `json:"integral"`
	// 累计储值总额
	TotalBalance float64 `json:"total_balance"`
	//累计储值次数
	TotalBalanceCounter uint64 `json:"total_balance_counter"`
	// 累计消费总额
	TotalCumulativeConsumption float64 `json:"total_cumulative_consumption"`
	// 累计消费总额
	TotalCumulativeConsumptionCounter uint64 `json:"total_cumulative_consumption_counter"`
	// 挂帐总额度
	DebitTotalLimit float64 `json:"debit_total_limit"`
	// 挂帐剩余额度
	DebitLeftLimit float64 `json:"debit_left_limit"`

	// 已用额度
	DebitUsedLimit float64 `json:"debit_used_limit"`
	// 实体卡号
	EntityCardId string `json:"entity_card_id"`
	// 开卡日期

	CardCreatedDate time.Time `json:"card_created_date"`
	// 有效期
	Expire string `json:"expire"`
}

// Save 保存实例
func (m MemberInfo) Save() {
	DataHandler.Create(&m)
}

// SaveCsvLine 保存实例
func (m MemberInfo) SaveCsvLine(key, value string) {
	// INSERT INTO `user_info` (`user_id`,`door_id`,`email`,`address`,`create_time`,`update_time`)
	// VALUES
	//(666,888,'test123@qq.com','北京市海淀区','2021-07-28 22:26:20.241','2021-07-28 22:26:20.241')
	// ON DUPLICATE KEY UPDATE `email`=VALUES(`email`),`address`=VALUES(`address`),`update_time`=VALUES(`update_time`)
	sql := "INSERT INTO member_infos (?) VALUES (?) ON DUPLICATE KEY UPDATE `phone`=VALUES(`phone`)"
	DataHandler.Exec(sql, key, value)
}
