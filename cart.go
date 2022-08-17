package model

import (
	"time"

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

	// 购物车id
	CartId string `json:"cart_id" gorm:"cart_id"`
	// 该商品数量
	Count int `json:"count" gorm:"count"`
	// 该商品总金额
	Amount float64 `json:"amount" gorm:"amount" `
	// 币种
	Currency string `json:"currency" gorm:"currency"`
	// 商品详情信息
	ItemInfo Item `json:"item_info" gorm:"item_info"`
}

func (m *CartItem) GetAmount() float64 {
	return float64(m.Count) * m.ItemInfo.Price
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
	// 商品总金额
	TotalAmount float64 `json:"total_amount" gorm:"total_amount"`
	// 商品总数
	TotalCount float64 `json:"total_count" gorm:"total_count"`
	// 购物项
	CartItems []*CartItem `json:"cart_items" gorm:"cart_items"`
}

type CartOutputModel struct {
	TotalProductNumber int     `json:"total_product_number" gorm:"total_product_number"`
	TotalProductPrice  float32 `json:"total_product_price" gorm:"total_product_price" `
}

// Save 保存实例
func (m Cart) Save() error {

	var cartInfoModel CartModel

	err := DataHandler.Transaction(func(tx *gorm.DB) error {

		// 查询条件
		cond := map[string]interface{}{
			"merchant_id": m.MerchantId,
			"store_id":    m.StoreId,
			"user_id":     m.UserId,
		}

		// 获取当前购物车列表
		err := tx.Model(&CartModel{}).Where(cond).First(&cartInfoModel).Error
		if err != nil {
			logger.Logger.WithField("cond", cond).
				WithError(err)
			// return err
		}

		// 还不存在则创建一个
		if cartInfoModel.TotalCount == 0 {
			logger.Logger.Info("ready to create a new object")
			// 单个商品首次添加
			if err := tx.Create(&m).Error; err != nil {
				// 返回任何错误都会回滚事务
				logger.Logger.WithField("m", m).
					WithError(err)
				return err
			}
			// 返回 nil 提交事务
			logger.Logger.Info("create a new object successful")
			return nil
		} else {
			logger.Logger.Info("ready to increment a cart number")
			tx.Model(&CartModel{}).
				Where(cond).
				UpdateColumn("total_amount", gorm.Expr("total_amount + ?", m.TotalAmount)).
				UpdateColumn("total_count", gorm.Expr("total_count + ?", m.TotalCount))
			logger.Logger.Info("increment a cart number successful")
		}
		return nil
	})

	return err

}

func (m Cart) GetCurrentCartInfo(item Item) (CartOutputModel, []*CartModel, error) {
	var cartOutputModel CartOutputModel
	cartListOutputModel := make([]*CartModel, 0)

	// 查询条件
	cond := map[string]interface{}{
		"merchant_id": m.MerchantId,
		"store_id":    m.StoreId,
		"user_id":     m.UserId,
	}

	// 查询当前购物车的状态
	DataHandler.Debug().Table("cart_models").
		Select("sum(product_number) as total_product_number, sum(total_price) as total_product_price").
		Where(cond).Find(&cartOutputModel)

	DataHandler.Debug().Table("cart_models").
		Select("user_id, product_id, product_name, pic, unit_price, product_number, characteristic, total_price, store_id, merchant_id, created_at, updated_at, id").
		Where(cond).Find(&cartListOutputModel)

	return cartOutputModel, cartListOutputModel, nil
}

func (m Cart) MinusCart() error {
	const productNumber = 1 // 每次减1

	// 查询条件
	condForDeleted := map[string]interface{}{
		"merchant_id": m.MerchantId,
		"store_id":    m.StoreId,
		"user_id":     m.UserId,
		// "product_number": 1,
	}

	condForDecrement := map[string]interface{}{
		"merchant_id": m.MerchantId,
		"store_id":    m.StoreId,
		"user_id":     m.UserId,
		// "product_id":  m.ProductId,
	}
	// 当product_number是1时，删除记录
	DataHandler.Debug().Table("cart").
		Where(condForDeleted).
		Delete(&CartModel{})

	DataHandler.Debug().Table("cart").
		Where(condForDecrement).
		// UpdateColumn("total_price", gorm.Expr("total_price - ?", m.UnitPrice)).
		UpdateColumn("product_number", gorm.Expr("product_number - ?", productNumber))
	return nil
}

// All 保存实例
func (m CartModel) All(merchantId string, storeId string, userId string) []CartModel {
	// 查询条件
	cond := map[string]interface{}{
		"merchant_id": m.MerchantId,
		"store_id":    m.StoreId,
		"user_id":     m.UserId,
	}
	instance := make([]CartModel, 0)
	DataHandler.Where(cond).
		Find(&instance)
	return instance
}
