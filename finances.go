package model

// Finances 品牌
type Finances struct {
	BaseModel

	// CreatedAt 创建人
	CreatedBy string `json:"created_by"`
	// UpdatedAt 修改人
	UpdatedBy string `json:"updated_by"`
	// FinanceId 金融编号
	FinanceId string `json:"finance_id"`
	// Name 金融名称
	Name string `json:"name"`
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m *Finances) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("finances", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("finances", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m *Finances) GetOne(instance interface{}) error {
	err := m.getOne("finances", instance)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除记录
func (m *Finances) Delete() error {
	err := m.delete("finances")
	if err != nil {
		return err
	}
	return nil
}

// Save 删除记录
func (m *Finances) Save() error {
	err := m.save(m)
	if err != nil {
		return err
	}
	return nil
}
