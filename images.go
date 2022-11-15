package model

type ImageBase struct {
	// Name 金融名称
	Name string `json:"name"`
	// Size 图片大小
	Size uint `json:"size"`
	// Type 图片类型
	Type string `json:"type" gorm:"default:avtar" gorm:"index:idx_type"`
	// Url 图片地址
	Url uint `json:"url"`
}

// Images 图片
type Images struct {
	BaseModel

	ImageBase

	// CreatedAt 创建人
	CreatedBy string `json:"created_by"`
	// UpdatedAt 修改人
	UpdatedBy string `json:"updated_by"`

	// MenuId 菜品id
	MenuId string `json:"menu_id"  gorm:"index:idx_menu_id"`
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m *Images) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("images", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("images", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m *Images) GetOne(instance interface{}) error {
	err := m.getOne("images", instance)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除记录
func (m *Images) Delete() error {
	err := m.delete("images")
	if err != nil {
		return err
	}
	return nil
}

// Save 删除记录
func (m *Images) Save() error {
	err := m.save(m)
	if err != nil {
		return err
	}
	return nil
}
