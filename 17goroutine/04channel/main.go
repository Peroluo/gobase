package main

import (
	"fmt"
	"sync"
)

func noBufChannel() {
	b := make(chan int)
	go func() {
		x := <-b
		fmt.Println("b渠道了", x)
	}()
	b <- 10
	fmt.Println("10发送到b中了")
}

func bufChannel() {
	b := make(chan int, 1) // 容量 里面只能存一个chan
	b <- 10
	fmt.Println("10发送到b中了")
	x := <-b
	fmt.Println("b渠道了", x)
	b <- 20
	fmt.Println("20发送到b中了")
	c := <-b
	fmt.Println("b渠道了", c)
	close(b)
}

// 1.驱动一个goroutune,生成100个数发送ch1
// 2.启动一个goroutine,从ch1中取值，其平方放到ch2
// 在main中 ch2打印

var wg sync.WaitGroup

func te1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func te2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}

func main() {
	// noBufChannel()
	// bufChannel()
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(2)
	go te1(a)
	go te2(a, b)
	for ret := range b {
		fmt.Println(ret)
	}
}
