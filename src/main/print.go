package main

import (
	"fmt"
	_ "fmt"
	"time"
)

//常量的定义
const (
	PI     = 3.14
	const1 = 1
	const2 = 2
)

//全局变量的声明与赋值
var name = "go-study"

//一般类型声明
type newType int

//结构的声明
type user struct {
}

//接口的声明
type login interface {
}

func main() {
	fmt.Print(name)
	print(time.After(10))
	println(time.RubyDate)
	Test()
	println(const1 == const2)
	var a int
	println(a)
}
