package main

import "fmt"

//变量赋值
//全局变量赋值
var (
	aaa       = "hello"
	a1, b1, c = 1, 2, 3
)

func assigns() {
	var d, e, f int
	d, e, f = 4, 5, 6
	var h, j, k int = 10, 2, 4
	r, t := 1, 32
	var i, o = 10, 20

	println(i, o)
	println(a, b, c, aaa)
	println(d, e, f)
	println(h, j, k)
	println(r, t)
}

//类型转换
func typeCast() {
	var a float32 = 1000.1
	fmt.Println(a)
	b := int(a)
	fmt.Println(b)
}
