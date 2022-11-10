package model

// Products 产品信息
type Products struct {
	BaseModel

	CategoryId  uint    `json:"category_id"`
	Reference   string  `json:"reference"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	Price       float64 `json:"price"`
	Thumbnail   string  `json:"thumbnail"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Stock       uint    `json:"stock"`

	Status string `json:"status"`
}

// All 获取所有数据
func (m Products) All(instance interface{}) error {
	err := m.all("products", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m Products) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("products", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("products", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m Products) GetOne(instance interface{}) error {
	err := m.getOne("products", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByFilterOffset 获取所有数据
// 以便管理员进行审核操作
func (m Products) ListByFilterOffset(instance interface{}, filter []string, offset int, limit int) (int64, error) {
	var counter int64 = 0

	filterParams := make([]string, 0)
	filterParams = append(filterParams, m.Status)

	err := m.counterByFilter("products", &counter, filter, filterParams)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}
	// 获取列表
	err = m.listByFilterOffset("products", instance, offset, limit, filter, filterParams)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// Delete 删除记录
func (m Products) Delete() error {
	err := m.delete("products")
	if err != nil {
		return err
	}
	return nil
}

// Update 更新
func (m Products) Update(newOne Products, columns []string) error {
	err := m.update("products", newOne, columns)
	if err != nil {
		return err
	}
	return nil
}
