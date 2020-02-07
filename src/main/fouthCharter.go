package main

//常量的定义
const a int = 1
const b = 'A'
const (
	text  = "123"
	lenth = len(text)
	num   = b * 20
)
const i, j, k = 1, "2", '3'
const (
	text2, lenth2, mum2 = "456", len(text2), k * 10
)

func constDefinition() {
	println(a, b)
}

//枚举 iota计数器
const (
	one = 'A'
	two
	three = iota
	four  = 2
)

func constAndEnum() {
	four := 10
	println(four)
	println(a, b)
}

//运算符
func operator() {
	println(^1)
	println(!true)
	println(1 << 10)
	println(2 << 1)
	/**
	 4: 0100
	 8: 1000
	---------
	 &  0000  0
	 |  1100  12
	 ^  1100  12
	 &^ 0100  4
	*/
	println(4&8,
		4|8,
		4^8,
		4&^8,
		false && true,
		false || true)
	//&&短路与
}

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB float64 = 1 << (iota * 10)
	ZB float64 = 1 << (iota * 10)
	YB float64 = 1 << (iota * 10)
)

//输出计算机存储单位
func printComputerStorageUnit() {
	println(KB, MB, GB, TB, PB, EB, ZB, YB)
}

//指针
func cursor() {
	a := 1
	var p *int = &a
	println(*p)
	a++
	println(*p)
}

func conditionExpression() {
	//if语句
	if a := 1; a > 1 {
		println("yes")
	} else {
		println("no")
	}
	//for语句 类似于java
	//one
	for i := 0; i < 10; i++ {
		print(i)
	}
	println()
	//two 无限循环需要退出状态
	a := 1
	for {
		a++
		if a > 10 {
			break
		}
		print(a)
	}
	println()
	//three 类似于while(b<10){b++}
	b := 1
	for b < 10 {
		b++
		print(b)
	}

	//switch
	//one
	p := 1
	switch p {
	case 0:
		println(0)
	case 1:
		println(1)
	default:
		println("default")
	}
	//two
	z := 1
	switch {
	case z > 0:
		println(z)
	case z < 0:
		println("-")
	default:
		println("not found")
	}
	//three
	switch z := 1; {
	case a > 0:
		println(z)
	case a < 0:
		println("zzz")
	default:
		println("not found")
	}

	//break
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

	//goto 注意标签位置放置前面会造成死循环
	for i := 0; i < 3; i++ {
		for {
			println(i)
			goto L
		}
	}
L:
}
