package model

import (
	"log"
	"time"
)

// MerchantApply 商户入驻申请
// 会进行增长的表后续都会进行数据迁移或者清理
type MerchantApply struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	// UpdatedAt 修改时间
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	// 自定义字段
	Email         string `json:"email" gorm:"email"`
	Name          string `json:"name" gorm:"name"`
	IdCard        string `json:"id_card" gorm:"id_card"`
	PrincipalName string `json:"principal_name" gorm:"principal_name"`
	Phone         string `json:"phone" gorm:"phone"`
	License       string `json:"license" gorm:"license"`
	Status        string `json:"status" gorm:"status"`
	// Type 商户类型(加盟、连锁、新)
	Type string `json:"type" gorm:"type"`
	// 申请回执
	ApplyCode string `json:"apply_code" gorm:"apply_code"`
	// 申请通过后生成的merchant_id
	MerchantId string `json:"merchant_id" gorm:"merchant_id"`
	// 申请通过后生成的密钥
	MerchantKey string `json:"merchant_key" gorm:"merchant_key"`
}

// MerchantInfo 商户信息表
// 当申请审核通过后保留全部申请信息
//type MerchantInfo struct {
//	gorm.Model
//	// 自定义字段
//	MerchantApply
//	// 申请通过后生成的密钥
//	MerchantKey string `json:"merchant_key" gorm:"merchant_key"`
//}

// Save 保存实例
func (m MerchantApply) Save() error {
	err := DataHandler.Create(&m).Error
	if err != nil {
		return err
	} else {
		// 保存成功可以进行消息通知操作
		// TODO send to mq
		log.Println("send to mq")
		return nil
	}
}

// ListAll 获取所有数据
// 以便管理员进行审核操作
func (m MerchantApply) ListAll() ([]Dishes, error) {
	instance := make([]MerchantApply, 0)
	err := DataHandler.Where("status = ?", m).Find(&instance).Error
	if err != nil {
		return nil, err
	} else {
		// 保存成功可以进行消息通知操作
		// TODO send to mq
		log.Println("send to mq")
		return nil, nil
	}
}

// FindOne 申请人可以查看申请进度
func (m MerchantApply) FindOne() (MerchantApply, error) {

	// 查询条件
	cond := map[string]interface{}{
		"apply_code": m.ApplyCode,
	}
	err := DataHandler.Debug().Table("orders").
		Select("*").
		Where(cond).First(&m).Error
	if err != nil {
		return m, err
	} else {
		// 保存成功可以进行消息通知操作
		// TODO send to mq
		log.Println("send to mq")
		return m, nil
	}
}

// UpdateStatus 更新状态
// 审批通过/失败后进行
func (m MerchantApply) UpdateStatus(status string) error {

	switch status {
	case "ok":
		log.Println("apply pass")
		m.MerchantKey = "generate key ...."
	case "reject":
		log.Println("apply reject")
	}
	err := DataHandler.Updates(m).
		Where("id = ?", m.Id).Error
	if err != nil {
		return err
	} else {
		// 保存成功可以进行消息通知操作
		// TODO send to mq
		log.Println("send to mq")
		return nil
	}
}
