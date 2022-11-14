package model

// MenuVariants 菜单做法
type MenuVariants struct {
	BaseModel

	// 对外提供列表
	MenuSpecification
	// CreatedAt 创建人
	CreatedBy string `json:"created_by"`
	// UpdatedAt 修改人
	UpdatedBy string `json:"updated_by"`
	// 菜品id (可以重复引用)
	MenuId string `json:"menu_id" gorm:"index:idx_menu_id"`
}

// All 获取所有数据
func (m MenuVariants) All(instance interface{}) error {
	err := m.all("menu_variants", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m *MenuVariants) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("menu_variants", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("menu_variants", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m *MenuVariants) GetOne(instance interface{}) error {
	err := m.getOne("menu_variants", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByFilterOffset 获取所有数据
// 以便管理员进行审核操作
func (m *MenuVariants) ListByFilterOffset(instance interface{}, filter []string, filterParams []string, offset int, limit int) (int64, error) {
	var counter int64 = 0
	err := m.counterByFilter("menu_variants", &counter, filter, filterParams)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}
	// 获取列表
	err = m.listByFilterOffset("menu_variants", instance, offset, limit, filter, filterParams)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// Delete 删除记录
func (m *MenuVariants) Delete() error {
	err := m.delete("menu_variants")
	if err != nil {
		return err
	}
	return nil
}

// Save 保存记录
func (m *MenuVariants) Save() error {
	err := m.save(m)
	if err != nil {
		return err
	}
	return nil
}
