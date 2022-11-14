package model

// Incomes 收入科目
type Incomes struct {
	BaseModel

	// CreatedAt 创建人
	CreatedBy string `json:"created_by"`
	// UpdatedAt 修改人
	UpdatedBy string `json:"updated_by"`
	// IncomeId 收入科目id
	IncomeId string `json:"income_id"`
	// Name 金融名称
	Name string `json:"name"`
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m *Incomes) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("incomes", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("incomes", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m *Incomes) GetOne(instance interface{}) error {
	err := m.getOne("incomes", instance)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除记录
func (m *Incomes) Delete() error {
	err := m.delete("incomes")
	if err != nil {
		return err
	}
	return nil
}

// Save 删除记录
func (m *Incomes) Save() error {
	err := m.save(m)
	if err != nil {
		return err
	}
	return nil
}
