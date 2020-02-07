# Go快速入门
## 什么是Go?
> go是一门并发支持、垃圾回收的编译型系统编程语言，旨在创造一门具有在静态编译语言的高性能和动态语言的
>高效开发之间拥有良好平衡点的一门语言。

## Go的主要特点有哪些
* 类型安全和内存安全
* 以非常直观和极低代价的方案实现高并发
* 高效的垃圾回收机制
* 快速编译(同时解决C语言中头文件太多的问题)
* 为多核计算机提供性能提升的方案
* UTF-8编码支持

## 常用命令
### Go命令
```text
在命令行输入go即可查看所有支持的命令
```
### 常用命令
- go get:获取远程包
- go run:直接运行程序
- go build:测试编译，检查是否有编译错误
- go fmt:格式化源码
- go install:编译包文件并编译整个程序
- go test:运行测试文件
- go doc:查看文档

## Go内置关键字
```text
break
default
func
interface
case
defer
go 
map
chan
else
goto
package
continue
for
import
return
select
struct
switch
type
var
```
### 使用方式
- Go程序是通过 package 来组织的（与python类似）
- 只有 package 名称为 main 的包可以包含 main 函数
- 一个可执行程序 有且仅有 一个main包

- 通过 import 关键字来导入其它非 main 包
- 通过 const 关键字来进行常量的定义
- 通过在函数体外部使用 var 关键字来进行全局变量的声明与赋值
- 通过 type 关键字来进行结构(struct)或接口(interface)的声明
- 通过 func 关键字来进行函数的声明

## Go的可见性规则
```text
Go语言中，使用`大小写`来决定该常量、变量、类型、接口、结构或函数 是否可以被外部包所调用：
     根据约定，函数名首字母小写，即为private
     函数名首字母大写，即为public
```
    
## Go的基本类型
- bool
  - 长度:1字节
  - 范围: true,false
  -注意事项: 不可以用数字代表true或false
- 整型: int/uint
  - 根据运行平台可能为32或者64位
- 8位整型: int8/uint8
  - 长度: 1字节
  - 取之范围: -128～127/0～255
- 字节型: byte(uint8别名)
- 16位整型：int16/uint16
  - 长度：2字节
  - 取值范围：-32768~32767/0~65535
- 32位整型：int32（rune）/uint32
  - 长度：4字节
  - 取值范围：-2^32/2~2^32/2-1/0~2^32-1
- 64位整型：int64/uint64
  - 长度：8字节
  - 取值范围：-2^64/2~2^64/2-1/0~2^64-1
- 浮点型：float32/float64
  - 长度：4/8字节
  - 小数位：精确到7/15小数位
- 复数：complex64/complex128
  - 长度：8/16字节
- 足够保存指针的 32 位或 64 位整数型：uintptr
- 其它值类型：
  - array、struct、string
- 引用类型：
  - slice、map、chan
- 接口类型：inteface
- 函数类型：func

## 类型默认值（零值）
```text
  通常值类型默认值都为0，bool为false，string为空字符串
```
## 变量赋值与声明
### 单个变量

- 变量的声明格式：var <变量名称> <变量类型>
- 变量的赋值格式：<变量名称> = <表达式>
- 声明的同时赋值：var <变量名称> [变量类型] = <表达式>
```go
	var b uint32
	b=190

	a:=10
	var c  ="hello"
	println(b,a,c)
```

### 多个变量
- 全局变量的声明可使用 var() 的方式进行简写
- 全局变量的声明不可以省略 var，但可使用并行方式
- 所有变量都可以使用类型推断
- 局部变量不可以使用 var() 的方式简写，只能使用并行方式

```go
package main

//变量赋值
//全局变量赋值
var (
	aaa     = "hello"
	a, b, c = 1, 2, 3
)

func assgin() {
   //并行方式
	var d, e, f int
	d, e, f = 4, 5, 6
	var h, j, k int = 10, 2, 4
	r,t:=1,32
	var i, o  =10,20
   //忽略赋值
   var _,p,l=1,2,3
    print(p,l)

	println(i,o)
	println(a, b, c, aaa)
	println(d, e, f)
	println(h, j, k)
	println(r,t)
}

```
## Go的类型转换
```text
Go中不存在隐式转换，所有类型转换必须显式声明
 转换只能发生在两种相互兼容的类型之间
 类型转换的格式：
 	<ValueA> [:]= <TypeOfValueA>(<ValueB>)
```
```go
func typeCast()  {
	var a float32=1000.1
	fmt.Println(a)
	b:=int(a);
	fmt.Println(b)
}
```
## 常量
### 常量的定义
- 常量的值在编译时就已经确定
- 常量的定义格式与变量基本相同
- 等号右侧必须是常量或者常量表达式
- 常量表达式中的函数必须是内置函数

