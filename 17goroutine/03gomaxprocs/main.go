package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("a:%d \n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("b:%d \n", i)
	}
}

func main() {
	runtime.GOMAXPROCS(8) //默认就是cpu的核心数
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
