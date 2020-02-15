package main

import "fmt"

func main() {

	// 切片指向一个数组
	// 切片的长度是它元素的个数
	// 切片的容量是从切片的第一个元素,到底层数组最后一个元素的长度

	// 声明切片
	var s1 []int           //定义一个int类型元素的切片
	fmt.Println(s1 == nil) //[]===nill true
	// 初始化
	s1 = []int{1, 2, 3}
	fmt.Println(s1)
	// 切片长度和容量
	fmt.Printf("len:%d cap:%d\n", len(s1), cap(s1))

	// 由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5}
	s2 := a1[0:4] // 基于一个数组 左闭右开；不包括4
	fmt.Println(s2)
	s3 := a1[1:]
	fmt.Println(s3)
	fmt.Printf("len:%d cap:%d\n", len(s3), cap(s3))
	s4 := a1[:4]
	fmt.Println(s4)
	s5 := a1[:]
	fmt.Println(s5)

	// make(类型,长度,容量)函数构造切片
	s6 := make([]int, 5, 10)
	fmt.Printf("s6=%v len(s6)=%d cap(s6)=%d", s6, len(s6), cap(s6)) //s6=[0 0 0 0 0] len(s6)=5 cap(s6)=10

	// for range
	for i, val := range s6 {
		fmt.Println(i, val)
	}
	// append(slice,slice...)
	s7 := []int{1, 2, 3}
	fmt.Println(s7)
	s8 := append(s7, 5, 6, 7)
	fmt.Println(s8)
	// 这样些才算对的
	s8 = append(s8, 8)
	fmt.Println(s8)

	// copy(new,old) //产生了新的底层数组
	s9 := []int{1, 2, 3}
	s10 := make([]int, 3, 3)
	copy(s10, s9)
	fmt.Println(s10, s9) //[1 2 3] [1 2 3]
	s9[0] = 4
	fmt.Println(s10, s9) //[1 2 3] [4 2 3]
	// 没有删除元素的方法
}
