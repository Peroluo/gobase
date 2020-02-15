package main

import "fmt"

type address struct {
	province string
	city     string
}
type workPlace struct {
	province string
	city     string
}
type person struct {
	name string
	address
	workPlace
}

type company struct {
	addr address
}

func main() {
	p1 := person{
		name: "liwen",
		address: address{
			province: "深圳",
			city:     "南山",
		},
		workPlace: workPlace{
			province: "深圳",
			city:     "南山",
		},
	}
	fmt.Println(p1.address.city)
	// fmt.Println(p1.city) // 先在自己结构体找到这个字段，找不到就去匿名结构体里面找,当有多个匿名字段，并还有该属性,不能直接访问city属性
	fmt.Println(p1.workPlace.city)
}
