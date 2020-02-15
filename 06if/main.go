package main

import "fmt"

func main() {
	// 普通写法
	age := 18
	if age > 18 {
		fmt.Println("成年了")
	} else if age < 18 {
		fmt.Println("18岁")
	} else {
		fmt.Println("88")
	}
	// 特殊写法(sex用作用域，在if内可用)
	if sex := 1; sex == 1 {
		fmt.Println("你是男的")
	} else {
		fmt.Println("你是女的")
	}
}
