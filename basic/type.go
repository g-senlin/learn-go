package main

import (
	"errors"
	"fmt"
)

/*
1、基础类型底层都是分配了一块内存，然后存储了相应的值.
2、Go之所以会那么简洁，是因为它有一些默认的行为：
	2.1、大写字母开头的变量是可导出的，也就是其它包可以读取的，是公用变量；小写字母开头的就是不可导出的，是私有变量
	2.2、大写字母开头的函数也是一样，相当于 class中的带 public 关键词的公有函数；小写字母开头的就是有 private 关键词的私有函数。

*/
func main() {
	// type 关键字
	//go语言支持type关键字，可以定义一种类型,但是有个关键点不同：不是同一个类型
	type Int int
	var intType int = 22
	var IntType Int

	// 下面被注释的这个语句会编译出错，因为 intType 是 int 类型，和 Int 不是同一类型，可以通过强制转换来实现下面的赋值
	//IntType = intType
	//强制转换 intType 为 Int 来实现赋值
	IntType = Int(intType)
	fmt.Println(IntType)

	/************************ Boolean 类型 ********************************/
	//在 Go 中，布尔值的类型为 bool ，值是 true 或 false ，默认为 false
	//布尔类型不可以被赋值为其它类型，也不接受自动或者强制类型转换

	var isSuccess = true
	fmt.Println(isSuccess)

	//enabled 被推导为bool类型
	enabled := (1 == 2)
	fmt.Println(enabled)

	/************************ 数值类型 ********************************/
	// Go里面也有直接定义好位数的类型： rune ,  int8 ,  int16 ,  int32 ,  int64 和 byte ,  uint8 ,  uint16 , uint32 ,  uint64 。其中 rune 是 int32 的别称， byte 是 uint8 的别称
	// 这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错。另外，尽管int的长度是32 bit, 但int 与 int32并不可以互用
	//var a int = 2
	//var b int32 = 100
	//b = a
	//a=b

	//浮点数的类型有 float32 和 float64 两种（没有 float 类型），默认是 float64
	var af = 2.3

	// 下面被注释的代码会报编译错误，小数默认类型是 float64 不能直接赋值给 float32 类型变量
	//var bf float32
	//bf = af
	//fmt.Println(bf)

	//小数默认类型是 float64 可以直接赋值给 float64 类型变量
	var cf float64
	cf = af
	fmt.Println(cf)

	//	Go还支持复数类型 complex64 (32位实数+32位虚数)，默认类型是 complex128 和 complex128 （64位实数+64位虚数）。复数的形式为 RE + IMi ，其中 RE 是实数部分， IM 是虚数部分，而最后的 i 是虚数单位
	var ac complex128 = 1 + 5i
	fmt.Printf("Value is: %v", ac)

	/************************ 字符串 ********************************/
	//	Go中的字符串都是采用 UTF-8 字符集编码。字符串是用一对双引号（ "" ）或反引号（ ` ` ）括起来定义，它的类型是 string

	// 字符串声明
	var as string
	//  声明了一个字符串变量，初始化为空字符串
	var bs string = ""
	// true
	fmt.Println(as == bs)

	var cs string = "hello go"
	// 下面注释的这行代码会报错，在Go中字符串是不可变的
	//cs[0] = 'H'

	// 如果真的想要修改字符串可以将字符串 cs  转换为 []byte  类型，修改后再转换回 string  类型
	//  将字符串 cs  转换为 []byte  类型
	c := []byte(cs)
	c[0] = 'H'
	//  再转换回 string  类型
	cs2 := string(c)
	fmt.Printf("%s\n", cs2)

	// Go中可以使用 + 操作符来连接两个字符串：
	ds := "hello,"
	es := " world"
	fs := ds + es
	fmt.Printf("%s\n", fs)

	// 修改字符串也可以通过把字符串进行切片操作
	s := "hello"
	s = "H" + s[1:]
	// 输出值：Hello
	fmt.Printf("%s\n", s)
	// 可以通过 ` 来声明一个多行字符串， ` 括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。
	m := `hello
	world`
	fmt.Println(m)

	/************************ 错误类型 ********************************/
	// Go内置有一个 error 类型，专门用来处理错误信息，Go 的 package 里面还专门有一个包 errors 来处理错误：
	err := errors.New("test error")
	if err != nil {
		fmt.Print(err)
	}

	/************************** 分组声明 *************************************/
	//在Go语言中，同时声明多个常量、变量，或者导入多个包时，可采用分组的方式进行声明

	/*
		//导入多个包
		import(
			"fmt"
		45
		"os"
		)

		// 声明多个常量，除非被显式设置为其它值或 iota ，每个 const 分组的第一个常量被默认设置为它的0值
		//第二及后续的常量被默认设置为它前面那个常量的值，如果前面那个常量的值是 iota ，则它也被设置为 iota
		const(
			i = 100
			pi = 3.1415
			prefix = "Go_"
		)

		// 声明多个变量
		var(
			i int
			pi float32
			prefix string
		)
	*/

	/************************** iota 枚举 *************************************/
	//Go里面有一个关键字 iota ，这个关键字用来声明 enum 的时候采用，它默认开始值是 0 ，每调用一次加 1
	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        //  常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说 w = iota ，因此 w == 3 。其实上面 y 和 z 可同样不用 "= iota"
	)
	fmt.Printf("x = %d, y = %d, z = %d, w = %d \n", x, y, z, w)

	//  每遇到一个 const 关键字， iota 就会重置，此时 v == 0
	const v = iota
	fmt.Printf("v = %d \n", v)

	/************************** array 数组 *************************************/
	// 数组的定义：var arr [n]type
	// 在 [n]type 中， n 表示数组的长度， type 表示存储元素的类型。对数组的操作和其它语言类似，都是通过 [] 来进行读取或赋值
	// 由于长度也是数组类型的一部分，因此 [3]int 与 [4]int 是不同的类型，数组也就不能改变长度。
	// 数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。如果要使用指针，就需要用到 slice 类型了
	//  声明了一个 int 类型的数组
	var arr [10]int
	// 数组下标是从 0 开始的
	arr[0] = 10
	// 赋值操作
	arr[1] = 20
	fmt.Printf("The first element is %d\n", arr[0])
	// 未赋值的数组元素，返回其元素类型的默认值
	fmt.Printf("The last element is %d\n", arr[9])

	//数组可以使用另一种 := 来声明
	// 声明了一个长度为 3 的 int 数组
	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)

	// 声明了一个长度为 10 的 int 数组，其中前三个元素初始化为 1 、 2 、 3 ，其它默认为 0
	arr2 := [10]int{1, 2, 3}
	fmt.Println(arr2)

	// 可以省略长度而采用 `...` 的方式， Go 会自动根据元素个数来计算长度
	arr3 := [...]int{4, 5, 6}
	fmt.Println(arr3)

	//  声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有 4 个 int 类型的元素
	arr4 := [2][3]int{[3]int{11, 12, 13}, [3]int{21, 22, 23}}
	fmt.Println(arr4)

	/************************** slice 切片 *************************************/
	// 在很多应用场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫 slice
	// slice 并不是真正意义上的动态数组，而是一个引用类型。 slice 总是指向一个底层 array ， slice 的声明也可以像 array 一样，只是不需要长度。
	// slice 的声明和声明 array 一样，只是少了长度
	//var variableName  []type

	//声明一个 slice ，并初始化数据
	slice := []byte{'a', 'b', 'c', 'd'}
	//[97 98 99 100]
	fmt.Println(slice)

	/*
		slice有一些简便的操作
		slice 的默认开始位置是0， ar[:n] 等价于 ar[0:n]
		slice 的第二个序列默认是数组的长度， ar[n:] 等价于 ar[n:len(ar)]
		如果从一个数组里面直接获取 slice ，可以这样 ar[:] ，因为默认第一个序列是0，第二个是数组的长度，即
		等价于 ar[0:len(ar)]
	*/
	// slice 可以从一个数组或一个已经存在的 slice 中再次声明。 slice 通过 array[i:j] 来获取，其中 i 是数组的开始位置， j 是结束位置(包含头不包含尾)，但不包含 array[j] ，它的长度是 j-i
	//从数组中声明
	slice1 := slice[1:3]
	//[98 99]
	fmt.Println(slice1)

	slice2 := slice[1:]
	//[98 99 100]
	fmt.Println(slice2)

	slice3 := slice[:3]
	//[97 98 99]
	fmt.Println(slice3)

	slice4 := slice[:]
	//[97 98 99 100]
	fmt.Printf("slice4 = %v \n", slice4)

	//从已经存在的 slice 中声明
	slice5 := slice4[1:3]
	fmt.Printf("slice5 = %v \n", slice5)

	// slice 是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值，slice4 与 slice5 引用同一个数组，所以改变 slice5 导致 slice4 也受到影响
	slice5[0] = 222
	//slice4 = [97 222 99 100]
	fmt.Printf("slice4 = %v \n", slice4)

	/*
		从概念上面来说 slice 像一个结构体，这个结构体包含了三个元素：
			- 一个指针，指向数组中 slice 指定的开始位置
			- 长度，即 slice 的长度
			- 最大长度，也就是 slice 开始位置到数组的最后位置的长度

		对于 slice 有几个有用的内置函数：
			len 获取 slice 的长度
			cap 获取 slice 的最大容量
			append 向 slice 里面追加一个或者多个元素，然后返回一个和 slice 一样类型的 slice
			copy 函数 copy 从源 slice 的 src 中复制元素到目标 dst ，并且返回复制的元素的个数
		注： append 函数会改变 slice 所引用的数组的内容，从而影响到引用同一数组的其它 slice 。 但当 slice 中没有剩
		余空间（即 (cap-len) == 0 ）时，此时将动态分配新的数组空间。返回的 slice 数组指针将指向这个空间，而原
		数组的内容将保持不变；其它引用此数组的 slice 则不受影响。
	*/

	/************************** map *************************************/
	//map 的格式为 map[keyType]valueType
	// map 的读取和设置也类似 slice 一样，通过 key 来操作，只是 slice 的 index 只能是｀int｀类型，而 map 多了很多类型，可以是 int ，可以是 string 及所有完全定义了 == 与 != 操作的类型

	/*
		map过程中需要注意的几点：
			-  map 是无序的，每次打印出来的 map 都会不一样，它不能通过 index 获取，而必须通过 key 获取
			-  map 的长度是不固定的，也就是和 slice 一样，也是一种引用类型
			- 内置的 len 函数同样适用于map ，返回 map 拥有的 key 的数量
			-  map 的值可以很方便的修改，通过 numbers["one"]=11 可以很容易的把key为one 的字典值改为 11
	*/

	//  声明一个 key 是字符串，值为 int 的字典 , 这种方式的声明需要在使用之前使用 make 初始化
	var map1 map[string]int
	map1 = make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3
	//map[one:1 two:2 three:3]
	fmt.Println(map1)

	// map 的初始化可以通过 key:val 的方式初始化值
	map2 := map[int]string{1: "one", 2: "two", 3: "three", 4: "four"}
	// map[1:one 2:two 3:three 4:four]
	fmt.Println(map2)

	// map 有两个返回值，第一个返回值是对应 key 的value 值，第二个返回值，如果不存在 key ，那么 ok 为 false ，如果存在 ok 为 true
	value, ok := map2[2]
	if ok {
		fmt.Printf("map2 exist key 2, map[2] = %v \n", value)
	} else {
		fmt.Printf("map2 not exist key 2, map2 = %v \n", map2)
	}

	//删除 map 中的 key ,如果 key 不存在就不操作
	delete(map2, 1)
	//map[2:two 3:three 4:four]
	fmt.Println(map2)

	//	map 和 slice 一样都是引用类型，如果两个 map 同时指向一个底层，那么一个改变，另一个也相应的改变
	map3 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	map4 := map3

	delete(map3, "key1")
	//map[key2:value2 key3:value3]
	fmt.Println(map4)

	map4["key2"] = "key2-key2"
	//map[key2:key2-key2 key3:value3]
	fmt.Println(map3)

	/*
		make、new操作
		make 用于内建类型（ map 、 slice 和 channel ）的内存分配。 new 用于各种类型的内存分配。
		内建函数 new 本质上说跟其它语言中的同名函数功能一样： new(T) 分配了零值填充的 T 类型的内存空间，并且返回其
		地址，即一个 *T 类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型 T 的零值。有一点非常重要：
		new 返回指针。
		内建函数 make(T, args) 与 new(T) 有着不同的功能，make只能创建 slice 、 map 和 channel ，并且返回一个有初
		始值(非零)的 T 类型，而不是 *T 。本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被
		初始化。例如，一个 slice ，是一个包含指向数据（内部 array ）的指针、长度和容量的三项描述符；在这些项目被
		初始化之前， slice 为 nil 。对于 slice 、 map 和 channel 来说， make 初始化了内部的数据结构，填充适当的值。

	*/
}
