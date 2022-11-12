package model

// Brands 品牌
type Brands struct {
	BaseModel

	// CreatedAt 创建人
	CreatedBy string `json:"created_by"`
	// UpdatedAt 修改人
	UpdatedBy string `json:"updated_by"`
	// BrandId 品牌ID
	BrandId string `json:"brand_id"`
	// Name 品牌名称
	Name string `json:"name"`
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m *Brands) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("brands", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("brands", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m *Brands) GetOne(instance interface{}) error {
	err := m.getOne("brands", instance)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除记录
func (m *Brands) Delete() error {
	err := m.delete("brands")
	if err != nil {
		return err
	}
	return nil
}

// Save 删除记录
func (m *Brands) Save() error {
	err := m.save(m)
	if err != nil {
		return err
	}
	return nil
}
