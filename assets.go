package model

// Assets 资产信息
type Assets struct {
	Balance float64 `json:"balance"`
	// 现金卡值
	CashCharge float64 `json:"cash_charge"`
	// 冻结卡值
	Freezing float64 `json:"freezing"`
	// 赠送卡值
	Gift float64 `json:"gift"`
	// 积分余额
	Integral uint64 `json:"integral"`
	// 累计储值总额
	TotalBalance float64 `json:"total_balance"`
	//累计储值次数
	TotalBalanceCounter uint64 `json:"total_balance_counter"`
	// 累计消费总额
	TotalCumulativeConsumption float64 `json:"total_cumulative_consumption"`
	// 累计消费总额
	TotalCumulativeConsumptionCounter uint64 `json:"total_cumulative_consumption_counter"`
	// 挂帐总额度
	DebitTotalLimit float64 `json:"debit_total_limit"`
	// 挂帐剩余额度
	DebitLeftLimit float64 `json:"debit_left_limit"`

	// 已用额度
	DebitUsedLimit float64 `json:"debit_used_limit"`
}
