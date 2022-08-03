package model

import (
	"errors"
	"time"
)

// Category 菜单分组
type Category struct {
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

	// 品牌名称
	BrandName string `json:"brand_name"`
	// 品牌ID
	BrandId string `json:"brand_id"`
	// 部门名称
	DepName string `json:"dep_name"`
	// 部门ID
	DepId string `json:"dep_id"`
	// 大分类
	BigCategory string `json:"big_category"`
	// 分类
	Type string `json:"type"`
	// 收入科目
	TakeIn string `json:"take_in"`
	// 排序用于展示效果
	SortIndex int `json:"sort_index"`
	// 自动增加(新增菜品库的时候)
	Ref int `json:"ref" gorm:"default:0"`
}

// All 保存实例
func (m Category) All(merchantId string) []Category {
	instance := make([]Category, 0)
	DataHandler.Where("merchant_id = ?", merchantId).Order("sort_index").Find(&instance)
	return instance
}

// LastOne 获取对象
func (m Category) LastOne(merchantId string) *Category {
	DataHandler.Where("merchant_id = ?", merchantId).Order("sort_index desc").First(&m)
	return &m
}

// GetByItemId 获取对象
func (m Category) GetByItemId(itemId string) *Category {
	DataHandler.Where("item_id = ?", itemId).First(&m)
	return &m
}

// GetByIndex 获取对象
func (m Category) GetByIndex(merchantId string, index int) *Category {
	DataHandler.Where("merchant_id = ? and sort_index = ?", merchantId, index).First(&m)
	return &m
}

// Increment 当新增该类型的菜品时会触发引用计数器加
func (m Category) Increment(itemId string) {
	DataHandler.Exec("UPDATE categories SET ref = ref + 1 WHERE item_id = ?", itemId)
}

// Decrement 当新增该类型的菜品时会触发引用计数器加
func (m Category) Decrement(itemId string) {
	DataHandler.Exec("UPDATE categories SET ref = ref - 1 WHERE item_id = ?", itemId)
}

func (m Category) SwapIndex(a string, b string, aIndex int, bIndex int) {
	DataHandler.Exec("UPDATE categories SET sort_index = ? WHERE item_id = ?", aIndex, a)
	DataHandler.Exec("UPDATE categories SET sort_index = ? WHERE item_id = ?", bIndex, b)
}

func (m Category) InsertItem(targetItemId string, sourceIndex int, DestinationIndex int) {

	DataHandler.Exec("UPDATE categories SET sort_index = ? WHERE item_id = ?", DestinationIndex, targetItemId)
	// 向下+1 往下移动
	if sourceIndex > DestinationIndex {
		DataHandler.Exec("UPDATE categories SET sort_index = sort_index + 1 WHERE sort_index between ? and ? and item_id != ?",
			DestinationIndex, sourceIndex, targetItemId)
	} else if sourceIndex < DestinationIndex {
		DataHandler.Exec("UPDATE categories SET sort_index = sort_index - 1 WHERE sort_index between ? and ? and item_id != ?",
			sourceIndex, DestinationIndex, targetItemId)
	}
}

// Delete 删除
func (m Category) Delete(id int) (Category, error) {
	// TODO 以后放到一个事务中执行
	// 删除前先查出记录
	DataHandler.
		Where("id = ? and ref = ?", id, 0).First(&m)

	if m.ItemId == "" {
		// 没有符合的分类
		// 分类下的菜品必选为空才能执行删除分类动作
		return m, errors.New("the ref is not zero, you can not delete category")
	}
	// 执行删除
	DataHandler.Where("item_id = ?", m.ItemId).Delete(&m)
	return m, nil
}
