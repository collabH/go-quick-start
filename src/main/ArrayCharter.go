package main

import "fmt"

//数组学习

func arrayDemo() {
	//长度不一致无法进行赋值，存在类型默认值
	var a [2]int
	var b [2]int
	b = a
	fmt.Println(b) //[0,0]

	//直接赋值
	arr := [2]int{1}
	println(arr[0])

	//索引方式,将索引为19的元素设置为1
	arr1 := [20]int{19: 1}
	fmt.Println(arr1)

	//自动计算长度
	arr2 := [...]int{1, 2, 3, 4}
	println(len(arr2))

	//去数组的指针 保存指针指向的值
	arr3 := [...]int{99: 1}
	var p *[100]int = &arr3
	println(p)
	//指针数组 保存两个指针的地址
	x, y := 1, 2
	arr4 := [...]*int{&x, &y}
	fmt.Println(arr4)

	//数组的表，需要类型一致,长度不同也无法进行比较
	var arr5 [10]string
	var arr6 [10]string
	println(arr5 == arr6)

	//多维数组
	arr7 := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	fmt.Print(len(arr7), len(arr7[0]), arr7)
}
