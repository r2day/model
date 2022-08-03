package model

// FormulaGroup 菜单分组
type FormulaGroup struct {
	Id          uint   `json:"id"`
	UserId      string `json:"user_id" gorm:"user_id"`
	MerchantId  string `json:"merchant_id" gorm:"merchant_id"`
	Name        string `json:"name" gorm:"name"`
	Alias       string `json:"alias" gorm:"alias"`
	PriceType   string `json:"price_type" gorm:"price_type"`
	PriceAmount int    `json:"price_amount" gorm:"price_amount"`
	PriceDesc   string `json:"price_desc" gorm:"price_desc"`
}
