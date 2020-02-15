package main

import (
	"fmt"
	"strings"
)

// 字符串是UTF-8编码 ""
// 字符是''
// 字节 1字节 = 8Bit(8个二进制)
// 1个字符=1个字节
// 1个utf-8编码汉子 = 一般3个字节
func main() {
	// s := "字符串"
	// c := 'A'
	// 多行字符串
	s1 := `
		luoguoxiong
		lala
	`
	fmt.Print(s1)
	// 字符串常用方法
	// len() 字符串长度
	s2 := "luoguoxiong"
	fmt.Println(len(s2))
	// 字符串拼接
	s3 := s2 + s2
	fmt.Println(s3)
	// 字符串分割
	ref := strings.Split(s2, "uo")
	fmt.Println(ref)
	// 包含
	fmt.Println(strings.Contains(s2, "luo"))
	// 前缀
	fmt.Println(strings.HasPrefix(s2, "luo"))
	// 后缀
	fmt.Println(strings.HasSuffix(s2, "luo"))
	// 某个字符的位置
	fmt.Println(strings.Index(s2, "luo"))
	fmt.Println(strings.LastIndex(s2, "guo"))
	// 拼接 ref["2","23"]
	fmt.Println(strings.Join(ref, "+"))
	for index, val := range s2 {
		fmt.Printf("%d\n", index)
		fmt.Printf("%c\n", val)
	}
	// 字符串修改 go字符串是不能修改的！！
	s4 := "白萝卜"
	// 中文rune,英文byte!!
	s5 := []rune(s4)        //把字符串强制转换成rune切片
	s5[0] = '红'             //s5是由字符组成的切片
	fmt.Println(string(s5)) // 把rune切片转换成字符串
	// 单引号与双引号的区别
	c1 := "H"
	c2 := '红'
	fmt.Printf("c1:%T,c2:%T", c1, c2) //c1:string,c2:int32(字符是由int32构成？？)

}
