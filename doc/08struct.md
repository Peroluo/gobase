# struct结构体

#### 1.结构体的定义

```go 
type Person struct{
    name string
    age int
}
```

#### 2. 初始化

```go
var p1 Person = Person{
    name: "luoguoxiong",
    age: 28
}
// 获取结构体指针
var p1 *Person = new(Person)
p1.name = "luoguoxiong"
p1.age = 28
// 获取结构体指针
p1:=&Person{
    name:"luoguoxiong",
    age:28
}
```

#### 3. 自定义构造函数

```go
func newPerson(name string, age int) *person{
    return &person{
        name,
        age
    }
}
```

#### 4.结构体方法

```go
func (p perosn)GetName()string{
    return p.name
}
```

```go
func (p *perosn)SetName(name string){
    p.name = name
}
```

#### 5. 匿名结构体

```go
type person struct {
	string
	int
}
p1 := person{
	"周林",
	9000,
}
fmt.Println(p1)
fmt.Println(p1.string)
fmt.Println(p1.int)
```

#### 6.结构体嵌套

```go
package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address // 匿名嵌套结构体
	workPlace
	// address:address
}

type company struct {
	name string
	address
}

func main() {
	p1 := person{
		name: "周林",
		age:  9000,
		address: address{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.address.city)
	// fmt.Println(p1.city) // 先在自己结构体找这个字段,找不到就去匿名嵌套的结构体中查找该字段
	fmt.Println(p1.address.city)
	fmt.Println(p1.workPlace.city)
}
```

> 1. 结构体嵌套访问属性值时，当自身有这个属性，优先返回，找不到再从匿名结构体中查询字段。可以简写`p1.city`访问属性。
>
> 2. 当匿名结构体有多个，并且都含有同一个属性`city`,不能简写，只能`p1.address.city`

#### 7. 继承

```go
type animal struct {
	name string
}
// 给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动!\n", a.name)
}

// 狗类
type dog struct {
	feet   uint8
	animal // animal拥有的方法,dog此时也有了
}

// 给dog实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s在叫:汪汪汪~\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{name: "周林"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.wang
	d1.move() // ?
}

```

#### 8.结构体tag

```go
type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "周林",
		Age:  9000,
	}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b))
	// 反序列化
	str := `{"name":"理想","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) // 传指针是为了能在json.Unmarshal内部修改p2的值
	fmt.Printf("%#v\n", p2)
}
```

