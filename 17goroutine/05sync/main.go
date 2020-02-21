package main

import (
	"fmt"
	"sync"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func run() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go run()
	go run()
	wg.Wait()
	fmt.Println(x)
}
