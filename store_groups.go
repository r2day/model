package model

// StoreGroup 分组
type StoreGroup struct {
	BaseModel

	// CreatedAt 创建人
	CreatedBy string `json:"created_by"`
	// UpdatedAt 修改人
	UpdatedBy string `json:"updated_by"`
	// GroupId 组ID
	GroupId string `json:"group_id"`
	// Name 组名
	Name string `json:"name"`
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m StoreGroup) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("store_groups", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("store_groups", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m StoreGroup) GetOne(instance interface{}) error {
	err := m.getOne("store_groups", instance)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除记录
func (m StoreGroup) Delete() error {
	err := m.delete("store_groups")
	if err != nil {
		return err
	}
	return nil
}

// Save 删除记录
func (m StoreGroup) Save() error {
	err := m.save(&m)
	if err != nil {
		return err
	}
	return nil
}
