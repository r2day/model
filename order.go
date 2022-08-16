package model

import (
	"time"

	logger "github.com/r2day/base/log"
	"gorm.io/gorm"
)

type Order struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// UserId 用户ID (登录dash平台的人)
	AdminId string `json:"admin_id"`
	// MerchantId 商户ID (例如: 黄李记作为一个商户存在)
	MerchantId string `json:"merchant_id"`
	// Status 状态
	Status string `gorm:"default:effected"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	// UpdatedAt 修改时间
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`

	StoreId       string  `json:"store_id" gorm:"store_id"`
	UserId        string  `json:"user_id" gorm:"user_id"`
	ProductId     string  `json:"product_id" gorm:"product_id" `
	ProductName   string  `json:"product_name" gorm:"product_name"`
	ProductNumber int     `json:"product_number" gorm:"product_number"`
	TotalPrice    float32 `json:"total_price" gorm:"total_price"`
	UnitPrice     float32 `json:"unit_price" gorm:"unit_price"`
	Pic           string  `json:"pic" gorm:"pic"`
	// 特性
	Desc string `json:"desc" gorm:"desc"`
	// 订单信息

	// 下单时间
	OrderTime string  `json:"order_time" gorm:"order_time"`
	// 下单门店
	StoreName       string  `json:"store_name" gorm:"store_name"`
	// 订单号
	OrderId        string  `json:"order_id" gorm:"order_id"`
	// 收货地址
	Address        string  `json:"address" gorm:"address"`
	// 取单号 (根据门店每天自动生成)
	Seq        string  `json:"seq" gorm:"seq"`
	// 就餐方式 (堂食、外卖)
	WayOfEating string  `json:"way_of_eating" gorm:"way_of_eating"`
	// 取餐时间
	PickUpTime string  `json:"pick_up_time" gorm:"pick_up_time"`
	// 备注
	Remark string  `json:"remark" gorm:"remark"`
}


// Save 保存实例
func (m Order) Save() error{
	// 查询购物车
	goodsInCart := make([]CartModel, 0)

	err := DataHandler.Transaction(func(tx *gorm.DB) error {

		// 查询条件
		cond := map[string]interface{}{
			"merchant_id": m.MerchantId,
			"store_id":    m.StoreId,
			"user_id":     m.UserId,
		}

		// 获取当前购物车列表
		err := tx.Model(&CartModel{}).Where(cond).Find(&goodsInCart).Error
		if err != nil {
			logger.Logger.WithField("cond", cond).
				WithError(err)
			// return err
		}

		if len(goodsInCart) == 0 {
			// 购物车为空
			logger.Logger.Warn("cart is empty")
			return nil
		} else {
			// 将购物车的物品搬到订单中
			for _, i := range goodsInCart {
				logger.Logger.Info("ready to place")
				// 从仓库中扣除
				// 
				o := Order{
					AdminId: i.AdminId,
					MerchantId: i.MerchantId,
					Status: "place",
					StoreId: i.StoreId,
					UserId: i.UserId,
					ProductId: i.UserId,
					ProductName: i.ProductName,
					TotalPrice: i.TotalPrice,
				}
				logger.Logger.Info("TODO send order to mq -->", o)

				tx.Model(&CartModel{}).
					Where(cond).
					UpdateColumn("total_price", gorm.Expr("total_price + ?", m.TotalPrice)).
					UpdateColumn("product_number", gorm.Expr("product_number + ?", m.ProductNumber))
				logger.Logger.Info("increment a cart number successful")
			}
		}
		return nil
	})

	return err

}