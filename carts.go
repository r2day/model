package model

type Cards struct {

	// 卡类型
	CardType string `json:"card_type"`
	// 卡状态
	CardStatus string `json:"card_status"`
	// 卡等级
	CardLevel string `json:"card_level"`
	// 开卡店铺
	CardFrom string `json:"card_from"`
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
