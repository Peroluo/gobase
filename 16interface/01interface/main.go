package main

import "fmt"

// 接口嵌套
type animal interface {
	move
	eat
}
type move interface {
	move()
}
type eat interface {
	eat(string)
}

type cat struct {
	name string
}

func (c *cat) move() {
	fmt.Println("猫move")
}
func (c *cat) eat(food string) {
	fmt.Printf("猫eat%s \n", food)
}

func main() {
	cat1 := cat{name: "miaomiao"}
	var a animal = &cat1 // cat实现animal接口
	a.eat("鱼")
	a.move()
}
