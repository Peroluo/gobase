package main

import "fmt"

// 变声明
var (
	name string // ''
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "guoxiong"
	fmt.Println(name)
	age = 27
	fmt.Println(age)
	isOk = false
	fmt.Println(isOk)
	// 类型推倒 声明并赋值
	var a = "55"
	fmt.Println(a)
	// 简短声明 只能在函数内部
	b := true
	fmt.Println(b)
	// 匿名变量声明 _
	fmt.Println("helloworld")
}
