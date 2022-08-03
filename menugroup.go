package model

// MenuGroup 菜单分组
type MenuGroup struct {
	Id         uint   `json:"id"`
	UserId     string `json:"user_id" gorm:"user_id"`
	MerchantId string `json:"merchant_id" gorm:"merchant_id"`
	Name       string `json:"name" gorm:"name"`
	Number     int    `json:"number" gorm:"number"`
	Desc       string `json:"desc" gorm:"desc"`
}