### 常量的初始化规则与枚举
```text
在定义常量组时，如果不提供初始值，则表示将使用上行的表达式
使用相同的表达式不代表具有相同的值
iota是常量的计数器，从0开始，组中每定义1个常量自动递增1
通过初始化规则与iota可以达到枚举的效果
每遇到一个const关键字，iota就会重置为0

```
## 运算符
- Go中的运算符均是从左到右结合
**优先级(从高到低)**
- `^` `!`                         (一元运算符)
- `*` `/` `%` `<<` `>>` `&` `&^`  
- `+` `-` `|` `^`                 (二元运算符)
- `==` `!=` `<` `<=` `>=` `>`
- `<-`                            (专门用于channel)
- `&&`
- `||`

## 指针
```text
Go虽然保留了指针，但与其它编程语言不同的是，在Go当中不
支持指针运算以及”->”运算符，而直接采用”.”选择符来操作指针
目标对象的成员

操作符”&”取变量地址，使用”*”通过指针间接访问目标对象
默认值为 nil 而非 NULL


递增递减语句

        在Go当中，++ 与 -- 是作为语句而并不是作为表达式
```

```go
//指针
func cursor()  {
	a:=1
	var p *int=&a;
	println(*p)
	a++
	println(*p)
}
```
## 控制语句
### if
- 条件表达式没有括号
- 支持一个初始化表达式（可以是并行方式）
- 左大括号必须和条件语句或else在同一行
- 支持单行模式
- 初始化语句中的变量为block级别，同时隐藏外部同名变量
- 1.0.3版本中的编译器BUG
```go
//if语句
	if a := 1; a > 1 {
		println("yes")
	} else {
		println("no")
	}
```
### for
- Go只有for一个循环语句关键字，但支持3种形式
- 初始化和步进表达式可以是多个值
- 条件语句每次循环都会被重新检查，因此不建议在条件语句中
- 使用函数，尽量提前计算好条件并以变量或常量代替
- 左大括号必须和条件语句在同一行
```go
//for语句 类似于java
	//one
	for i:=0;i<10;i++  {
		print(i)
	}
	println()
	//two 无限循环需要退出状态
	a:=1
	for  {
		a++
		if a > 10 {
			break
		}
		print(a)
	}
	println()
	//three 类似于while(b<10){b++}
	b:=1
	for b<10  {
		b++
		print(b)
	}
```

### switch
- 可以使用任何类型或表达式作为条件语句
- 不需要写break，一旦条件符合自动终止
- 如希望继续执行下一个case，需使用fallthrough语句
- 支持一个初始化表达式（可以是并行方式），右侧需跟分号
- 左大括号必须和条件语句在同一行
```go
//switch
	//one
	p:=1
	switch p {
	case 0:
		println(0)
	case 1:
		println(1)
	default:
		println("default")
	}
	//two
	z:=1
	switch {
	case z>0:
		println(z)
	case z<0:
		println("-" )
	default:
		println("not found")
	}
	//three
	switch z:=1; {
	case a > 0:
		println(z)
	case a<0:
		println("zzz")
	default:
		println("not found")
	}
```
### 跳转语句goto，break，continue
- 三个语法都可以配合标签使用
- 标签名区分大小写，若不使用会造成编译错误
- break与continue配合标签可用于多层循环的跳出
- goto是调整执行位置，与其他2个语句配合标签的结果并不相同

```go
LABEL:
	for {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				break LABEL
			} else {
				println(i)
			}
		}
	}
	// continue
Flag:
	for i := 0; i < 3; i++ {
		for {
			println(i)
			continue Flag
		}
	}

	//goto 
	for i := 0; i < 3; i++ {
		for {
			println(i)
			goto L
		}
	}
L:
```
## 数组array
- 定义数组的格式：var <varName> [n]<type>，n>=0
- 数组长度也是类型的一部分，因此具有不同长度的数组为不同类型
- 注意区分指向数组的指针和指针数组
- 数组在Go中为值类型
- 数组之间可以使用==或!=进行比较，但不可以使用<或>
- 可以使用new来创建数组，此方法返回一个指向数组的指针
- Go支持多维数组