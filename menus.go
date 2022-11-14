package model

// 根据huanglj样本设计

// MenuBase 菜单基本信息
type MenuBase struct {
	// MenuId 菜品编码
	MenuId string `json:"menu_id"`

	// POS 分类
	PosCategory string `json:"pos_category"`
	// 线上 分类
	OnlineCategory string `json:"online_category"`

	// 收入科目
	IncomeAccount string `json:"income_account"`
	// 出品部门
	ProductionDepartment string `json:"production_department"`
}

// MenuSpecification 规格
type MenuSpecification struct {
	// 规格名称
	SpecName string `json:"spec_name"`
	// 常规售价
	CommonPrice float64 `json:"common_price"`
	// 常规会员价
	CommonVipPrice float64 `json:"common_vip_price"`
	// 外卖价
	TakeoutPrice float64 `json:"takeout_price"`
	// 外卖会员价
	TakeoutVipPrice float64 `json:"takeout_vip_price"`
	// 堂食价
	DinePrice float64 `json:"dine_price"`
	// 堂食会员价
	DineVipPrice float64 `json:"dine_vip_price"`

	// 自提价
	PickUpPrice float64 `json:"pick_up_price"`
	// 堂食价
	PickUpVipPrice float64 `json:"pick_up_vip_price"`
	// 原价
	OriginalPrice float64 `json:"original_price"`
	// 成本
	Cost float64 `json:"cost"`
}

type MenuSupport struct {
	// 支持业务（列表)
	SupportBusiness string `json:"support_business"`
	// 支持口味（列表)
	SupportTaste string `json:"support_taste"`
	// 支持做法（列表)
	SupportPractice string `json:"support_practice"`
}

type MenuEnable struct {
	// 菜品启用
	IsEnable bool `json:"is_enable"`
	// 菜品上架
	IsShelf bool `json:"is_shelf"`
	// 是否招牌
	IsFascia bool `json:"is_fascia"`
	// 是否新菜
	IsNew bool `json:"is_new"`
	// 是否推荐
	IsRecommend bool `json:"is_recommend"`
	// 是否打折
	IsDiscount bool `json:"is_discount"`
	// 是否推荐
	IsSupportDecimals bool `json:"is_support_decimals"`
	// 是否打
	IsWeigh bool `json:"is_weigh"`
	// 是否开启 开台自动加入
	IsAutoJoin bool `json:"is_auto_join"`
	// 是否支持单独销售
	IsSaleAlone bool `json:"is_sale_alone"`
}

// MenuSales 售卖属性
type MenuSales struct {
	MenuSupport

	MenuEnable

	// 最小启售数量
	SaleMiniNum uint `json:"sale_mini_num"`
	// 销售提点
	SaleCommission float64 `json:"sale_commission"`
	// 打包费用
	Pack float64 `json:"pack"`
}

// Menus 菜单
type Menus struct {
	BaseModel
	// 菜品基本信息
	MenuBase
	// 菜品销售信息
	MenuSales
	//// 存储序列化后的数据
	//Baskets string `json:"-"`
	//// 对外提供列表
	Basket []MenuSpecification `json:"basket" gorm:"-"`
	// CreatedAt 创建人
	CreatedBy string `json:"created_by"`
	// UpdatedAt 修改人
	UpdatedBy string `json:"updated_by"`
	// Name 菜品名称
	Name string `json:"name"`
}

//
//func (m Menus) MarshalJSON() ([]byte, error) {
//	// 命名别名，避免MarshalJson死循环
//	type AliasMenus Menus
//	err := json.Unmarshal([]byte(m.Baskets), &m.Basket)
//	if err != nil {
//		return nil, err
//	}
//
//	return json.Marshal(struct {
//		AliasMenus
//	}{AliasMenus(m)})
//}

// All 获取所有数据
func (m Menus) All(instance interface{}) error {
	err := m.all("menus", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m *Menus) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("menus", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("menus", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m *Menus) GetOne(instance interface{}) error {
	err := m.getOne("menus", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByFilterOffset 获取所有数据
// 以便管理员进行审核操作
func (m *Menus) ListByFilterOffset(instance interface{}, filter []string, filterParams []string, offset int, limit int) (int64, error) {
	var counter int64 = 0
	err := m.counterByFilter("menus", &counter, filter, filterParams)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}
	// 获取列表
	err = m.listByFilterOffset("menus", instance, offset, limit, filter, filterParams)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// Delete 删除记录
func (m *Menus) Delete() error {
	err := m.delete("menus")
	if err != nil {
		return err
	}
	return nil
}

// Save 保存记录
func (m *Menus) Save() error {
	err := m.save(m)
	if err != nil {
		return err
	}
	return nil
}
