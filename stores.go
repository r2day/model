package model

import "strings"

type Position struct {
	// 位置信息
	Country           string `json:"country"`
	Province          string `json:"province"`
	City              string `json:"city"`
	Area              string `json:"area"`
	Street            string `json:"street"`
	Address           string `json:"address"`
	LatitudeLongitude string `json:"latitude_longitude"`
}

type StoresInfo struct {
	// 联系方式
	Phone string `json:"phone"`
	// 公告
	BBS string `json:"bbs"`
	// 营业时间
	BusinessHours string `json:"business_hours"`
	// 门店状态
	Status string `json:"status"`
}

// Stores 发票
type Stores struct {
	BaseModel

	// CreatedAt 创建人
	CreatedBy string `json:"created_by"`
	// UpdatedAt 修改人
	UpdatedBy string `json:"updated_by"`
	// GroupId 分组id
	GroupId string `json:"group_id"`
	// BrandName 品牌名称
	BrandName string `json:"brand_name"`
	// StoreName 店铺名称
	StoreName string `json:"store_name"`
	// StoreId 门店id
	StoreId string `json:"store_id" gorm:"index:idx_store_id,unique"`
	// OrganizeId
	OrganizeId string `json:"organize_id"`
	// FinanceName 财务主体
	FinanceName string `json:"finance_name"`
	// FinanceId 财务
	FinanceId string `json:"finance_id"`
	// Tag 标签
	Tag string `json:"tag"`
	// ManageOrganize
	ManageOrganize string `json:"manage_organize"`

	// 位置
	Position

	StoresInfo
	// 运营摸索
	BusinessMode string `json:"business_mode"`
	// 系统类型
	SystemType string `json:"system_type"`
	// 分组
	StoreGroupName string `json:"store_group_name"`
}

// SaveALine 保存实例
func (m Stores) SaveALine(value []string) {

	// 因为经纬度包含了，（逗号），因此为27columns
	// 否则可以忽略
	//if len(value) != 27 {
	//	return
	//}

	m.GroupId = value[0]
	m.BrandName = value[1]
	m.StoreName = value[2]
	m.StoreId = value[3]
	m.OrganizeId = value[4]
	m.FinanceName = value[5]
	m.FinanceId = value[6]
	m.Tag = value[7]
	m.ManageOrganize = value[8]

	m.Country = value[9]
	m.Province = value[10]
	m.City = value[11]
	m.Area = value[12]
	m.Street = value[13]
	m.Address = value[14]

	m.Phone = value[15]
	m.BBS = value[16]

	m.BusinessHours = value[17]
	m.Status = value[18]
	m.BusinessMode = value[19]
	m.SystemType = value[20]

	// 地理位置
	la := strings.Trim(value[21], `"`)
	lg := strings.Trim(value[22], `"`)
	m.LatitudeLongitude = la + "," + lg
	// 创建人
	m.CreatedBy = value[23]

	// 更新时间 value[24]
	m.UpdatedBy = value[25]

	// 分组名
	m.StoreGroupName = value[len(value)-1]
	DataHandler.Create(&m)
}

// All 获取所有数据
func (m Stores) All(instance interface{}) error {
	err := m.all("stores", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m Stores) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("stores", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("stores", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m Stores) GetOne(instance interface{}) error {
	err := m.getOne("stores", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByFilterOffset 获取所有数据
// 以便管理员进行审核操作
func (m Stores) ListByFilterOffset(instance interface{}, filter []string, filterParams []string, offset int, limit int) (int64, error) {
	var counter int64 = 0
	err := m.counterByFilter("stores", &counter, filter, filterParams)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}
	// 获取列表
	err = m.listByFilterOffset("stores", instance, offset, limit, filter, filterParams)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// Delete 删除记录
func (m Stores) Delete() error {
	err := m.delete("stores")
	if err != nil {
		return err
	}
	return nil
}

// Save 删除记录
func (m Stores) Save() error {
	m.save(&m)
	return nil
}
