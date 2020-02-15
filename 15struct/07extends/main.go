package main

import "fmt"

type animal struct {
	name string
}

func (a animal) run() {
	fmt.Printf("%s移动!\n", a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d dog) wangwang() {
	fmt.Printf("%s旺旺！\n", d.name)
}
func main() {
	dog1 := dog{
		feet: 4,
		animal: animal{
			name: "小黄",
		},
	}
	fmt.Println(dog1)
	dog1.run()
	dog1.wangwang()
}
