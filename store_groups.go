package model

// StoreGroup 分组
type StoreGroup struct {
	BaseModel

	// CreatedAt 创建人
	CreatedBy string `json:"created_by"`
	// UpdatedAt 修改人
	UpdatedBy string `json:"updated_by"`
	// Name 组名
	Name string `json:"name"`
}
