package model

import (
	"errors"
	"time"

	logger "github.com/r2day/base/log"
	btime "github.com/r2day/base/time"
	"github.com/r2day/base/util"
	"github.com/r2day/enum"
	"gorm.io/gorm"
)

// OrderItem 购物车物品
// 存储: es
// 写入: 客户
// 读: 客户/管理员
// 高频: 读
// 说明: 可以分析商品在购物车中的组合规律
type OrderItem struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// Status 状态
	Status string `gorm:"default:effected"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	// UpdatedAt 修改时间
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`

	// 订单id (主订单id)
	OrderId string `json:"order_id" gorm:"order_id"`
	// OrderId
	OrderItemId string `json:"order_item_id" gorm:"order_item_id"`
	// Itemid TODO 索引
	ItemId string `json:"item_id" gorm:"item_id"`
	// 该商品数量
	Count int `json:"count" gorm:"count"`
	// 该商品总金额
	Amount float64 `json:"amount" gorm:"amount" `
	// 币种
	Currency string `json:"currency" gorm:"currency"`
	// 商品详情信息
	Item Item `json:"item" gorm:"item"`
}

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
	// StoreId 店铺id
	StoreId string `json:"store_id" gorm:"store_id"`
	// 用户id
	UserId string `json:"user_id" gorm:"user_id"`

	// 订单列表
	ItemList []OrderItem `json:"item_list" gorm:"item_list"`

	// 保温袋
	PackageFee string `json:"package_fee" gorm:"package_fee"`
	// 配送费
	DeliveryFee string `json:"delivery_fee" gorm:"delivery_fee"`
	// 支付方式
	PaymentMethod string `json:"payment_method" gorm:"payment_method"`
	// 金额
	Balance float64 `json:"balance" gorm:"balance"`
	// 实付
	ActuallyPaid float64 `json:"actually_paid" gorm:"actually_paid"`

	// 订单信息
	// 订单状态(1: 下单成功；
	// 2: 已经支付；
	// 3: 仓库无货;
	// 4: 正在排队制作;
	// 5: 已经制作完成;
	// 6: 等待配送
	// 7: 外卖已经接单
	// 8: 已经送达
	// 9: 已经完成
	// 10: 发起退款
	// 11: 进入退款流程
	OrderStatus enum.OrderStatusEnum `json:"order_status" gorm:"order_status"`
	// 下单时间 2022.08.14 15:27
	OrderTime string `json:"order_time" gorm:"order_time"`
	// 下单门店
	StoreName string `json:"store_name" gorm:"store_name"`
	// 订单号
	OrderId string `json:"order_id" gorm:"order_id"`
	// 收货地址
	Address string `json:"address" gorm:"address"`
	// 取单号 (根据门店每天自动生成) 5001
	Seq int `json:"seq" gorm:"seq"`
	// 就餐方式 (堂食、外卖)
	WayOfEating string `json:"way_of_eating" gorm:"way_of_eating"`
	// 取餐时间
	PickUpTime string `json:"pick_up_time" gorm:"pick_up_time"`
	// 备注
	Remark string `json:"remark" gorm:"remark"`
}

func (m Order) PlaceOrder() error {
	err := DataHandler.Transaction(func(tx *gorm.DB) error {

		theOrderId := util.GetOrderId()
		// step01 查询购物车的信息
		cart := Cart{}
		cond := map[string]interface{}{
			"merchant_id": m.MerchantId,
			"store_id":    m.StoreId,
			"user_id":     m.UserId,
		}
		// 获取当前购物车列表
		err := tx.Model(&Cart{}).Where(cond).First(&cart).Error
		if err != nil {
			logger.Logger.WithField("cond", cond).
				WithError(err)
			return err
		}
		if cart.TotalCount == 0 {
			return errors.New("nothing need to place, please go to add something into cart")
		}

		// step02 查询购物车中的商品信息
		cartItem := make([]CartItem, 0)
		cond2 := map[string]interface{}{
			"cart_id": cart.CartId,
		}
		// 获取当前购物车列表
		err = tx.Model(&CartItem{}).Where(cond2).Find(&cartItem).Error
		if err != nil {
			logger.Logger.WithField("cond", cond).
				WithError(err)
			return err
		}
		// step03 将购物车的物品变为订单
		orderItem := make([]OrderItem, 0)
		for _, i := range cartItem {
			singleOrderItem := OrderItem{
				OrderId:     theOrderId,
				OrderItemId: util.GetOrderItemId(),
				ItemId:      i.ItemId,
				Count:       i.Count,
				Amount:      i.Amount,
				Currency:    i.Currency,
				Item:        i.Item,
			}
			// TODO 发出消息
			// 1. 厨房; 2. es; 3. 仓库; 4. 打印机 etc
			orderItem = append(orderItem, singleOrderItem)
		}

		DataHandler.Create(&orderItem)
		// step04 移除购物车的商品
		DataHandler.Debug().Table("cart_items").
			Where(cond2).
			Delete(&CartItem{})

		// step05 回填订单信息
		// order := Order{}

		// 这里可以增加会员计算优惠
		m.ActuallyPaid = cart.TotalAmount
		m.ItemList = orderItem
		m.OrderId = theOrderId
		m.OrderStatus = enum.Init
		m.OrderTime = btime.GetCurrentTime()

		DataHandler.Create(&m)

		// step06 更新购物车信息
		tx.Model(&Cart{}).
			Where(cond).
			UpdateColumn("total_amount", 0).
			UpdateColumn("total_count", 0)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// Save 保存实例
func (m Order) Save() error {
	return nil

}
