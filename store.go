package model

// StoreModel 店铺信息
type StoreModel struct {
	UserId   string `json:"user_id"`
	Id       uint   `json:"id" gorm:"unique"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Position string `json:"position"`
	Addr     string `json:"addr"`
	Phone    string `json:"phone"`
	Pic      string `json:"pic"`
}
