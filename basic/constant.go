package main

import "fmt"

func main() {
	// 如果需要，也可以明确指定常量的类型
	const Pi float32 = 3.1415926
	fmt.Println(Pi)

	// 可以不同指定常量类型，编译器会根据初始化的值自动推导出相应的类型
	const prefix = "g-senlin"
	fmt.Println(prefix)

	//在 Go 中，布尔值的类型为 bool ，值是 true 或 false ，默认为 false
	const sucess = true
	fmt.Println(sucess)
}
