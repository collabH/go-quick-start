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
