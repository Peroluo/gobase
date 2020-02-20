# goroutine 并发

#### 1. goroutine 

> go 关键字

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second * 2) // 避免goroutine结束前，程序停止
}

```

#### 2. sync.WaitGroup

> 上述demo中，我们不能知道goroutine什么时候执行完，可借助sync.WaitGroup的Add、Done、Wait方法。	

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Microsecond * 100)
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("1000个goroutine已执行完！")
}
```

#### 3. workerpoll 线程池

```go
package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 输出结果
	for a := 1; a <= 5; a++ {
		// 如果没有接受到值，就会造成main一直等待，直到完成接收到值
		fmt.Println(<-results)
	}
}

```

