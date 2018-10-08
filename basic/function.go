package main

import (
	"fmt"
	"os"
)

/*
函数是Go里面的核心设计，它通过关键字 func 来声明，它的格式如下：
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
	// 这里是处理逻辑代码
	// 返回多个值
	return value1, value2
}

- 关键字 func 用来声明一个函数 funcName
- 函数可以有一个或者多个参数，每个参数后面带有类型，通过 , 分隔
- 函数可以返回多个值
- 上面返回值声明了两个变量 output1 和 output2 ，如果你不想声明也可以，直接就两个类型
- 如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
- 如果没有返回值，那么就直接省略最后的返回信息
- 如果有返回值， 那么必须在函数的外层添加return语句
*/
func main() {
	a := 2
	b := 3
	fmt.Printf("max = %v \n", max(a, b))

	add, multiplied := addAndMultiplied(3, 4)
	fmt.Printf("add = %v, multiplied = %v \n", add, multiplied)

	add, multiplied = addAndMultiplied(5, 6)
	fmt.Printf("add = %v, multiplied = %v \n", add, multiplied)

	//	可变参数
	variableParamFunc(1, 2, 3, 4, 5, 6)

	/*传值与传指针传值与传指针*/
	//	当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一份copy，当在被调用函数中修改参数值的时候，调用函数中相应实参不会发生任何变化，因为数值变化只作用在copy上。
	aa := 2
	addOne(aa)
	fmt.Printf("传值，修改后的值 aa = %d \n", aa)

	bb := 2
	//变量在内存中是存放于一定地址上的，修改变量实际是修改变量地址处的内 存。 只有变量所在的地址，才能修改变量的值。
	// 所以我们需要将 bb 所在地址 &bb 传入函数，并将函数的参数的类型由 int 改为 *int ，即改为指针类型，才能在函数中修改 bb 变量的值。
	// 此时参数仍然是按copy传递的，只是copy的是一个指针。
	// 同过在变量前加 &  取该变量的地址
	addOneWithPointer(&bb)
	fmt.Printf("传值，修改后的值 aa = %d \n", bb)

	/*
		传指针的好处
			- 传指针使得多个函数能操作同一个对象。
			- 传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
			- Go 语言中 string ， slice ， map 这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变 slice 的长度，则仍需要取地址传递指针）
	*/

	/************************************* deffer ****************************************/
	//Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。
	// 特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题。

	/****************************** 函数作为值、类型 *************************************/
	//	在 Go 中函数也是一种变量，我们可以通过 type 来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型。
	// 函数当做值和类型在我们写一些通用接口的时候非常有用，一个函数类型类似策略接口，可以接收不同的策略类
	// 声明函数类型 type typeName func(input1 inputType1 [, input2 inputType2 [, ...]) (result1 resultType，result2 resultType)

	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	//  把函数当做值来传递了，只要奇数
	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice are: ", odd)
	//  把函数当做值来传递了，只要偶数
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)

	/**************************** Panic和Recover **************************************/
	//	Go没有像Java那样的异常机制，它不能抛出异常，而是使用了 panic 和 recover 机制。
	// 一定要记住，你应当把它作为最后的手段来使用，也就是说，你的代码中应当没有，或者很少有 panic 的东西。

	/*
		Panic
		是一个内建函数，可以中断原有的控制流程，进入一个令人恐慌的流程中。当函数 F 调用 panic ，函数F的执行被中
		断，但是 F 中的延迟函数会正常执行，然后F返回到调用它的地方。在调用的地方， F 的行为就像调用了 panic 。这一
		过程继续向上，直到发生 panic 的 goroutine 中所有调用的函数返回，此时程序退出。恐慌可以直接调用 panic 产
		生。也可以由运行时错误产生，例如访问越界的数组。

		Recover
		是一个内建的函数，可以让进入令人恐慌的流程中的 goroutine 恢复过来。 recover 仅在延迟函数中有效。在正常
		的执行过程中，调用 recover 会返回 nil ，并且没有其它任何效果。如果当前的 goroutine 陷入恐慌，调用
		recover 可以捕获到 panic 的输入值，并且恢复正常的执行。
	*/

	/**************************** main 函数和 init 函数 *********************************/

	/*
		Go里面有两个保留的函数： init 函数（能够应用于所有的 package ）和 main 函数（只能应用于 package main ）。
	这两个函数在定义时不能有任何的参数和返回值。虽然一个 package 里面可以写任意多个 init 函数，但这无论是对
	于可读性还是以后的可维护性来说，我们都强烈建议用户在一个 package 中每个文件只写一个 init 函数。
	Go程序会自动调用 init() 和 main() ，所以你不需要在任何地方调用这两个函数。每个 package 中的 init 函数都是
	可选的，但 package main 就必须包含一个 main 函数。

		程序的初始化和执行都起始于 main 包。如果 main 包还导入了其它的包，那么就会在编译时将它们依次导入。有时一
	个包会被多个包同时导入，那么它只会被导入一次（例如很多包可能都会用到 fmt 包，但它只会被导入一次，因为没
	有必要导入多次）。当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包导入进来，然后再对这些包
	中的包级常量和变量进行初始化，接着执行 init 函数（如果有的话），依次类推。等所有被导入的包都加载完毕
	了，就会开始对 main 包中的包级常量和变量进行初始化，然后执行 main 包中的 init 函数（如果存在的话），最后执
	行 main 函数。
	*/

	/*
		上面这个fmt是Go语言的标准库，其实是去goroot下去加载该模块，当然Go的import还支持如下两种方式来加载自己
	写的模块：
	60
	1. 相对路径
	import “./model” //当前文件同一目录的model目录，但是不建议这种方式来import
	2. 绝对路径
	import “shorturl/model” //加载gopath/src/shorturl/model模块
	上面展示了一些import常用的几种方式，但是还有一些特殊的import，让很多新手很费解，下面我们来一一讲解一下
	到底是怎么一回事
	1. 点操作
	我们有时候会看到如下的方式导入包
	import(
	. "fmt"
	)
	这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调
	用的fmt.Println("hello world")可以省略的写成Println("hello world")
	2. 别名操作
	别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字
	import(
	f "fmt"
	)
	别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("hello world")
	3. _操作
	这个操作经常是让很多人费解的一个操作符，请看下面这个import
	import (
	"database/sql"
	_ "github.com/ziutek/mymysql/godrv"
	)
	_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数
	*/
}

