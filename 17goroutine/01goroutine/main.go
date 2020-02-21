package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int, 8)
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ { // 超过缓冲区8会报错
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			ch1 <- i
		}(i)
	}
	wg.Wait() // 等待所有go func执行完~
	close(ch1)
	for c := range ch1 {
		fmt.Println(c)
	}
}
