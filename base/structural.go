package main

import (
	"fmt"
)

// 结构体声明
// 字段标记；-：忽略
// 多字段声明
// 匿名字段
type Student struct {
	Id, Score uint `json:"-" gorm:"column:<name>" json:"score" gorm:"column:<name>"`
	Name      string
	// 嵌套结构体
	Subject Subject
	bool
}

// 结构体函数
func (stu Student) setName(name string) {
	stu.Name = name
}

// 指针才能修改数据，值无法修改数据
// 结构体赋值给变量后，等于进行了深拷贝，变量原值不变
func (stu *Student) setPName(name string) {
	stu.Name = name
}

type Subject struct {
	Id   uint8
	Name string
}

func setStuName(stu Student, subName string) {
	stu.Name = subName
}

func setPStuName(stu *Student, subName string) {
	stu.Name = subName
}

func main() {
	fmt.Println("结构体")
	// 匿名结构体
	// 其中赋值后都必须加","
	// 匿名结构体主要用来组成测试数据
	// 匿名结构体可以空空
	// 空空匿名结构体常用于channel信息
	nounName := struct {
		id uint
	}{
		id: 100,
	}

	blank := struct{}{}

	noNameStructCh := make(chan struct{}, 1)

	fmt.Println(nounName.id)
	fmt.Println(blank)
	fmt.Println(noNameStructCh)

	subject := Subject{Id: 1, Name: "英语"}
	student := Student{Id: 001, Name: "张三", Subject: subject}
	fmt.Println(student.Name)
	student.setName("李四")
	fmt.Println(student.Name)
	student.setPName("李四")
	fmt.Println(student.Name)

	setStuName(student, "王五")
	fmt.Println(student.Name)
	setPStuName(&student, "王五")
	fmt.Println(student.Name)

}
