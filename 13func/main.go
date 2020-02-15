package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

// 没有返回值
func f1(a int) {
	fmt.Println(a)
}

// 命名返回值，简写
func f2(a int) (ret int) {
	ret = a + 1
	return
}

// 多个返回值
func f3() (x int, y int) {
	x = 1
	y = 2
	return x, y
}

// 同类型简写
func f4(a, b, c int) int {
	return a + b + c
}

// 可变长参数,必须放在参数最后
func f5(y ...int) {
	fmt.Println(y) // y是slice
}

// 函数也是个类型
func teFn(x func(int) int, val int) func(int, int) int {
	c := x(val)
	return func(a int, b int) int {
		return a + b + c
	}
}

func main() {
	c := sum(1, 2)
	fmt.Println(c)
	_, y := f3()
	fmt.Println(y)
	f5(1, 2, 3, 4)
	// 函数类型
	fmt.Printf("%T\n", sum)
	// 函数类型参数和返回值
	fnnew := teFn(func(a int) int {
		return a + 4
	}, 4)(2, 4)
	fmt.Println(fnnew)
}
