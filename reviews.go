package model

// Reviews 客户评价
type Reviews struct {
	BaseModel

	// 门店id
	StoreId string `json:"store_id"`
	// 产品id
	ProductId string `json:"product_id"`
	// 客户id
	CustomerId string `json:"customer_id"`
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

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m Reviews) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("reviews", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("reviews", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// ListByFilterOffset 获取所有数据
// 以便管理员进行审核操作
func (m Reviews) ListByFilterOffset(instance interface{}, filter []string, filterParams []string, offset int, limit int) (int64, error) {
	var counter int64 = 0
	err := m.counterByFilter("reviews", &counter, filter, filterParams)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}
	// 获取列表
	err = m.listByFilterOffset("reviews", instance, offset, limit, filter, filterParams)
	if err != nil {
		return 0, err
	}
	return counter, nil
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
