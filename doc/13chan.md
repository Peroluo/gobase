# channel 通道

#### 1. chan声明

```go
var ch1 chan int //声明ch1是int类型的通道
```

##### 2. chan初始化

```go
var ch1 chan int //声明ch1是int类型的通道
fmt.Println(ch1==nill) // true
ch1 = make(chan int, 1) // ch1初始化 ,1:为通道的缓存区数量
```

#### 3.chan <-

```go
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
	wg.Wait()
	for c := range ch1 {
		fmt.Println(c)
	}
}
```

#### 3. 单向通道

> 单向通道一般用于当做函数参数进行传值
>
> ch1 := make(chan int, 10)  双向通道
>
> ch2 := make(chan<- int, 10)  只接收
>
> ch3 := make(<-chan int, 10) 只发送

```go
package main

import (
	"fmt"
	"sync"
)

func f1(ch chan<- int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		v, ok := <-ch1
		if !ok {
			fmt.Println("还没有接收到ch1通道的值")
			break
		}
		ch2 <- v * v
	}
	close(ch2)
}

var wg sync.WaitGroup

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	wg.Add(2)
	go f1(ch1)
	go f2(ch1, ch2)
	wg.Wait()
	for c := range ch2 {
		fmt.Println(c)
	}
}

```

#### 4. close

> 使用for range获取通道值，记得close

```go
ch1 := make(chan int, 2)
ch1 <- 1
ch1 <- 3
// close(ch1) // fatal error: all goroutines are asleep - deadlock!
for c := range ch1 {
    fmt.Println(c)
}
```

