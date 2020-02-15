package main

import "fmt"
// 整型 int8、int16、int32、int64
// 无符号整型 uint8、uint16、uint32、uint64
// 特殊整型 uint;32位系统就uint32、64位系统就是uint64
// 特殊整型 int;32位系统就int32、64位系统就是int64
// 特殊整型 uintptr 无符号整型，用于存放指针

// 8进制 16进制
func main()  {
	a:=101
	fmt.Printf("%d\n",a)
	fmt.Printf("%b\n",a)// 二进制
	fmt.Printf("%o\n",a)// 八进制
	fmt.Printf("%x\n",a)// 十六进制
	fmt.Printf("%T\n",a)// 变量类型
	// 八进制
	i2:=077
	fmt.Printf("%d\n",i2)
	i3:=0x123
	fmt.Printf("%d\n",i3)
	// 声明int8类型
	i4:=int8(10)
	fmt.Printf("%T\n",i4)
}