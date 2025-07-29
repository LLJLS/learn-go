package main

import (
	"fmt"
	"math"
)

/*
1.题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	long float64
	wide float64
}

func (sh Rectangle) Area() {
	fmt.Println("长方形的面积：", sh.long*sh.wide)
}

func (sh Rectangle) Perimeter() {
	fmt.Println("长方形的周长：", (sh.long+sh.wide)*2)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() {
	fmt.Printf("圆形的面积：%.2f\n", math.Pi*math.Pow(c.radius, 2))
}

func (c Circle) Perimeter() {
	fmt.Printf("圆形的周长：%.2f\n", 2*math.Pi*c.radius)
}

/*
2.题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工id:%d,姓名：%s,年龄：%d\n", e.EmployeeID, e.Person.Name, e.Person.Age)
}

func main() {
	// Q1
	r := Rectangle{long: 20, wide: 10}
	c := Circle{radius: 5}
	f := func(s Shape) {
		s.Area()
		s.Perimeter()
	}
	f(r)
	f(c)
	// Q2
	p := Person{Name: "张三", Age: 18}
	e := Employee{Person: p, EmployeeID: 1}
	e.PrintInfo()

}
