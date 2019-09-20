package main

import (
	"fmt"
)

// Bank 银行
type Bank struct {
	saving map[string]int // 每账户的存款金额
}

// Request 银行存取操作
type Request struct {
	op    string       // 存、取、查
	name  string       // 操作的账号
	value int          // 操作金额
	retCh chan *Result // 存放银行处理结果的通道
}

// Result 执行结果
type Result struct {
	success bool // 成功
	value   int  // 查询时使用：余额
}

// NewBank 工厂函数
func NewBank() *Bank {
	b := &Bank{
		saving: make(map[string]int),
	}
	return b
}

// Loop 银行处理客户请求
func (b *Bank) Loop(reqCh chan *Request) {
	for req := range reqCh {
		switch req.op {
		case "deposite":
			b.Deposit(req)
		case "withdraw":
			b.Withdraw(req)
		case "query":
			b.Query(req)
		default:
			// 响应
			ret := &Result{
				false,
				0,
			}
			req.retCh <- ret
		}
	}

	// 无请求时银行退出
	fmt.Println("Bank exit")
}

// Deposit 存款
func (b *Bank) Deposit(req *Request) {
	name := req.name
	amount := req.value

	if _, ok := b.saving[name]; !ok {
		b.saving[name] = 0
	}
	b.saving[name] += amount

	// 响应
	ret := &Result{
		true,
		0,
	}
	req.retCh <- ret
}

// Withdraw 取款，不足时取款失败
func (b *Bank) Withdraw(req *Request) {
	name := req.name
	amount := req.value

	var status bool
	if balance, ok := b.saving[name]; !ok || balance < amount {
		status = false
		amount = 0
	} else {
		status = true
		b.saving[name] -= amount
	}

	// 响应
	ret := &Result{
		status,
		amount,
	}
	req.retCh <- ret
}

// Query 查询余额
func (b *Bank) Query(req *Request) {
	name := req.name

	var (
		ok      bool
		balance int
	)

	if balance, ok = b.saving[name]; !ok {
		balance = 0
	}

	// 响应
	ret := &Result{
		true,
		balance,
	}
	req.retCh <- ret
}
