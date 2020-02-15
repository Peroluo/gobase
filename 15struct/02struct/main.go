package main

import "fmt"

type person struct {
	name  string
	age   int
	hobby []string
}

func main() {
	var p person
	p.name = "lixiang"
	p.age = 90
	p.hobby = []string{"篮球"}
	fmt.Println(p)
	fmt.Printf("%T \n", p)
}
