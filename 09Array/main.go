package main

import "fmt"

func main() {
	// 数组声明 值=默认元素的零值
	// 数组长度不可变
	var a1 [3]int
	fmt.Println(a1) //[0 0 0]
	var a2 [3]bool
	fmt.Println(a2) //[false false false]
	var a3 [3]string
	fmt.Println(a3) //['' '' '']
	// 数组初始化

	//初始化1
	a4 := [3]bool{true, true, true}
	fmt.Println(a4)

	// 初始化2,...根据初始值,推断数组长度
	a5 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a5)

	// 初始化3 根据索引初始化
	a6 := [5]int{0: 1, 4: 5}
	fmt.Println(a6)

	// 多维数组
	// [[1 2] [3 4] [5 6]]
	var a7 [3][2]int
	a7 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a7)
}
