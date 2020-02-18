package main

import (
	"fmt"
	"time"
)

func main()  {
	for i:=0;i<1000;i++{
		go func(i int)  {
			fmt.Println("hello",i)
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
	fmt.Println("？？")
}