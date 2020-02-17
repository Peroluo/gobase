package main

import "fmt"

type person struct {
	name string
	age  int
}

func f(x person) {
	x.name = "23"
}

func f2(x *person) {
	// (*x).name = "23" //跟下面是一样的
	x.name = "23" //根据内存地址找到原变量 语法糖
}
func main() {
	var p person
	p.age = 8
	p.name = "te"
	f(p)
	// f没有改到值
	fmt.Println(p)          //{te 8}
	fmt.Printf("%T\n", p)   //person
	fmt.Printf("%p \n", &p) //0xc0000044a0
	f2(&p)
	fmt.Println(p)

	// 指针
	var p2 = new(person)
	fmt.Printf("%T \n", p2)  //person
	fmt.Printf("%p \n", p2)  //0xc000004520 p2的值就是一个内存地址
	fmt.Printf("%p \n", &p2) //p2的内存地址 0xc000004540
	fmt.Println(p2)          //零值 &{ 0}

	var p3 = &person{name: "SD", age: 19}
	fmt.Printf("%T \n", p3)  //person
	fmt.Println(p3)          //&{SD 19}
	fmt.Printf("%p \n", &p3) //0xc000004580

	var p4 = &person{"sd", 12}
	fmt.Println(p4)          //&{sd 12}
	fmt.Printf("%p \n", &p4) //0xc0000045c0
}
