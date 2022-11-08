package model

// CustomerGroups 客户分组
type CustomerGroups struct {
	BaseModel

	Name string `json:"name" gorm:"index:name"`
}

// Save 保存实例
func (m CustomerGroups) Save() {
	DataHandler.Create(&m)
}

// All 获取所有数据
func (m CustomerGroups) All(instance []CustomerGroups) error {
	err := m.all("customer_groups", &instance)
	if err != nil {
		return err
	}
	return nil
}
