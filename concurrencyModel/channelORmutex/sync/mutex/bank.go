package main

import "sync"

// Bank 银行结构体
type Bank struct {
	sync.Mutex
	saving map[string]int  // 存款金额
}

// NewBank 生成银行对象
func NewBank() *Bank{
	b := &Bank{
		saving : make(map[string]int),
	}
	return b
}

// Deposit 存款
func (b *Bank) Deposit(name string, amount int){
	b.Lock()
	defer b.Unlock()

	if _, ok := b.saving[name]; !ok{
		b.saving[name] = 0
	}
	b.saving[name] += amount
}

// Withdraw 取款，返回实际取到的金额
func (b *Bank) Withdraw(name string, amount int) int{
	b.Lock()
	defer b.Unlock()

	if _, ok := b.saving[name];!ok{
		return 0
	}

	if b.saving[name] < amount{
		amount = b.saving[name]
	}

	b.saving[name] -= amount
	return amount
}

// Query 查询余额
func (b *Bank) Query(name string) int {
	b.Lock()
	defer b.Unlock()

	if _, ok := b.saving[name]; !ok {
		return 0
	}

	return b.saving[name]
}