package model

// Categories 商品分类
type Categories struct {
	BaseModel

	Name string `json:"name"`

	Status string `json:"status"`
}

// All 获取所有数据
func (m Categories) All(instance interface{}) error {
	err := m.all("categories", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m Categories) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("categories", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("categories", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m Categories) GetOne(instance interface{}) error {
	err := m.getOne("categories", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByFilterOffset 获取所有数据
// 以便管理员进行审核操作
func (m Categories) ListByFilterOffset(instance interface{}, filter []string, offset int, limit int) (int64, error) {
	var counter int64 = 0

	filterParams := make([]string, 0)
	filterParams = append(filterParams, m.Status)

	err := m.counterByFilter("categories", &counter, filter, filterParams)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}
	// 获取列表
	err = m.listByFilterOffset("categories", instance, offset, limit, filter, filterParams)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// Delete 删除记录
func (m Categories) Delete() error {
	err := m.delete("categories")
	if err != nil {
		return err
	}
	return nil
}

// Update 更新
func (m Categories) Update(newOne Categories, columns []string) error {
	err := m.update("categories", newOne, columns)
	if err != nil {
		return err
	}
	return nil
}
