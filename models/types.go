package models

import (
	"errors" // 用来报错

	"github.com/shopspring/decimal"
)

// Account: 账户结构 (升级版)
type Account struct {
	ID   int
	Name string
	// 核心升级：使用 Map 存储多币种余额
	// Key = 币种名称 ("USDT"), Value = 数量
	Balances map[string]decimal.Decimal
}

// Order: 订单结构 (不变)
type Order struct {
	ID       int
	Symbol   string
	Side     string
	Price    decimal.Decimal
	Quantity decimal.Decimal
	User     *Account
}

// --- 下面是新加的“资金操作”方法 ---

// Deposit: 充值
func (a *Account) Deposit(symbol string, amount decimal.Decimal) {
	if a.Balances == nil {
		a.Balances = make(map[string]decimal.Decimal)
	}

	current := a.Balances[symbol]
	a.Balances[symbol] = current.Add(amount)
}

// Withdraw: 提现 (带余额检查)
func (a *Account) Withdraw(symbol string, amount decimal.Decimal) error {
	if a.Balances == nil {
		return errors.New("wallet is empty")
	}

	current := a.Balances[symbol]
	if current.LessThan(amount) {
		return errors.New("insufficient balance")
	}

	a.Balances[symbol] = current.Sub(amount)
	return nil
}
