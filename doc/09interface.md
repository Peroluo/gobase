# interface 接口

#### 1.接口定义

```go
type spreaker interfacr{
    run()
    eat()
}
```

#### 2.结构体实现接口

```go
package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步...")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	// 实现接口
	var a1 animal = cat{
		name: "猫",
		feet: 8,
	}
	fmt.Printf("%T\n", a1)
	fmt.Println(a1)
}
```

> 如果结构体方法是对象是指针类型，则实现接口必须通过指针获取对象（new或&）

#### 3.匿名接口

```go
package main

import "fmt"

// 空接口

// interface: 关键字
// interface{} :空接口类型

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "周林"
	m1["age"] = 9000
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)
}

```

