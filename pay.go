package model

import (
	"errors"
	"time"

	logger "github.com/r2day/base/log"
	"github.com/r2day/base/util"
	"gorm.io/gorm"
)

// OrderItem 购物车物品
// 存储: mysql/es
// 写入: 客户
// 读: 客户/管理员
// 高频: 读
// 说明: 可以分析商品在购物车中的组合规律
type PaymentFlow struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// UserId 用户ID (登录dash平台的人)
	UserId string `json:"user_id"`
	// MerchantId 商户ID (例如: 黄李记作为一个商户存在)
	MerchantId string `json:"merchant_id"`
	// Status 状态
	Status string `gorm:"default:effected"`
	// CreatedAt 创建时间
	CreatedAt time.Time
	// UpdatedAt 修改时间
	UpdatedAt time.Time

	// 订单id (主订单id)
	OrderId string `json:"order_id" gorm:"order_id"`
	// OrderId
	AccountId string `json:"account_id" gorm:"account_id"`
	// 支付金额
	Amount float64 `json:"amount" gorm:"amount" `
	// 币种
	// Fkind enum.Fkind `json:"fkind" gorm:"fkind"`
}

// Pay 支付
// 支付方式:
// balance: 余额; wechat: 微信; ...
func (m PaymentFlow) Pay(payMethod string) (string, error) {

	// 交易流水号
	transactionId := util.TransactionId()

	err := DataHandler.Transaction(func(tx *gorm.DB) error {
		// step01 查询购物车的信息

		// switch payMethod {
		// case enum.Balance:
		//
		// }
		// 1. 查询账号余额
		fin := Finance{}
		cond1 := map[string]interface{}{
			"account_id": m.AccountId,
		}
		err := tx.Model(&AccountInfo{}).Where(cond1).First(&fin).Error
		if err != nil {
			logger.Logger.WithField("cond", cond1).
				WithError(err)
			return err
		}

		orderInfo := Order{}
		cond2 := map[string]interface{}{
			"order_id": m.MerchantId,
		}

		// 2. 查询当前订单的金额币种信息
		err = tx.Model(&Order{}).Where(cond2).First(&orderInfo).Error
		if err != nil {
			logger.Logger.WithField("cond", cond2).
				WithError(err)
			return err
		}
		if orderInfo.ActuallyPaid > fin.Balance {
			return errors.New("balance no enough to pay")
		}

		// step06 更新账号余额
		tx.Model(&Finance{}).
			Where(cond1).
			UpdateColumn("balance", gorm.Expr("balance - ?", orderInfo.ActuallyPaid))

		// 写流水 (先写到mysql 后续会同步到es并且删除mysql的流水)
		m.Amount = orderInfo.ActuallyPaid
		// m.Fkind = enum.
		DataHandler.Create(&m)
		// 发消息...
		return nil
	})

	if err != nil {
		return transactionId, err
	}
	return transactionId, nil
}
