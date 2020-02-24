package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var notify bool

func f(ctx context.Context) {
FORLOOP:
	for {
		fmt.Println("SFS")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
	defer wg.Done()
}

func main() {
	ctx, cancel := context.WithCa ncel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 10)
	cancel()
	wg.Wait()
}
