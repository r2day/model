package model

import (
	"fmt"
	"time"
)

const (
	Sep = " and %s = ?"
)

// BaseModel 基本数据表模型
type BaseModel struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`

	// MerchantId 商户ID (例如: 黄李记作为一个商户存在)
	MerchantId string `json:"merchant_id" gorm:"index:idx_merchant"`

	// BaseStatus 基本状态
	BaseStatus string `json:"base_status" gorm:"default:effected" gorm:"index:idx_base_status"`

	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`

	// UpdatedAt 修改时间
	UpdatedAt time.Time `json:"updated_at"  gorm:"column:updated_at;autoUpdateTime"`
}

// Save 保存实例
func (m BaseModel) save(instance interface{}) {
	DataHandler.Create(instance)
}

// all 获取所有数据
func (m BaseModel) all(table string, instance interface{}) error {
	//instance := make([]BaseModel, 0)
	err := DataHandler.Table(table).Debug().
		Where("base_status = ? and merchant_id = ?", m.BaseStatus, m.MerchantId).
		Find(instance).Error
	if err != nil {
		return err
	}
	return nil
}

// listByOffset 根据分页规则获取所有数据
func (m BaseModel) listByOffset(table string,
	instance interface{}, offset int, limit int) error {
	err := DataHandler.Table(table).Debug().
		Where("base_status = ? and merchant_id = ?", m.BaseStatus, m.MerchantId).
		Offset(offset).Limit(limit).
		Find(instance).Error
	if err != nil {
		return err
	}
	return nil
}

// counter 获取数据记录数
func (m BaseModel) counter(table string, counter *int64) error {
	err := DataHandler.Table(table).Debug().
		Where("base_status = ? and merchant_id = ?", m.BaseStatus, m.MerchantId).
		Count(counter).Error
	if err != nil {
		return err
	}
	return nil
}

// counter 获取数据记录数
func (m BaseModel) counterByFilter(table string, counter *int64, filterColumns []string, filterParams interface{}) error {
	filter := joinQueryFields(filterColumns)
	err := DataHandler.Table(table).Debug().
		Where("base_status = ? and merchant_id = ?"+filter, m.BaseStatus, m.MerchantId, filterParams).
		Count(counter).Error
	if err != nil {
		return err
	}
	return nil
}

// GetMany 获取指定的客户信息
func (m BaseModel) getMany(table string, ids []string, instance interface{}) error {
	err := DataHandler.Table(table).
		Where("base_status = ? and merchant_id = ? and id IN ?",
			m.BaseStatus, m.MerchantId, ids).
		Find(instance).Error
	if err != nil {
		return err
	}
	return nil
}

// GetOne 获取单个详情
func (m BaseModel) getOne(table string, instance interface{}) error {
	err := DataHandler.Table(table).
		Where("base_status = ? and merchant_id = ? and id = ?",
			m.BaseStatus, m.MerchantId, m.Id).
		First(&instance).Error
	if err != nil {
		return err
	}
	return nil
}

// listByOffset 根据分页规则获取所有数据
// filter: ["status", "category_id"]
// filterParams: ["pending", 5]
func (m BaseModel) listByFilterOffset(table string,
	instance interface{}, offset int, limit int, filterColumns []string, filterParams interface{}) error {
	filter := joinQueryFields(filterColumns)
	err := DataHandler.Table(table).Debug().
		Where("base_status = ? and merchant_id = ?"+filter, m.BaseStatus, m.MerchantId, filterParams).
		Offset(offset).Limit(limit).
		Find(instance).Error
	if err != nil {
		return err
	}
	return nil
}

// delete 删除一条记录
func (m BaseModel) delete(table string) error {
	err := DataHandler.Table(table).
		Where("base_status = ? and merchant_id = ? and id = ?",
			m.BaseStatus, m.MerchantId, m.Id).
		Delete(&m).Error
	if err != nil {
		return err
	}
	return nil
}

// update 更新一条记录
func (m BaseModel) update(table string, newOne interface{}, columns []string) error {
	err := DataHandler.Table(table).
		Select(columns).
		Where("base_status = ? and merchant_id = ? and id = ?",
			m.BaseStatus, m.MerchantId, m.Id).
		Updates(newOne).Error
	if err != nil {
		return err
	}
	return nil
}

// joinQueryFields 拼接查询条件
func joinQueryFields(columns []string) string {
	final := ""
	for _, name := range columns {
		f := fmt.Sprintf(Sep, name)
		final += f
	}
	return final
}

// JoinQueryFields 拼接查询条件
func (m *BaseModel) JoinQueryFields(columns []string) string {
	return joinQueryFields(columns)
}
