package model

import (
	"time"

	"github.com/r2day/base/util"
	logger "github.com/r2day/base/log"
	"gorm.io/gorm"
)

// Item 物品
// 存储: mysql
// 写入: 管理员
// 读: 客户/管理员
// 高频: 读
type Item struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// Status 状态
	Status string `gorm:"default:effected"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	// UpdatedAt 修改时间
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`

	// Name 名称
	Name string `json:"name" gorm:"name"`
	// ItemId 商品编号
	ItemId string `json:"item_id" gorm:"item_id"`
	// Price 价格
	Price float64 `json:"price" gorm:"price"`
	// Currency 币种
	Currency string `json:"currency" gorm:"currency"`

	// Alias 别名
	Alias string `json:"alias" gorm:"alias"`
	// Badge 徽章 (新品、微辣、..)
	Badge string `json:"badge" gorm:"badge"`
	// Category 分类
	Category string `json:"category" gorm:"category"`
	// Pic 图片
	Pic string `json:"pic" gorm:"pic"`
	// Urls 更多信息
	Urls string `json:"urls" gorm:"urls"`
	// Desc 描述
	Desc string `json:"desc" gorm:"desc"`
	// Sales 销量信息 通过查询Sales.GetSales(itemId)
	SalesInfo Sales `json:"sales_info" gorm:"sales_info"`
}

// CartItem 购物车物品
// 存储: es
// 写入: 客户
// 读: 客户/管理员
// 高频: 读
// 说明: 可以分析商品在购物车中的组合规律
type CartItem struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// Status 状态
	Status string `gorm:"default:effected"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	// UpdatedAt 修改时间
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`

	// 购物车id 加索引?????
	CartId string `json:"cart_id" gorm:"cart_id"`
	// CartItemId
	CartItemId string `json:"cart_item_id" gorm:"cart_item_id"`
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

func (m *CartItem) CalculateAmount() float64 {
	return float64(m.Count) * m.Item.Price
}

// Cart 购物车
// 存储: mysql
// 写入: 客户
// 读: 客户/管理员
// 高频: 读
type Cart struct {
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

	// 店铺id
	StoreId string `json:"store_id" gorm:"store_id"`
	// 用户id
	UserId string `json:"user_id" gorm:"user_id"`
	// cart id
	CartId string `json:"cart_id" gorm:"cart_id"`
	// 商品总金额
	TotalAmount float64 `json:"total_amount" gorm:"total_amount"`
	// 商品总数
	TotalCount int `json:"total_count" gorm:"total_count"`
	// 购物项
	CartItems []*CartItem `json:"cart_items" gorm:"cart_items"`
}

type CartOutputModel struct {
	TotalProductNumber int     `json:"total_product_number" gorm:"total_product_number"`
	TotalProductPrice  float32 `json:"total_product_price" gorm:"total_product_price" `
}

// 如果不存在cart则创建一个，否则直接返回当前的cart
func (m Cart) Init() Cart {
	// 购物车初始化
	// 查询cartItem条件
	cond := map[string]interface{}{
		"merchant_id":     m.MerchantId,
		"store_id":     m.StoreId,
		"user_id":     m.UserId,
	}
	// 获取当前购物车列表
	err := DataHandler.Model(&Cart{}).Where(cond).First(&m).Error
	if err != nil {
		logger.Logger.WithField("cond", cond).
			WithError(err)
		return m
	}
	// 确认是否已经查询到有效的购物车
	// 如果不存在，则创建一个
	if m.Status != "effected" {
			DataHandler.Create(&m)
	}
	return m
}

