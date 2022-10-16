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
	CreatedAt time.Time
	// UpdatedAt 修改时间
	UpdatedAt time.Time

	// 自定义字段
	Email         string `json:"email"`
	Name          string `json:"name"`
	IdCard        string `json:"id_card"`
	PrincipalName string `json:"principal_name"`
	// 每个手机号仅能申请一次
	Phone   string `json:"phone" gorm:"index:idx_phone,unique"`
	License string `json:"license"`
	Status  string `json:"status"`
	// Type 商户类型(加盟、连锁、新)
	Type string `json:"type"`
	// 申请回执
	ApplyCode string `json:"apply_code" gorm:"index:idx_apply_code,unique"`
	// 申请通过后生成的merchant_id
	MerchantId string `json:"merchant_id"`
	// 申请通过后生成的密钥
	MerchantKey string `json:"merchant_key"`
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

// FindIfPhoneHasRegister 检查当前手机号是不是已经注册过
func (m MerchantApply) FindIfPhoneHasRegister() (MerchantApply, error) {

	// 查询条件
	cond := map[string]interface{}{
		"phone": m.Phone,
	}
	err := DataHandler.Debug().Table("merchant_applies").
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
func (m MerchantApply) ListAll() ([]MerchantApply, error) {
	instance := make([]MerchantApply, 0)
	err := DataHandler.Table("merchant_applies").Where("status = ?", m.Status).Find(&instance).Error
	if err != nil {
		return nil, err
	} else {
		// 保存成功可以进行消息通知操作
		// TODO send to mq
		log.Println("send to mq")
		return instance, nil
	}
}

// FindOne 申请人可以查看申请进度
func (m MerchantApply) FindOne() (MerchantApply, error) {

	// 查询条件
	cond := map[string]interface{}{
		"apply_code": m.ApplyCode,
	}
	err := DataHandler.Debug().Table("merchant_applies").
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

// FindByMerchantId 通过商户id查询
// 审批通过后才有商户id
func (m MerchantApply) FindByMerchantId() (MerchantApply, error) {

	// 查询条件
	cond := map[string]interface{}{
		"merchant_id": m.MerchantId,
		"status":      "ok", // TODO 状态统一定义到enum中
	}
	err := DataHandler.Debug().Table("merchant_applies").
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
func (m MerchantApply) UpdateStatus(status string) (err error) {
	switch status {
	case "ok":
		log.Println("apply pass")
		merchantId := "xxx"
		merchantKey := "yyy"
		cond := map[string]interface{}{
			"id": m.Id,
		}
		err = DataHandler.Model(&MerchantApply{}).
			Where(cond).
			UpdateColumn("status", status).
			UpdateColumn("merchant_id", merchantId).
			UpdateColumn("merchant_key", merchantKey).Error

	case "reject":
		log.Println("apply reject")
		cond := map[string]interface{}{
			"id": m.Id,
		}
		err = DataHandler.Model(&MerchantApply{}).
			Where(cond).
			UpdateColumn("status", status).Error
	}

	if err != nil {
		return err
	} else {
		// 保存成功可以进行消息通知操作
		// TODO send to mq
		log.Println("send to mq")
		return nil
	}
}
