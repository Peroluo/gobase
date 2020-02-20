# var 变量声明

#### 1. 单个声明

```go
// var 变量名 变量类型
var name string
```

#### 2. 批量声明

```go
// var ()
var (
	name string // ""
    age int // 0
    isok bool // false
)
```

> 注意：默认声明的值都是对应类型的零值

#### 3.声明变量并赋值

```go
var s1 string = "who"
```

#### 4. 类型推倒

```go
var s2 = "20"
```

#### 5.简单变量声明

```go
s1:="string"
```

> 注意：简单变量声明，只能用于函数内部