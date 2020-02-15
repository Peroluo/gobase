package main

import "fmt"

// defer作用域不会影响其他
// 多用于函数结束之前释放资源(文件句并、数据库连接、socket连接)
func deferDemo() {
	fmt.Println("heihei1")
	defer //defer把后面的语句延迟到函数即将返回的时候再执行
	fmt.Println("heihei2")
	defer fmt.Println("XX")
	fmt.Println("heihei3")
	// 多个defer后进先出
	// heihei1
	// heihei3
	// XX
	// heihei2
}

// Go的return不是原子操作，在底层是分两步执行
// 第一步："返回值"赋值
// defer
// 第二步：返回
// 函数如果存在defer，那么defer的执行实际在第一步和第二步直接
func main() {
	deferDemo()
	fmt.Println("??")
}
