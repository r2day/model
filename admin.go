package model

import (
	"log"
	"time"
)

// AdminAccount 商户入驻申请
// 会进行增长的表后续都会进行数据迁移或者清理
type AdminAccount struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// CreatedAt 创建时间
	CreatedAt time.Time
	// UpdatedAt 修改时间
	UpdatedAt time.Time

	// Phone
	// 每个手机号仅能申请一次
	Phone string `json:"phone" gorm:"index:idx_phone,unique"`
	// Password 密码
	Password string `json:"password"`
}

// FindByPhone 通过手机号查询
// 审批通过后才有商户id
func (m AdminAccount) FindByPhone() (AdminAccount, error) {

	// 查询条件
	cond := map[string]interface{}{
		"phone": m.Phone,
	}
	err := DataHandler.Debug().Table("admin_accounts").
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
func (m AdminAccount) Save() error {
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
