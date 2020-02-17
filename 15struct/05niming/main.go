package main

import "fmt"

type person struct {
	name string
	age  int
	bool // 匿名字段类型只有一个
}

func main() {
	p1 := person{
		"周玲",
		2323,
		false,
	}
	// 访问匿名字段
	fmt.Println(p1.bool)
}
