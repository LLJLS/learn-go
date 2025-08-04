package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

type Accounts struct {
	Id      uint `gorm:"primaryKey"`
	Balance float64
}

type Transactions struct {
	Id              uint `gorm:"primaryKey"`
	From_account_id uint
	From            Accounts `gorm:"foreignKey:from_account_id"`
	To_account_id   uint
	To              Accounts `gorm:"foreignKey:to_account_id"`
	Amount          float64
}

func transfer(db *gorm.DB, from uint, to uint, amount float64) bool {
	stuF := Accounts{}
	db.Where("id=?", from).First(&stuF)
	newBalanceF := stuF.Balance - amount
	if newBalanceF < 0 {
		fmt.Printf("账户[%d]:余额为%.2f,余额不足\n", stuF.Id, stuF.Balance)
		return false
	}

	stuT := Accounts{}
	db.Where("id=?", to).First(&stuT)
	tx := db.Begin()
	// update From account
	stuFA := Accounts{Id: from}
	res := db.Debug().Model(&stuFA).Updates(map[string]interface{}{"Balance": newBalanceF})
	if res.Error != nil {
		tx.Rollback()
		panic(res.Error)
		return false
	}
	// update To account
	stuTA := Accounts{Id: to}
	res = db.Debug().Model(&stuTA).Updates(map[string]interface{}{"Balance": stuT.Balance + amount})
	if res.Error != nil {
		tx.Rollback()
		panic(res.Error)
		return false
	}
	// update transaction
	tr := Transactions{From_account_id: from, To_account_id: to, Amount: amount}
	res = db.Debug().Create(tr)
	if res.Error != nil {
		tx.Rollback()
		panic(res.Error)
		return false
	}
	tx.Commit()
	return true
}

func main() {
	dsn := "remote:mysql_remote@tcp(192.168.40.10:3306)/web3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database connection failed")
		panic(err)
	}

	// create table
	// db.AutoMigrate(&Accounts{}, &Transactions{})

	// create simples
	// db.Create(&Accounts{Balance: 100})
	// db.Create(&Accounts{Balance: 50})

	// transfer
	isSucc := transfer(db, 1, 2, 100)
	if isSucc {
		fmt.Println("转账成功")
	} else {
		fmt.Println("转账失败")
	}
}
