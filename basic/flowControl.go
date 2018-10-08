package main

import "fmt"

//流程控制：流程控制包含分三大类：条件判断，循环控制和无条件跳转

func main() {

	/*********************************** if 条件判断 *********************************/

	//Go 里面 if 条件判断语句中不需要括号，如下代码所示
	x := 5
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	//Go 的 if 还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其 他地方就不起作用了，如下所示
	if y := 20; y > 10 {
		fmt.Println("y is greater than 10")
	} else {
		fmt.Println("y is less than 10")
	}
	// 这个地方如果这样调用就编译出错了，因为 y 是条件里面的变量
	//fmt.Println(y)

	//多个条件的时候如下所示：
	integer := 10
	if integer == 3 {
		fmt.Println("The integer is equal to 3")
	} else if integer < 3 {
		fmt.Println("The integer is less than 3")
	} else {
		fmt.Println("The integer is greater than 3")
	}

	/*********************************** goto 无条件跳转 *********************************/
	gotoExampleFunc()

	/*********************************** for 循环 ****************************************/
	//Go里面最强大的一个控制逻辑就是 for ，它既可以用来循环读取数据，又可以当作 while 来控制逻辑，还能迭代操作。
	//在循环里面有两个关键操作 break 和 continue , break 操作是跳出当前循环， continue 是跳过本次循环。当嵌套 过深的时候， break 可以配合标签使用，即跳转至标签所指定的位置。
	/*
		for expression1; expression2; expression3 {
		//...
		}

		expression1 、 expression2 和 expression3 都是表达式，其中 expression1 和 expression3 是变量声明或者
		函数调用返回值之类的， expression2 是用来条件判断， expression1 在循环开始之前调用， expression3 在每
		轮循环结束之时调用
	*/

	sum1 := 0
	for index := 0; index < 10; index++ {
		sum1 += index
	}
	fmt.Println("sum1 is equal to ", sum1)

	//有些时候如果我们忽略 expression1 和 expression3 ：
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println("sum2 is equal to ", sum2)

	//for 配合 range 可以用于读取 slice 和 map 的数据
	// map 的初始化可以通过 key:val 的方式初始化值
	map2 := map[int]string{1: "one", 2: "two", 3: "three", 4: "four"}
	for k, v := range map2 {
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}
	//由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用 _ 来丢弃不需要的返回值
	for _, v := range map2 {
		fmt.Println("map's val:", v)
	}

	/******************************* switch ****************************************/
	//有些时候你需要写很多的 if-else 来实现一些逻辑处理，这个时候代码看上去就很丑很冗长，而且也不易于以后的维护，这个时候 switch 就能很好的解决这个问题

	//在 Go 中可以把很多值聚合在了一个 case 里面
	i := 3
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}

	// Go 里面 switch 默认相当于每个 case 最后带有 break ，匹配成功后不会自动向下执行其他case，而是跳出整个 switch , 但是可以使用 fallthrough 强制执行后面的case代码
	num := 6
	switch num {
	case 4:
		fmt.Println("The num was <= 4")
		fallthrough
	case 5:
		fmt.Println("The num was <= 5")
		fallthrough
	case 6:
		fmt.Println("The num was <= 6")
		fallthrough
	case 7:
		fmt.Println("The num was <= 7")
		fallthrough
	case 8:
		fmt.Println("The num was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

}
func gotoExampleFunc() {
	i := 0
	// 这行的第一个词，以冒号结束作为标签,标签名是大小写敏感的。
Here:
	println(i)
	i++
	if i == 1000 {
		return
	}
	//Go 有 goto 语句——请明智地使用它。用 goto 跳转到必须在当前函数内定义的标签。例如假设这样一个循环：
	// 跳转到 Here 去
	goto Here
}
