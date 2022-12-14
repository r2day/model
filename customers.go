package model

import (
	"strconv"
)

// Customers 客户账号信息
type Customers struct {
	BaseModel

	// CustomerId 客户编号
	CustomerId string `json:"customer_id"`
	// 姓名
	Name string `json:"name"`
	// 性别
	Gender string `json:"gender"`
	// 手机号
	Phone string `json:"phone" gorm:"index:idx_phone"`
	// 会有卡(正常)
	VipCard uint64 `json:"vip_card"`
	// 优惠券
	Tickets uint64 `json:"tickets"`
	// 来源
	From string `json:"from"`
	// 注册时间
	RegisterDate string `json:"register_date"`
	// 生日类型 阳历/阴历
	BirthType string `json:"birth_type"`
	// 出生日期
	BirthDay string `json:"birth_day"`
}

// SaveALine 保存实例
func (m Customers) SaveALine(value []string) {

	// m.IndexId = value[0]
	m.CustomerId = value[1]
	m.Name = value[2]
	m.Gender = value[3]
	m.Phone = value[4]
	m.VipCard, _ = strconv.ParseUint(value[5], 10, 64)
	m.Tickets, _ = strconv.ParseUint(value[6], 10, 64)
	m.From = value[7]
	m.RegisterDate = value[8]
	m.BirthType = value[9]
	m.BirthDay = value[10]

	DataHandler.Create(&m)
}

// All 获取所有数据
func (m Customers) All(instance interface{}) error {
	err := m.all("customers", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m Customers) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("customers", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("customers", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m Customers) GetOne(instance interface{}) error {
	err := m.getOne("customers", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByFilterOffset 获取所有数据
// 以便管理员进行审核操作
func (m Customers) ListByFilterOffset(instance interface{}, filter []string, filterParams []string, offset int, limit int) (int64, error) {
	var counter int64 = 0
	err := m.counterByFilter("customers", &counter, filter, filterParams)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}
	// 获取列表
	err = m.listByFilterOffset("customers", instance, offset, limit, filter, filterParams)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// Delete 删除记录
func (m Customers) Delete() error {
	err := m.delete("customers")
	if err != nil {
		return err
	}
	return nil
}
