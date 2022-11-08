package model

import (
	"encoding/json"
	"strings"
)

// CustomerGroups 客户分组
type CustomerGroups struct {
	BaseModel

	Name     string   `json:"name" gorm:"index:name"`
	Segments string   `json:"-"`
	Groups   []string `json:"groups" gorm:"-"`
}

// Save 保存实例
func (m CustomerGroups) Save() {
	DataHandler.Create(&m)
}

// All 获取所有数据
func (m CustomerGroups) All(instance interface{}) error {
	err := m.all("customer_groups", instance)
	if err != nil {
		return err
	}
	return nil
}

func (m CustomerGroups) MarshalJSON() ([]byte, error) {
	// 命名别名，避免MarshalJson死循环
	type Alias CustomerGroups
	if m.Segments == "" {
		m.Groups = strings.Split(m.Segments, ",")
	} else {
		m.Groups = make([]string, 0)
	}
	return json.Marshal(struct {
		Alias
	}{Alias(m)})
}
