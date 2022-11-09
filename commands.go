package model

import (
	"encoding/json"
)

// id: integer
//reference: string
//date: date
//customer_id: integer
//basket: [{ product_id: integer, quantity: integer }]
//status: 'ordered' | 'delivered' | 'canceled'
//returned: boolean

// BaseFee 基本费
type BaseFee struct {
	// 总额外税
	TotalExTaxes float64 `json:"total_ex_taxes"`
	// 运输费用
	DeliveryFees float64 `json:"delivery_fees"`
	// 费率
	TaxRate float64 `json:"tax_rate"`
	// 税
	Taxes float64 `json:"taxes"`
	// 总金额
	Total float64 `json:"total"`
}

// Product 商品
type Product struct {
	// 商品id
	ProductId uint `json:"product_id"`
	// 商品数量
	Quantity uint `json:"quantity"`
	// 商品单价
	UnitPrice uint `json:"unit_price"`
}

// Commands 订单信息
type Commands struct {
	BaseModel

	// CustomerId 客户编号(id)
	CustomerId uint `json:"customer_id"`

	// Reference 订单
	Reference string `json:"reference"`

	// 基本费用
	BaseFee

	// 商品序列号后的信息
	Baskets string `json:"-"`
	// 用于突出对外商品信息
	Basket []Product `json:"basket" gorm:"-"`

	// 订单状态
	Status string `json:"status"`
	// 是否为退款
	Returned bool `json:"returned"`
}

func (m Commands) MarshalJSON() ([]byte, error) {
	// 命名别名，避免MarshalJson死循环
	type AliasCommands Commands
	err := json.Unmarshal([]byte(m.Baskets), m.Basket)
	if err != nil {
		return nil, err
	}

	return json.Marshal(struct {
		AliasCommands
	}{AliasCommands(m)})
}

// All 获取所有数据
func (m Commands) All(instance interface{}) error {
	err := m.all("commands", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m Commands) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("commands", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("commands", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m Commands) GetOne(instance interface{}) error {
	err := m.getOne("commands", instance)
	if err != nil {
		return err
	}
	return nil
}
