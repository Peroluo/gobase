package main

import "fmt"

func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 变种1
	i := 2
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	a := 1
	// 变种2
	for a < 2 {
		fmt.Println(a)
		a++
	}
	// 无限循环
	// for{
	// 	fmt.Println("???")
	// }

	// for range
	s := "luoguoxiong罗国雄"
	for index, val := range s {
		fmt.Printf("%d,%c\n", index, val)
		// 0,l
		// 1,u
		// 2,o
		// 3,g
		// 4,u
		// 5,o
		// 6,x
		// 7,i
		// 8,o
		// 9,n
		// 10,g
		// 11,罗
		// 14,国
		// 17,雄
	}
}
