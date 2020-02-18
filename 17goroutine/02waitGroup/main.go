package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		a := rand.Intn(10)
		fmt.Println(a)
	}
}

func f1(i int) {
	defer wait.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

var wait sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go f1(i)
	}
	wait.Wait()
	fmt.Println("???执行完了")
}
