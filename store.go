package model

import (
	"strings"
	"time"
)

// StoreModel 店铺信息
type StoreModel struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// UserId 用户ID (登录dash平台的人)
	UserId string `json:"user_id"`
	// MerchantId 商户ID (例如: 黄李记作为一个商户存在)
	MerchantId string `json:"merchant_id" gorm:"index:idx_merchant"`
	// Status 状态
	Status string `gorm:"default:effected" gorm:"index:idx_status"`
	// StoreId store id
	StoreId string `json:"store_id"`
	// CreatedAt 创建时间
	CreatedAt time.Time
	// UpdatedAt 修改时间
	UpdatedAt time.Time

	Name     string `json:"name"`
	Position string `json:"position"`
	Addr     string `json:"addr"`
	Phone    string `json:"phone"`
	Pic      string `json:"pic"`
}

// Save 保存实例
func (m StoreModel) Save() {
	DataHandler.Create(&m)
}

func (m StoreModel) GetStore(merchantId string) []StoreModel {
	instance := make([]StoreModel, 0)
	DataHandler.Where("merchant_id = ?", merchantId).Find(&instance)
	return instance
}

// StoreInfo 门店信息
type StoreInfo struct {
	// Id 自增唯一id
	Id uint `json:"id" gorm:"unique"`
	// MerchantId 商户ID (例如: 黄李记作为一个商户存在)
	MerchantId string `json:"merchant_id" gorm:"index:idx_merchant"`
	// Status 状态
	Status string `gorm:"default:effected" gorm:"index:idx_status"`
	// CreatedAt 创建时间
	CreatedAt time.Time
	// UpdatedAt 修改时间
	UpdatedAt time.Time

	// CreatedAt 创建人
	CreatedBy string `json:"created_by"`
	// UpdatedAt 修改人
	UpdatedBy string `json:"updated_by"`

	GroupId   string `json:"group_id"`
	BrandName string `json:"brand_name"`
	StoreName string `json:"store_name"`
	// StoreId store id
	StoreId     string `json:"store_id" gorm:"index:idx_store_id,unique"`
	OrganizeId  string `json:"organize_id"`
	FinanceName string `json:"finance_name"`
	FinanceId   string `json:"finance_id"`
	// 标签
	Tag string `json:"tag"`

	ManageOrganize string `json:"manage_organize"`

	// 位置信息
	Country           string `json:"country"`
	Province          string `json:"province"`
	City              string `json:"city"`
	Area              string `json:"area"`
	Street            string `json:"street"`
	Address           string `json:"address"`
	LatitudeLongitude string `json:"latitude_longitude"`

	// 联系方式
	Phone string `json:"phone"`
	// 公告
	BBS string `json:"bbs"`
	// 营业时间
	BusinessHours string `json:"business_hours"`
	// 门店状态
	StoreStatus string `json:"store_status"`
	// 运营摸索
	BusinessMode string `json:"business_mode"`
	// 系统类型
	SystemType string `json:"system_type"`
	// 分组
	StoreGroupName string `json:"store_group_name"`
}

// SaveALine 保存实例
func (m StoreInfo) SaveALine(value []string) {

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
	m.StoreStatus = value[18]
	m.BusinessMode = value[19]
	m.SystemType = value[20]

	// 地理位置
	la := strings.Trim(value[21], `"`)
	lg := strings.Trim(value[22], `"`)
	m.LatitudeLongitude = la + "," + lg
	// 创建人
	m.CreatedBy = value[23]

	// 创建时间 value[24]
	m.UpdatedBy = value[25]

	// 更新时间 value[26]
	// 分组名
	m.StoreGroupName = value[27]
	DataHandler.Create(&m)
}

// ListAll 获取所有数据
// 店铺信息
func (m StoreInfo) ListAll() ([]StoreInfo, error) {
	instance := make([]StoreInfo, 0)
	err := DataHandler.Table("store_infos").
		Where("status = ? and merchant_id = ?", m.Status, m.MerchantId).
		Find(&instance).Error
	if err != nil {
		return nil, err
	} else {
		// 保存成功可以进行消息通知操作
		return instance, nil
	}
}
