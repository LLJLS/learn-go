package main

import (
	"fmt"

	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

type Students struct {
	Id    uint `gorm:"primarykey"`
	Name  string
	Age   uint
	Grade string
}

func printRes(db *gorm.DB, res *[]Students) {
	if db.Error != nil {
		log.Fatal("query failed")
	}
	for _, v := range *res {
		fmt.Println(v)
	}
}

func main() {
	dsn := "remote:mysql_remote@tcp(192.168.40.10:3306)/web3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database connection failed")
		panic(err)
	}
	// 建表
	db.Debug().AutoMigrate(&Students{})

	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	stu := Students{Name: "张三", Age: 20, Grade: "三年级"}
	db.Debug().Create(&stu)

	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	stus := []Students{}
	res := db.Debug().Where("age > ?", 18).Find(&stus)
	printRes(res, &stus)
	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	stu1 := Students{Id: 2}
	res = db.Debug().Model(&stu1).Updates(Students{Grade: "三年级"})
	printRes(res, &[]Students{stu1})
	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	db.Debug().Where("age < ?", 15).Delete(&Students{})

}
