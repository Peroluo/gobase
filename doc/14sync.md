# sync

#### 1.sync.Mutex 互斥锁

```go
package main
import (
    "fmt"
    "sync"
    "time"
)
type Book struct {
    BookName string
    L        *sync.Mutex
}
func (bk *Book) SetName(wg *sync.WaitGroup, name string) {
    defer func() {
        fmt.Println("Unlock set name:", name)
        bk.L.Unlock()
        wg.Done()
    }()
    bk.L.Lock()
    fmt.Println("Lock set name:", name)
    time.Sleep(1 * time.Second)
    bk.BookName = name
}
func main() {
    bk := Book{}
    bk.L = new(sync.Mutex)
    wg := &sync.WaitGroup{}
    books := []string{"《三国演义》", "《道德经》", "《西游记》"}
    for _, book := range books {
        wg.Add(1)
        go bk.SetName(wg, book)
    }
    wg.Wait()
}

```

#### 2.sync.RWMutex 读写锁

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	var rwMutex sync.RWMutex
	Data := 0
	for i := 0; i < 10; i++ {
		// 读
		go func read(t int) {
			rwMutex.RLock()
			defer rwMutex.RUnlock()
			fmt.Printf("Read data: %v\n", Data)
			wg.Done()
			time.Sleep(2 * time.Second)
			// 这句代码第一次运行后，读解锁。
			// 循环到第二个时，读锁定后，这个goroutine就没有阻塞，同时读成功。
		}(i)
		// 写
		go func write(t int) {
			rwMutex.Lock()
			defer rwMutex.Unlock()
			Data += t
			fmt.Printf("Write Data: %v %d \n", Data, t)
			wg.Done()
			// 这句代码让写锁的效果显示出来，写锁定下是需要解锁后才能写的。
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}
```

> 当写锁定，读不会被锁，读是自由的，写锁会等待上个写锁解锁，才会执行！！

#### 3.sync.Map

```go
package main

import (
	"fmt"
	"strconv"
	"sync"
)

// sync.Map 是一个开箱即用的并发安全的map

var m2 = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)         // 必须使用sync.Map内置的Store方法设置键值对
			value, _ := m2.Load(key) // 必须使用sync.Map提供的Load方法根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

