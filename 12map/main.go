package main

import "fmt"

// map

func main() {
	// 初始化
	m1 := make(map[string]int, 10)
	m1["李翔"] = 100
	fmt.Println(m1)
	// 判断是不是有值
	v, isok := m1["xx"]
	if !isok {
		fmt.Println("没有key值")
	} else {
		fmt.Println(v)
	}
	// for range遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 删除
	delete(m1, "李翔")
	fmt.Println(m1)

	// make和slice组合
	var s1 = make([]map[string]string, 10, 10)
	s1[0] = make(map[string]string, 1)
	s1[0]["te"] = "other"
	fmt.Println(s1)

	// 值为切片类型的map
	var m2 = make(map[string][]int, 10)
	m2["背景"] = []int{10, 20, 30}
	fmt.Println(m2)
}
