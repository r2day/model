package model

// Invoices 发票
type Invoices struct {
	BaseModel

	BaseFee

	// CustomerId 客户编号(id)
	CustomerId uint `json:"customer_id"`

	// 订单id
	CommandId uint `json:"command_id"`

	// status 发票状态
	Status string `json:"status"`
}

// All 获取所有数据
func (m Invoices) All(instance interface{}) error {
	err := m.all("invoices", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByOffset 获取所有数据
// 以便管理员进行审核操作
func (m Invoices) ListByOffset(instance interface{}, offset int, limit int) (int64, error) {
	var counter int64 = 0

	err := m.counter("invoices", &counter)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}

	// 获取列表
	err = m.listByOffset("invoices", instance, offset, limit)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// GetOne 获取单个数据
// 以便管理员进行审核操作
func (m Invoices) GetOne(instance interface{}) error {
	err := m.getOne("invoices", instance)
	if err != nil {
		return err
	}
	return nil
}

// ListByFilterOffset 获取所有数据
// 以便管理员进行审核操作
func (m Invoices) ListByFilterOffset(instance interface{}, filter []string, offset int, limit int) (int64, error) {
	var counter int64 = 0

	filterParams := make([]string, 0)
	filterParams = append(filterParams, m.Status)

	err := m.counterByFilter("invoices", &counter, filter, filterParams)
	if err != nil {
		return 0, err
	} else if counter == 0 {
		return 0, nil
	}
	// 获取列表
	err = m.listByFilterOffset("invoices", instance, offset, limit, filter, filterParams)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

// Delete 删除记录
func (m Invoices) Delete() error {
	err := m.delete("invoices")
	if err != nil {
		return err
	}
	return nil
}
