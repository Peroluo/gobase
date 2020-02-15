package main

import "fmt"

func main()  {
	// &取地址
	// *去地址的值
	n:=18
	ponintN:=&n //*int
	fmt.Println(n,ponintN)
	// make和new都是申请内存的
	// new很少用，一般都是跟基本类型申请内存的，返回的是对应类型的指针
	// make给slice、map、chan申请内存。返回对应的三个类型的本身
	var a2=new(int)
	fmt.Println(a2)
	fmt.Println(*a2)
}