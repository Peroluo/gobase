package main

import "fmt"

func main() {
	switch a := 12; a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("DEFAULT")
	}
	// 多个类型
	switch n := 2; n {
	case 1, 3, 5, 6, 7, 9:
		fmt.Println("奇数")
	default:
		fmt.Println("偶数")
	}
	// fallthrough //匹配到，再执行下一项
	switch s := "a"; s {
	case "a":
		fmt.Println("a")
		fallthrough
	case "a+":
		fmt.Println("a+")
	}
}
