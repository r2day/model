package model

import (
	"time"

	logger "github.com/r2day/base/log"
	"gorm.io/gorm"
)

// Sales 销量
type Sales struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// Status 状态
	Status string `gorm:"default:effected"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	// UpdatedAt 修改时间
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`

  // Id 自增唯一id
  ItemId string `json:"item_id" gorm:"item_id"`
  // 总销量
  TotalSales int64 `json:"total_sales" gorm:"total_sales"`
  // 总销量
  // TotalSales int64 `json:"total_sales" gorm:"total_sales"`
}
