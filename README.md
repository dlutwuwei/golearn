# golearn

# 变量、类型、关键字

## 声明和赋值

- 1.3使用`:=`声明和赋值一步完成

- 对声明却未使用的变量报错

- 声明语法

```go
var (
  x int
  b bool
)
```
const, import也支持

只声明
```go
var x, y int
```

平行赋值
```
a, b := 20, 16
```

丢弃赋值
```go
`_, b := 34, 35`丢弃_的值
```

## 布尔类型
bool
true false

## 数字类型

int, uint长度有计算机字长决定

int8,int16,int32,int64

byte,uint8,uint16,uint32, uint64

float32 和 float64，没有float

## 常量

```go
const (
  a= iota
  b // 隐含iota
)
```
使用 iota 生成枚举值

## 字符串

一旦给变量赋值,字符串就不能修改了:在 Go 中字符串是不可变的。

```go
s := "Starting part" +
     "Ending part"

s := `Starting part
     Ending part`
```
## 复数

```go
var c complex64 = 5+5i;
fmt.Printf("Value is: %v", c)
```
## 错误

var e error

# 运算符

```
* / % << >> & &^
+ - | ^
== != < <= > >= <-
&&
||
```

虽然 Go 不支持运算符重载(或者方法重载),而一些内建运算符却支持重载。例如 + 可以用于整数、浮点数、复数和字符串(字符串相加表示串联它们)。

# 关键字

# 控制结构

if-else

goto

for

range

switch

