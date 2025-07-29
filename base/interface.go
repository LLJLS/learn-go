package main

import "fmt"

// PaymentMethod 接口定义了支付方法的基本操作
type PaymentMethod interface {
	Account
	Pay(amount int) bool
}

type Account interface {
	GetBalance() int
}

// CreditCard 结构体实现 PaymentMethod 接口
type CreditCard struct {
	// 已用额度
	balance int
	// 限额
	limit int
}

func (c *CreditCard) Pay(amount int) bool {
	if c.balance+amount <= c.limit {
		c.balance += amount
		fmt.Printf("信用卡支付成功：%d\n", amount)
		return true
	}
	fmt.Println("信用卡支付失败：超出额度")
	return false
}

func (c *CreditCard) GetBalance() int {
	return c.balance
}

// DebitCard 结构体实现 PaymentMethod 接口
type DebitCard struct {
	// 余额
	balance int
}

func (d *DebitCard) Pay(amount int) bool {
	if d.balance >= amount {
		d.balance -= amount
		fmt.Printf("借记卡支付成功：%d\n", amount)
		return true
	}
	fmt.Println("借记卡支付失败：余额不足")
	return false
}

func (d *DebitCard) GetBalance() int {
	return d.balance
}

// 使用PaymentMethod 接口的函数
func purchaseItem(p PaymentMethod, price int) {
	if p.Pay(price) {
		fmt.Printf("购买成功，剩余余额：%d\n", p.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}

type Test struct {
	id   int
	name string
}

// 只要实现了接口的结构体都可以转换
// 接口可以接受任何类型
func main() {
	creditCard := &CreditCard{balance: 0, limit: 1000}
	debitCard := &DebitCard{balance: 500}

	fmt.Println("使用信用卡购买：")
	purchaseItem(creditCard, 800)

	fmt.Println("使用借记卡购买：")
	purchaseItem(debitCard, 300)

	fmt.Println("再次使用借记卡购买：")
	purchaseItem(debitCard, 200)

	// var accountA Account = creditCard
	// fmt.Println("获取账户A的信息", accountA)

	var accountA Account = debitCard
	fmt.Println("获取账户A的余额", accountA.GetBalance())

}
