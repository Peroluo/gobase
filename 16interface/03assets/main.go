package main

import "fmt"

func main() {
	var a interface{} = 100
	fmt.Printf("%T\n", a)
	// 类型断言
	v1, ok := a.(int8)
	if ok {
		fmt.Println("int8", v1)
	} else {
		fmt.Println("不是int8")
	}

	// switch
	switch v2 := a.(type) {
	case int:
		fmt.Println("int", v2)
	default:
		fmt.Println("不明类型数据")
	}
}
