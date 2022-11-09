package model

// Reviews 客户评价
type Reviews struct {
	BaseModel

	// 门店id
	StoreId uint `json:"store_id"`
	// 产品id
	ProductId uint `json:"product_id"`
	// 客户id
	CustomerId uint `json:"customer_id"`
	// 星级
	Rating uint `json:"rating"`
	// 用户评论
	Comment string `json:"comment"`
	// 评论状态
	Status string `json:"status"`
}

// All 获取所有数据
func (m Reviews) All(instance interface{}) error {
	err := m.all("reviews", instance)
	if err != nil {
		return err
	}
	return nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m Reviews) GetOne(instance interface{}) error {
	err := m.getOne("reviews", instance)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除记录
func (m Reviews) Delete() error {
	err := m.delete("reviews")
	if err != nil {
		return err
	}
	return nil
}

// Update 更新
func (m Reviews) Update(newOne Reviews, columns []string) error {
	err := m.update("reviews", newOne, columns)
	if err != nil {
		return err
	}
	return nil
}
