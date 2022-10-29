package model

import (
	"strconv"
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
	CardId string `json:"card_id" gorm:"index:idx_card_id,unique"`
	// CustomerId 客户编号
	CustomerId string `json:"customer_id"`
	// 手机号
	Phone string `json:"phone"`
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

	CardCreatedDate string `json:"card_created_date"`
	// 有效期
	Expire string `json:"expire"`
	// From 来源
	From string `json:"from"`
	// Channel 渠道
	Channel string `json:"channel"`
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

// SaveALine 保存实例
func (m MemberInfo) SaveALine(value []string) {

	m.CardId = value[0]
	m.CustomerId = value[1]
	m.Phone = value[2]
	m.Name = value[3]
	m.Gender = value[4]
	m.BirthDay = value[5]
	m.CardType = value[6]
	m.CardStatus = value[7]
	m.CardLevel = value[8]
	m.CardFrom = value[9]
	m.Balance, _ = strconv.ParseFloat(value[10], 64)
	m.CashCharge, _ = strconv.ParseFloat(value[11], 64)
	m.Freezing, _ = strconv.ParseFloat(value[12], 64)
	m.Gift, _ = strconv.ParseFloat(value[13], 64)
	m.Integral, _ = strconv.ParseUint(value[14], 10, 64)
	m.TotalBalance, _ = strconv.ParseFloat(value[15], 64)
	m.TotalBalanceCounter, _ = strconv.ParseUint(value[16], 10, 64)

	m.TotalCumulativeConsumption, _ = strconv.ParseFloat(value[17], 64)
	m.TotalCumulativeConsumptionCounter, _ = strconv.ParseUint(value[18], 10, 64)
	m.DebitTotalLimit, _ = strconv.ParseFloat(value[19], 64)
	m.DebitLeftLimit, _ = strconv.ParseFloat(value[20], 64)
	m.DebitUsedLimit, _ = strconv.ParseFloat(value[21], 64)
	m.EntityCardId = value[22]
	m.CardCreatedDate = value[23]
	m.Expire = value[24]

	DataHandler.Create(&m)
}
