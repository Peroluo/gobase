package main

import (
	"encoding/json"
	"fmt"
)

// json转换要使用大写开头
type person struct {
	Name string `json:"name",db:"name"` //不同的转换对应不同的名称
	Age  int    `json:"age",db:"age"`
}

func main() {
	p1 := person{
		Name: "sfs",
		Age:  12,
	}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", string(b))
	}
	// 反序列化
	str := `{"name":"李翔","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Println(p2)
}
