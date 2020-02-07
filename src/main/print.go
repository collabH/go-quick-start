package main

import (
	"fmt"
	_ "fmt"
	"github.com/bwmarrin/snowflake"
	"os"
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
	n, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id", id)
		fmt.Println(
			"node: ", id.Node(),
			"step: ", id.Step(),
			"time: ", id.Time(),
			"\n",
		)
	}
	//变量赋值
	var b uint32
	b = 190

	a := 10
	var c = "hello"
	println(b, a, c)
	//fmt.Print(name)
	//print(time.After(10))
	//println(time.RubyDate)
	//Test()
	//println(const1 == const2)
	//var a int
	//println(a)

	assigns()

	//常量
	constAndEnum()

	//运算符
	operator()

	//输出计算机存储单位
	printComputerStorageUnit()

	//指针
	cursor()

	//条件判断语句
	conditionExpression()

	//数组
	arrayDemo()
}
