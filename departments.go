package model

// Departments 部门
type Departments struct {
	BaseModel

	// CreatedAt 创建人
	CreatedBy string `json:"created_by"`
	// UpdatedAt 修改人
	UpdatedBy string `json:"updated_by"`
	// DepartmentId 部门id
	DepartmentId string `json:"department_id"`
	// Name 部门名称名称
	Name string `json:"name"`
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m *Departments) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("departments", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("departments", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m *Departments) GetOne(instance interface{}) error {
	err := m.getOne("departments", instance)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除记录
func (m *Departments) Delete() error {
	err := m.delete("departments")
	if err != nil {
		return err
	}
	return nil
}

// Save 删除记录
func (m *Departments) Save() error {
	err := m.save(m)
	if err != nil {
		return err
	}
	return nil
}
