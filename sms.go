package model

// SmsConf 短信基本配置
type SmsConf struct {
	// 手机号
	Phone string `json:"phone" gorm:"phone"`
	// 第一个参数
	First string `json:"first" gorm:"first"`
	// 第二个参数
	Second string `json:"second" gorm:"second"`
	// 第三个参数
	Third string `json:"third" gorm:"third"`
}
