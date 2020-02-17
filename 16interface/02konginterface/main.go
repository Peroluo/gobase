package main

import "fmt"

//interface{}//空接口，没有必要起名字。通常定义成这种格式

// 所有的类型都实现了空接口，也就是任意类型的变量都能保存到空接口中

func show(a interface{}) {
	// 类型断言
	v, ok := a.(string)
	if !ok {
		fmt.Println("不是string")
	} else {
		fmt.Println(v)
	}
	// switch a.(type)
	switch va := a.(type) {
	case string:
		fmt.Println("is string", va)
	case int:
		fmt.Println("is int", va)
	case bool:
		fmt.Println("is bool", va)
	}
	fmt.Printf("%T\n  ", a)
}
func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "zhouling"
	m1["age"] = 18
	fmt.Println(m1)
	show(m1["name"])
	show(m1["age"])
}