// Save 购物车
// 单例模式
func (m Cart) Save(item Item) error {

	var cartItem CartItem

	// 查询条件
	cartCond := map[string]interface{}{
		"merchant_id": m.MerchantId,
		"store_id":    m.StoreId,
		"user_id":     m.UserId,
	}
	err := DataHandler.Transaction(func(tx *gorm.DB) error {

		// 查询cartItem条件
		cond := map[string]interface{}{
			"item_id":     item.ItemId,
		}

		// 获取当前购物车列表
		err := tx.Model(&CartItem{}).Where(cond).First(&cartItem).Error
		if err != nil {
			logger.Logger.WithField("cond", cond).WithError(err)
			// return err (if not record is ok, we create one, don't return it )
		}

		// 还不存在则创建一个cartItem
		if cartItem.Count == 0 {
			logger.Logger.Info("ready to create a new cartItem")

			// 赋值商品详情信息
			cartItem.CartItemId = util.GetCartItemId()
			cartItem.CartId = m.CartId
			cartItem.ItemId = item.ItemId
			cartItem.Item = item // 以后不再需要重复赋值
			cartItem.Count = 1 // 首次添加 (往后直接累加)
			cartItem.Amount = cartItem.CalculateAmount()
			cartItem.Currency = item.Currency

			// 单个商品首次添加
			if err := tx.Create(&cartItem).Error; err != nil {
				// 返回任何错误都会回滚事务
				logger.Logger.WithField("item", item).WithField("cartItem", cartItem).WithError(err)
				return err
			}
			// 返回 nil 提交事务
			logger.Logger.Info("create a new cartItem successful")
			// TODO send to MQ / ES save it

			return nil
		} else {
			logger.Logger.Info("ready to increment a cart number")
			// 更新cartItem
			tx.Model(&CartItem{}).
				Where(cond).
				UpdateColumn("count", gorm.Expr("count + ?", 1)).
				UpdateColumn("amount", gorm.Expr("amount + ?", item.Price))

			logger.Logger.Info("increment a cart number successful")
		}

		// 更新购物车
		tx.Model(&Cart{}).
			Where(cartCond).
			UpdateColumn("total_amount", gorm.Expr("total_amount + ?", item.Price)).
			UpdateColumn("total_count", gorm.Expr("total_count + ?", 1)) // 每次增加一个？
		return nil
	})

	return err

}

// GetCartInfo 获取购物车信息
func (m Cart) GetCartInfo() (Cart, []*CartItem, error) {
	var cart Cart
	cartItems := make([]*CartItem, 0)

	// 查询条件
	cond := map[string]interface{}{
		"merchant_id": m.MerchantId,
		"store_id":    m.StoreId,
		"user_id":     m.UserId,
	}

	// 查询当前购物车的状态
	DataHandler.Debug().Table("cart").
		Select("total_amount, total_count, cart_id").
		Where(cond).First(&cart)


		cond2 := map[string]interface{}{
			"cart_id": cart.CartId,
		}

	DataHandler.Debug().Table("cart_item").
		Select("count, amount, item_id").
		Where(cond2).Find(&cartItems)

	// 循环查询缓存获得item 详细信息
	// ....
	// for _, i := range cartItems {
	// 	// query from cache
	// 	i.Item = nil //....
	// }

	return cart, cartItems, nil
}

// MinusCart 从购物车中移除商品
func (m Cart) MinusCart(item Item) error {
	const productNumber = 1 // 每次减1



	// 查询条件
	cartItemCond := map[string]interface{}{
		"item_id":     item.ItemId,
		"count":     1, // 如果当前是1则直接移除
	}

	// 查询条件
	cartItemCond2 := map[string]interface{}{
		"item_id":     item.ItemId,
	}

	cartCond := map[string]interface{}{
		"merchant_id": m.MerchantId,
		"store_id":    m.StoreId,
		"user_id":     m.UserId,
	}
	// 当product_number是1时，删除记录
	DataHandler.Debug().Table("cart_item").
		Where(cartItemCond).
		Delete(&CartItem{})

	// 如果不为最后一个则，直接减1
	DataHandler.Debug().Table("cart_item").
		Where(cartItemCond2).
		UpdateColumn("amount", gorm.Expr("amount - ?", item.Price)).
		UpdateColumn("count", gorm.Expr("count - ?", 1))

	// 购物车总体减一次
	DataHandler.Debug().Table("cart").
		Where(cartCond).
		UpdateColumn("total_amount", gorm.Expr("total_amount - ?", item.Price)).
		UpdateColumn("total_count", gorm.Expr("total_count - ?", 1))
	return nil
}

// All 管理员查看当前所有人的购物车
func (m Cart) All() []Cart {
	cond := map[string]interface{}{
		"merchant_id": m.MerchantId,
	}
	instance := make([]Cart, 0)
	DataHandler.Where(cond).
		Find(&instance)
	return instance
}
