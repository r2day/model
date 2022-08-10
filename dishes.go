package model

import (
	"time"

	logger "github.com/r2day/base/log"
)

// Dishes 菜单库
type Dishes struct {
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

	// Name 业务名称
	ItemName string `json:"item_name"`
	// ItemId 项目编号(也是唯一值)
	ItemId string `json:"item_id"` // db func

	// 以下为业务具体的字段
	NormalPrice float32 `json:"normal_price"`
	// 会员价格
	VipPrice float32 `json:"vip_price"`
	// 分类名称
	CategoryName string `json:"category_name"`
	// 分类ID
	CategoryId string `json:"category_id"`
	// 菜品规格
	UnitName string `json:"unit_name"`
	// 菜品规格
	UnitId string `json:"unit_id"`
	// 品牌名称
	BrandName string `json:"brand_name"`
	// 品牌ID
	BrandId string `json:"brand_id"`
	// 部门名称
	DepName string `json:"dep_name"`
	// 部门ID
	DepId string `json:"dep_id"`
	// 是否支持外卖
	SupportTakeout string `json:"support_takeout"`
	// 标签
	Badge string `json:"badge"`
	// 描述
	Desc string `json:"desc"`
	// 图片
	Pic string `json:"pic"`
}

// Save 保存实例
func (m Dishes) Save() {
	DataHandler.Create(&m)
	Category{}.Increment(m.CategoryId)
}

// UpdateStatus 更新状态
func (m Dishes) UpdateStatus(id int, status string) {
	DataHandler.UpdateColumn("status", status).
		Where("id = ?", id)
}

// Delete 删除
func (m Dishes) Delete(id int) Dishes {
	// TODO 以后放到一个事务中执行
	// 删除前先查出记录
	DataHandler.
		Where("id = ?", id).First(&m)
	// 从类中删除计数器
	Category{}.Decrement(m.CategoryId)
	// 执行删除
	DataHandler.Where("id = ?", id).Delete(&m)
	return m
}

// GroupCategoryId 保存实例
func (m Dishes) GroupCategoryId(categoryId string) []Dishes {
	instance := make([]Dishes, 0)
	DataHandler.Where("category_id = ?", categoryId).Find(&instance)
	return instance
}

// Detail 菜品详情
func (m Dishes) Detail (productId string) (*Dishes, error) {

	// 查询条件
	cond := map[string]interface{}{
		"merchant_id": m.MerchantId,
		"item_id":  productId,
	}

	// var productInfoModel ProductInfoModel
	err := DataHandler.Where(cond).
	First(&m).Error
	
	if err != nil {
		// 返回任何错误都会回滚事务
		logger.Logger.WithField("m", m).
			WithError(err)
		return nil, err
	}
	return &m, nil
}