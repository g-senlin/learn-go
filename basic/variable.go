package main

import "fmt"

func main() {
	/******************************* 变量定义 ******************************************/
	//Go 对于已声明但未使用的变量会在编译阶段报错
	// 定义变量
	var a int
	a = 1
	fmt.Println(a)

	// 定义变量并初始化
	var aa int = 11
	fmt.Println(aa)

	// 定义多个同类型变量
	var b, c int
	b = 2
	c = 3
	fmt.Println(b, c)

	// 定义多个同类型变量并初始化
	var bb, cc int = 22, 33
	fmt.Println(bb, cc)

	// 编译器会根据初始化的值自动推导出相应的类型
	var bbb, ccc = 222, 333
	fmt.Println(bbb, ccc)

	// 省略 var 和 type 关键字的简写形式
	// 这种简写形式只能用在函数内部，在函数外部使用则会无法编译通过，所以一般用 var 方式来定义全局变量
	bbbb, cccc := 2222, 3333
	fmt.Println(bbbb, cccc)

	// _ （下划线）是个特殊的变量名，任何赋予它的值都会被丢弃
	_, d := 66, 4
	fmt.Println(d)
}
