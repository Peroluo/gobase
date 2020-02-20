# const常量声明

#### 1. 单个声明

```go
const pi = 3.14
```

#### 2.批量声明

```go
// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)// n1=n2=n3
```

#### 3.iota(1)

```go
// iota
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)
```

#### 4.iota(2)

```go
const (
	b1 = iota // 0
	b2 = iota // 1
	_  = iota // 2
	b3 = iota // 3
)
```

#### 5.iota(插队)

```go
// 插队
const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4
)
```

#### 6.多个常量声明

```go
// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1:1 d2:2
	d3, d4 = iota + 1, iota + 2 // d3:2 d4:3
)
```

#### 7. 定义数量级

```go
// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
```

#### 总结： iota从0开始，往下依次加1

