package main

import "fmt"

// 自定义类型
type myInt int

type youInt = int

// 给自定义类型添加方法，不能给别的包里内添加方法(只在当前包)
func (m myInt) hello() {
	fmt.Println("HELLO")
}
func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n) // main.myInt

	n1 := myInt(11)
	n1.hello()

	var y youInt
	y = 100
	fmt.Printf("%T\n", y) //int

	var c rune
	c = '中'
	fmt.Printf("%T\n", c) //int32
}
