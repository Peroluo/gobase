package main

import "fmt"

// interface是一种数据类型
// 一个结构体可以实现多个接口
type cat struct{}

type dog struct{}

// 接口定义
// type 接口名 interface{
// 方法名1(参数)(返回值)
// 方法名2(参数)(返回值)
// }
type spreaker interface {
	spreak()
}

func (c cat) spreak() {
	fmt.Println("喵喵")
}

func (d dog) spreak() {
	fmt.Println("旺旺")
}

func spreak(x spreaker) {
	// 传什么参数，什么参数调用spreak
	x.spreak()
}

// 接口实现
// 一个变量实现了接口规定的“所有方法”,那么这个变量就实现了这个接口。
type animal interface {
	move()
	eat(string)
}

type fish struct {
	name string
}

func (f fish) eat(foot string) {
	fmt.Printf("%s吃%s\n", f.name, foot)
}

func (f fish) move() {
	fmt.Printf("%ss开始跑动\n", f.name)
}

type person struct {
	name string
}

func (f *person) eat(foot string) {
	fmt.Printf("%s吃%s\n", f.name, foot)
}
func (f *person) move() {
	fmt.Printf("%s开始跑动\n", f.name)
}

func main() {
	// 1.接口定义
	var c cat
	var d dog
	spreak(c)
	spreak(d)
	// 2.接口实现
	var a1 animal //定义一个类型接口变量
	fish := fish{
		name: "鱼",
	}
	a1 = fish //接口实现
	a1.eat("草")
	// person的eat方法是指针接受者，所以必须使用“&person”实现接口
	a1 = &person{
		name: "xiaoluo",
	} //接口实现
	a1.eat("肉")
}