// 返回 a 、b 中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// a + b， a * b
func addAndMultiplied(a, b int) (int, int) {
	return a + b, a * b
}

// a + b， a * b
//官方建议：最好命名返回值，因为不命名返回值，虽然使得代码更加简洁了，但是会造成生成的文档可读性差。
func addAndMultiplied2(a, b int) (add int, multiplied int) {
	return a + b, a * b
}

//Go 函数支持变参。接受变参的函数是有着不定数量的参数的。为了做到这点，首先需要定义函数使其接受变参：
//arg ...int 告诉 Go 这个函数接受不定数量的参数。
// 注意，这些参数的类型全部是 int 。在函数体中，变量 arg 是一 个 int 的 slice
func variableParamFunc(arg ...int) {
	for index, n := range arg {
		fmt.Printf("this index is %d and the number is: %d\n", index, n)
	}
}

// a + 1
func addOne(a int) {
	a = a + 1
}

// a + 1
func addOneWithPointer(a *int) {
	*a = *a + 1
}

// 是否为奇数
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

// 是否为偶数
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 声明函数类型
type strategy func(a int) bool

//  声明的函数类型在这个地方当做了一个参数
func filter(slice []int, s strategy) []int {
	var result []int
	for _, value := range slice {
		if s(value) {
			result = append(result, value)
		}
	}
	return result
}

var user = os.Getenv("USER")

func init() {
	if user == "" {
		panic("no value for $USER")
	}
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	// 执行函数 f ，如果 f 中出现了 panic ，那么就可以恢复回来
	f()
	return
}
