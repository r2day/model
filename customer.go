package model

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Customers 客户账号信息
type Customers struct {
	BaseModel

	// CardId 卡号
	CardId string `json:"card_id" gorm:"index:idx_card_id,unique"`
	// CustomerId 客户编号
	CustomerId string `json:"customer_id"`
	// 手机号
	Phone string `json:"phone" gorm:"index:idx_phone"`
	// 姓名
	Name string `json:"name"`
	// 性别
	Gender string `json:"gender"`
	// 出生日期
	BirthDay string `json:"birth_day"`
	// 资产
	Assets

	// 卡信息
	Cards
	// Avatar 头像基本地址
	// 例如: https://avatar.r2day.club/<customer_id>/64x64
	Avatar string `json:"avatar"`
	// Segments 分类标签 逗号分隔的数据
	Segments string   `json:"-"`
	Groups   []string `json:"groups" gorm:"-"`
}

func (m Customers) MarshalJSON() ([]byte, error) {
	// 命名别名，避免MarshalJson死循环
	type AliasCustomer Customers
	if m.Segments != "" {
		m.Groups = strings.Split(m.Segments, ",")
	} else {
		m.Groups = make([]string, 0)
	}
	return json.Marshal(struct {
		AliasCustomer
	}{AliasCustomer(m)})
}

// SaveALine 保存实例
func (m Customers) SaveALine(value []string) {

	m.CardId = value[0]
	m.CustomerId = value[1]
	m.Phone = value[2]
	m.Name = value[3]
	m.Gender = value[4]
	m.BirthDay = value[5]
	m.CardType = value[6]
	m.CardStatus = value[7]
	m.CardLevel = value[8]
	m.CardFrom = value[9]
	m.Balance, _ = strconv.ParseFloat(value[10], 64)
	m.CashCharge, _ = strconv.ParseFloat(value[11], 64)
	m.Freezing, _ = strconv.ParseFloat(value[12], 64)
	m.Gift, _ = strconv.ParseFloat(value[13], 64)
	m.Integral, _ = strconv.ParseUint(value[14], 10, 64)
	m.TotalBalance, _ = strconv.ParseFloat(value[15], 64)
	m.TotalBalanceCounter, _ = strconv.ParseUint(value[16], 10, 64)

	m.TotalCumulativeConsumption, _ = strconv.ParseFloat(value[17], 64)
	m.TotalCumulativeConsumptionCounter, _ = strconv.ParseUint(value[18], 10, 64)
	m.DebitTotalLimit, _ = strconv.ParseFloat(value[19], 64)
	m.DebitLeftLimit, _ = strconv.ParseFloat(value[20], 64)
	m.DebitUsedLimit, _ = strconv.ParseFloat(value[21], 64)
	m.EntityCardId = value[22]
	m.CardCreatedDate = value[23]
	m.Expire = value[24]

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
