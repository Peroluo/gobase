# Map

#### 1.声明

```go
var m1 map[string]int
fmt.Println(m1 == nil)        // 还没有初始化（没有在内存中开辟空间）
```

#### 2.声明并赋值

```go
m1 := make(map[string]int, 10) // 要估算好该map容量，避免在程序运行期间再动态扩容
m1["理想"] = 18
m1["jiwuming"] = 35
```

#### 3.遍历

```go
// map的遍历
for k, v := range m1 {
fmt.Println(k, v)
}
```

#### 4.删除

```go
delete(m1, "jiwuming")
```

#### 5.判断是否有某个key

```go
value, ok := m1["娜扎"]
if !ok {
    fmt.Println("查无此key")
} else {
    fmt.Println(value)
}
```

