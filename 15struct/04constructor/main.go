package main

import "fmt"

type person struct {
	name string
	age  int
}

// 构造函数
// 返回是结构体还是结构体指针
// 当结构体比较大的时候尽量使用结构体指针，减少程序内存开销
func newPerson(name string, age int) *person {
	return &person{
		name,
		age,
	}
}

// func(接受者变量 接受者类型) 方法名(参数列表)(返回参数){函数体}
func (p person) say1() {
	fmt.Printf("%s已经%d岁了！\n", p.name, p.age)
}

// 值接受者
// 接受者是拷贝代价比较大的大对象
// 保证一致性，如果某个方法使用了值接受者类型，要保持一致
func (p *person) guonian() {
	p.age++
}

func main() {
	p1 := newPerson("luoguoxiong", 12)
	fmt.Println(p1) //&{luoguoxiong 12}
	p1.say1()
	p1.guonian()
	fmt.Println(p1.age)
	p2 := person{
		name: "ewe",
		age:  19,
	}
	(&p2).guonian()
	fmt.Println(p2.age)
}
