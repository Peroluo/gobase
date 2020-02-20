# array数组

#### 1. 数组声明

```go
// 声明一个长度为3，类型为string的数组a
var a [3]string // [false,false,false]
```

#### 2.初始化

##### 初始化1

```go
a:=[3]bool[true,true,false]
```

##### 根据初始化值自动推断数组长度

```go
a:=[...]int[1,2,3,4]
```

##### 根据索引短

```go
a:=[5]int{0:1,4:2}
```

#### 3.数组变历

```go
for i := 0; i < len(citys); i++ {
    fmt.Println(citys[i])
}
// 2. for range遍历
for i, v := range citys {
    fmt.Println(i, v)
}
```

#### 4.多维数组

```go
a11 := [3][2]int{
    [2]int{1, 2},
    [2]int{3, 4},
    [2]int{5, 6},
}
```
