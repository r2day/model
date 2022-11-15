package model

// MenuTaste 菜单做法
type MenuTaste struct {
	BaseModel

	// 对外提供列表
	MenuTasteEnable

	// CreatedAt 创建人
	CreatedBy string `json:"created_by"`
	// UpdatedAt 修改人
	UpdatedBy string `json:"updated_by"`
	// 菜品id (可以重复引用)
	MenuId string `json:"menu_id" gorm:"index:idx_menu_id"`
}

// All 获取所有数据
func (m MenuTaste) All(instance interface{}) error {
	err := m.all("menu_tastes", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m *MenuTaste) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("menu_tastes", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("menu_tastes", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m *MenuTaste) GetOne(instance interface{}) error {
	err := m.getOne("menu_tastes", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByFilterOffset 获取所有数据
// 以便管理员进行审核操作
func (m *MenuTaste) ListByFilterOffset(instance interface{}, filter []string, filterParams []string, offset int, limit int) (int64, error) {
	var counter int64 = 0
	err := m.counterByFilter("menu_tastes", &counter, filter, filterParams)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}
	// 获取列表
	err = m.listByFilterOffset("menu_tastes", instance, offset, limit, filter, filterParams)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// Delete 删除记录
func (m *MenuTaste) Delete() error {
	err := m.delete("menu_tastes")
	if err != nil {
		return err
	}
	return nil
}

// Save 保存记录
func (m *MenuTaste) Save() error {
	err := m.save(m)
	if err != nil {
		return err
	}
	return nil
}
