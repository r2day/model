package model

import (
	"time"

	logger "github.com/r2day/base/log"
	"gorm.io/gorm"
)

type CartModel struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// UserId 用户ID (登录dash平台的人)
	AdminId string `json:"admin_id"`
	// MerchantId 商户ID (例如: 黄李记作为一个商户存在)
	MerchantId string `json:"merchant_id"`
	// Status 状态
	Status string `gorm:"default:effected"`
	// CreatedAt 创建时间
	CreatedAt time.Time
	// UpdatedAt 修改时间
	UpdatedAt time.Time

	StoreId       string  `json:"store_id" gorm:"store_id"`
	UserId        string  `json:"user_id" gorm:"user_id"`
	ProductId     string  `json:"product_id" gorm:"product_id" `
	ProductName   string  `json:"product_name" gorm:"product_name"`
	ProductNumber int     `json:"product_number" gorm:"product_number"`
	TotalPrice    float32 `json:"total_price" gorm:"total_price"`
	UnitPrice     float32 `json:"unit_price" gorm:"unit_price"`
	Pic           string  `json:"pic" gorm:"pic"`
	// 特性
	Characteristic string `json:"characteristic" gorm:"characteristic"`
}

type CartOutputModel struct {
	// gorm.Model
	TotalProductNumber int `json:"total_product_number" gorm:"total_product_number"`
	TotalProductPrice  int `json:"total_product_price" gorm:"total_product_price" `
}

// Save 保存实例
func (m CartModel) Save() error {

	var cartInfoModel CartModel

	err := DataHandler.Transaction(func(tx *gorm.DB) error {

		// 查询条件
		cond := map[string]interface{}{
			"merchant_id": m.MerchantId,
			"store_id":    m.StoreId,
			"product_id":  m.ProductId,
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
		if cartInfoModel.ProductNumber == 0 {
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
				UpdateColumn("total_price", gorm.Expr("total_price + ?", m.TotalPrice)).
				UpdateColumn("product_number", gorm.Expr("product_number + ?", m.ProductNumber))
			logger.Logger.Info("increment a cart number successful")
		}
		return nil
	})

	return err

}

func (m CartModel) GetCurrentCartInfo() (CartOutputModel, []*CartModel, error) {
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
		Where("user_id = ?", m.UserId).Find(&cartOutputModel)

	DataHandler.Debug().Table("cart_models").
		Select("user_id, product_id, product_name, pic, unit_price, product_number, characteristic, total_price").
		Where(cond).Find(&cartListOutputModel)

	return cartOutputModel, cartListOutputModel, nil
}

func (m CartModel) MinusCart() error {
	const productNumber = 1 // 每次减1

	// 查询条件
	condForDeleted := map[string]interface{}{
		"merchant_id": m.MerchantId,
		"store_id":    m.StoreId,
		"user_id":     m.UserId,
		"product_id":  m.ProductId,
		"product_number": 1,
	}

	condForDecrement := map[string]interface{}{
		"merchant_id": m.MerchantId,
		"store_id":    m.StoreId,
		"user_id":     m.UserId,
		"product_id":  m.ProductId,
	}
	// 当product_number是1时，删除记录
	DataHandler.Debug().Table("cart_models").
		Where(condForDeleted).
		Delete(&CartModel{})

	DataHandler.Debug().Table("cart_models").
		Where(condForDecrement).
		UpdateColumn("total_price", gorm.Expr("total_price - ?", m.UnitPrice)).
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
