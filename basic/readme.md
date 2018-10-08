## Go Basic
### hello go
```go
package main

import "fmt"

func main() {
	fmt.Printf("hello go")
}
````
Go程序是通过 package 来组织的，`package main` 这一行告诉我们当前文件属于哪个包，而包名 main 则告诉我们它是一个可独立运行的包，它在编译后会产生可执行文件。
除了 main 包之外，其它的包最后都会生成 *.a 文件（也就是包文件）并放置在 $GOPATH/pkg 下。
每一个可独立运行的 Go 程序，必定包含一个 `package main` ，在这个 main 包中必定包含一个入口函数 main ，main 函数既没有参数，也没有返回值。

通过关键字 func 定义了一个 main 函数，函数体被放在 {} （大括号）中，就像我们平时写 C 或 Java 时一样。

为了打印 hello go ，我们调用了 fmt 包中的 Printf 函数，通过 `import "fmt"` 导入了系统级别的 fmt 包。

包名和包所在的文件夹名可以是不同的，但一般都定义成相同。

Go 使用 package 来组织代码。 main.main() 函数(这个函数主要位于主包）是每一个独立的可运行程序的入口点。Go使用UTF-8字符串和标识符(因为UTF-8的发明者也就是 Go 的发明者)，所以它天生就具有多
语言的支持。

Go 是天生支持 UTF-8 的，任何字符都可以直接输出，你甚至可以用 UTF-8 中的任何字符作为标识符。

#### 通过 package 来组织的好处：
- 模块化：能够把你的程序分成多个模块
- 可重用性：每个模块都能被其它应用程序反复使用

### 定义变量
- 使用 var 关键字是Go最基本的定义变量方式，与 Java 语言不同的是 Go 把变量类型放在变量名后面
- Go 对于已声明但未使用的变量会在编译阶段报错
-  _ （下划线）是个特殊的变量名，任何赋予它的值都会被丢弃
```go
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
```

### 定义常量
常量是指在程序编译阶段就确定下来的值，而程序在运行时则无法改变该值。在 Go 程序中，常量可定义为数值、布尔值或字符串等类型,且需要在声明时赋值

```go
// 如果需要，也可以明确指定常量的类型
const Pi float32 = 3.1415926
fmt.Println(Pi)

// 可以不同指定常量类型，编译器会根据初始化的值自动推导出相应的类型
const prefix = "g-senlin"
fmt.Println(prefix)

//在 Go 中，布尔值的类型为 bool ，值是 true 或 false ，默认为 false
const sucess = true
fmt.Println(sucess)
```

### 内置类型

#### type关键字
Go 语言支持 type 关键字，可以定义一种类型，这个关键字大体上和C语言的 typedef 相似，但是有个关键点不同：新类型与原来的类型不是同一个类型，不能直接赋值。
     格式如下：
```go
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
```     


#### Boolean
- 在Go中，布尔值的类型为 bool ，值是 true 或 false ，默认为 false
- 布尔类型不可以被赋值为其它类型，也不接受自动或者强制类型转换
```go
var isSuccess = true
fmt.Println(isSuccess)

//enabled 被推导为bool类型
enabled := (1 == 2)
fmt.Println(enabled)
```

####　数值类型

##### 整数
- 整数类型有无符号和带符号两种。 Go 同时支持 int 和 uint ，这两种类型的长度相同，但具体长度取决于不同编译器的实现和操作系统
- Go里面也有直接定义好位数的类型： rune ,  int8 ,  int16 ,  int32 ,  int64 和 byte ,  uint8 ,  uint16 ,uint32 ,  uint64 。其中 rune 是 int32 的别称， byte 是 uint8 的别称
- 预定义的类型，各自间是不一样的，比如在32位机器上，int和int32也是不同的数据类型，不能自动进行类型转换，如果需要，可以用强制类型转换解决赋值或逻辑比较等问题
 ```go
// 这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错。另外，尽管int的长度是32 bit, 但int 与 int32并不可以互用
//var a int = 2
//var b int32 = 100
//b = a
//a = b
```
##### 小数
浮点数的类型有 float32 和 float64 两种（没有 float 类型），默认是 float64
```go
var af = 2.3

// 下面被注释的代码会报编译错误，小数默认类型是 float64 不能直接赋值给 float32 类型变量
//var bf float32
//bf = af
//fmt.Println(bf)

//小数默认类型是 float64 可以直接赋值给 float64 类型变量
var cf float64
cf = af
fmt.Println(cf)
```

##### 复数
Go 还支持复数类型 complex64 (32位实数+32位虚数)，默认类型是 complex128 和 complex128 （64位实数+64位虚数）。复数的形式为 RE + IMi ，其中 RE 是实数部分， IM 是虚数部分，而最后的 i 是虚数单位
```go
//声明一个复数并赋值
var ac complex128 = 1 + 5i
fmt.Printf("Value is: %v", ac)
```

#### 字符串
```go
//	Go中的字符串都是采用 UTF-8 字符集编码。字符串是用一对双引号（ "" ）或反引号（ ` ` ）括起来定义，它的类型是 string

// 字符串声明
var as string
//  声明了一个字符串变量，初始化为空字符串
var bs string = ""
//true
fmt.Println(as == bs)

var cs string = "hello go"
//下面注释的这行代码会报错，在Go中字符串是不可变的
//cs[0] = 'H'

//如果真的想要修改字符串可以将字符串 cs  转换为 []byte  类型，修改后再转换回 string  类型
//  将字符串 cs  转换为 []byte  类型
c := []byte(cs)
c[0] = 'H'
//  再转换回 string  类型
cs2 := string(c)
fmt.Printf("%s\n", cs2)

//Go中可以使用 + 操作符来连接两个字符串：
ds := "hello,"
es := " world"
fs := ds + es
fmt.Printf("%s\n", fs)

//修改字符串也可以通过把字符串进行切片操作
s := "hello"
s = "H" + s[1:]
//输出值：Hello
fmt.Printf("%s\n", s)
//可以通过 ` 来声明一个多行字符串， ` 括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。
m := `hello
world` 
fmt.Println(m)
```

#### 错误类型
Go内置有一个 error 类型，专门用来处理错误信息，Go 的 package 里面还专门有一个包 errors 来处理错误：
```go
err := errors.New("test error")
if err != nil {
    fmt.Print(err)
}
```

#### 分组声明
在Go语言中，同时声明多个常量、变量，或者导入多个包时，可采用分组的方式进行声明

```go
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
```
	

#### iota 枚举
Go里面有一个关键字 iota ，这个关键字用来声明 enum 的时候采用，它默认开始值是 0 ，每调用一次加 1

```go
const(
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w //  常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说 w = iota ，因此 w == 3 。其实上面 y 和 z 可同样不用 "= iota"
	)
fmt.Printf("x = %d, y = %d, z = %d, w = %d \n",x,y,z,w)

//  每遇到一个 const 关键字， iota 就会重置，此时 v == 0
const v = iota
fmt.Printf("v = %d",v)
```

#### array 数组 
数组的定义：var arr [n]type
 - 在 [n]type 中， n 表示数组的长度， type 表示存储元素的类型。对数组的操作和其它语言类似，都是通过 [] 来进行读取或赋值
 - 由于长度也是数组类型的一部分，因此 [3]int 与 [4]int 是不同的类型，数组也就不能改变长度。
 - 数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。如果要使用指针，就需要用到 slice 类型了

```go
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
```
	
#### slice 切片
在很多应用场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫 slice,slice 并不是真正意义上的动态数组，而是一个引用类型。 slice 总是指向一个底层 array ， slice 的声明也可以像 array 一样，只是不需要长度。

从概念上面来说 slice 像一个结构体，这个结构体包含了三个元素：
- 一个指针，指向数组中 slice 指定的开始位置
- 长度，即 slice 的长度
- 最大长度，也就是 slice 开始位置到数组的最后位置的长度

对于 slice 有几个有用的内置函数：
- len 获取 slice 的长度
- cap 获取 slice 的最大容量
- append 向 slice 里面追加一个或者多个元素，然后返回一个和 slice 一样类型的 slice
- copy 函数 copy 从源 slice 的 src 中复制元素到目标 dst ，并且返回复制的元素的个数
注： append 函数会改变 slice 所引用的数组的内容，从而影响到引用同一数组的其它 slice 。 但当 slice 中没有剩余空间（即 (cap-len) == 0 ）时，此时将动态分配新的数组空间。返回的 slice 数组指针将指向这个空间，而原
数组的内容将保持不变；其它引用此数组的 slice 则不受影响。

slice 的声明和声明 array 一样，只是少了长度

var variableName  []type
```go
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
```

#### map 
map 的格式为 map[keyType]valueType

map 的读取和设置也类似 slice 一样，通过 key 来操作，只是 slice 的 index 只能是｀int｀类型，而 map 多了很多类型，可以是 int ，可以是 string 及所有完全定义了 == 与 != 操作的类型

map过程中需要注意的几点： 
-  map 是无序的，每次打印出来的 map 都会不一样，它不能通过 index 获取，而必须通过 key 获取 
-  map 的长度是不固定的，也就是和 slice 一样，也是一种引用类型 
- 内置的 len 函数同样适用于map ，返回 map 拥有的 key 的数量 
-  map 的值可以很方便的修改，通过 numbers["one"]=11 可以很容易的把key为one 的字典值改为 11
	
```go
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
map3 := map[string]string{"key1": "value1","key2": "value2","key3": "value3"}
map4 := map3

delete(map3, "key1")
//map[key2:value2 key3:value3]
fmt.Println(map4)

map4["key2"] = "key2-key2"
//map[key2:key2-key2 key3:value3]
fmt.Println(map3)
```
#### make、new操作
	make 用于内建类型（ map 、 slice 和 channel ）的内存分配。 new 用于各种类型的内存分配。
	内建函数 new 本质上说跟其它语言中的同名函数功能一样： new(T) 分配了零值填充的 T 类型的内存空间，并且返回其
	地址，即一个 *T 类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型 T 的零值。有一点非常重要：
	new 返回指针。
	
	内建函数 make(T, args) 与 new(T) 有着不同的功能，make只能创建 slice 、 map 和 channel ，并且返回一个有初
	始值(非零)的 T 类型，而不是 *T 。本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被
	初始化。
	
	例如，一个 slice ，是一个包含指向数据（内部 array ）的指针、长度和容量的三项描述符；在这些项目被
	初始化之前， slice 为 nil 。对于 slice 、 map 和 channel 来说， make 初始化了内部的数据结构，填充适当的值。

### 流程控制
流程控制：流程控制包含分三大类：条件判断，循环控制和无条件跳转

#### if 条件判断
Go 里面 if 条件判断语句中不需要括号
```go
x := 5
if x > 10 {
    fmt.Println("x is greater than 10")
} else {
    fmt.Println("x is less than 10")
}
```

Go 的 if 还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其 他地方就不起作用了
```go
if y := 20; y > 10 {
    fmt.Println("y is greater than 10")
} else {
    fmt.Println("y is less than 10")
}
// 这个地方如果这样调用就编译出错了，因为 y 是条件里面的变量
//fmt.Println(y)
````

多个条件的时候
```go
integer := 10
if integer == 3 {
    fmt.Println("The integer is equal to 3")
} else if integer < 3 {
    fmt.Println("The integer is less than 3")
} else {
    fmt.Println("The integer is greater than 3")
}
```

#### goto 无条件跳转
```go
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
```

#### for 循环
Go 里面最强大的一个控制逻辑就是 for ，它既可以用来循环读取数据，又可以当作 while 来控制逻辑，还能迭代操作。

在循环里面有两个关键操作 break 和 continue , break 操作是跳出当前循环， continue 是跳过本次循环。当嵌套 过深的时候， break 可以配合标签使用，即跳转至标签所指定的位置。
	for expression1; expression2; expression3 {
	//...
	}

	expression1 、 expression2 和 expression3 都是表达式，其中 expression1 和 expression3 是变量声明或者
	函数调用返回值之类的， expression2 是用来条件判断， expression1 在循环开始之前调用， expression3 在每
	轮循环结束之时调用
	
```go
sum1 := 0;
for index := 0; index < 10; index++ {
    sum1 += index
}
fmt.Println("sum1 is equal to ", sum1)

//有些时候如果我们忽略 expression1 和 expression3 ：
sum2 := 1
for ; sum2 < 1000; {
    sum2 += sum2
}
fmt.Println("sum2 is equal to ", sum2)


//for 配合 range 可以用于读取 slice 和 map 的数据
// map 的初始化可以通过 key:val 的方式初始化值
map2 := map[int]string{1: "one", 2: "two", 3: "three", 4: "four"}
for k,v:=range map2 {
    fmt.Println("map's key:",k)
    fmt.Println("map's val:",v)
}
//由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用 _ 来丢弃不需要的返回值
for _, v := range map2{
    fmt.Println("map's val:", v)
}
```


#### switch
有些时候你需要写很多的 if-else 来实现一些逻辑处理，这个时候代码看上去就很丑很冗长，而且也不易于以后的维护，这个时候 switch 就能很好的解决这个问题

在 Go 中可以把很多值聚合在了一个 case 里面
```go
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
```

Go 里面 switch 默认相当于每个 case 最后带有 break ，匹配成功后不会自动向下执行其他case，而是跳出整个 switch , 但是可以使用 fallthrough 强制执行后面的case代码

```go
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
```

### 函数
函数是Go里面的核心设计，它通过关键字 func 来声明，它的格式如下：
```go
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
	// 这里是处理逻辑代码
	// 返回多个值
	return value1, value2
}
```

- 关键字 func 用来声明一个函数 funcName
- 函数可以有一个或者多个参数，每个参数后面带有类型，通过 , 分隔
- 函数可以返回多个值
- 上面返回值声明了两个变量 output1 和 output2 ，如果你不想声明也可以，直接就两个类型
- 如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
- 如果没有返回值，那么就直接省略最后的返回信息
- 如果有返回值， 那么必须在函数的外层添加return语句


#### 函数的使用
```go
a := 2
b := 3
fmt.Printf("max = %v \n", max(a, b))

add, multiplied := addAndMultiplied(3, 4)
fmt.Printf("add = %v, multiplied = %v \n", add, multiplied)

add, multiplied = addAndMultiplied(5, 6)
fmt.Printf("add = %v, multiplied = %v \n", add, multiplied)
```

````GO
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
````


#### 可变参数
Go 函数支持变参。接受变参的函数是有着不定数量的参数的.实际上这些参数是一个同类型 的 slice


````go
variableParamFunc(1, 2, 3, 4, 5, 6)
````

```go
//为了做到这点，首先需要定义函数使其接受变参：arg ...int 告诉 Go 这个函数接受不定数量的参数。实际上这些参数是一个 int 类型 的 slice
func variableParamFunc(arg ...int) {
	for index, n := range arg {
		fmt.Printf("this index is %d and the number is: %d\n", index, n)
	}
}
````

#### 传值与传指针传值与传指针
当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一份copy，当在被调用函数中修改参数值的时候，调用函数中相应实参不会发生任何变化，因为数值变化只作用在copy上。

```go
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
```

```go
// a + 1
func addOne(a int) {
	a = a + 1
}

// a + 1
func addOneWithPointer(a *int) {
	*a = * a + 1
}
```

传指针的好处
- 传指针使得多个函数能操作同一个对象。
- 传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
- Go 语言中 string ， slice ， map 这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变 slice 的长度，则仍需要取地址传递指针）

#### defer 
Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题。

#### 函数作为值、类型
在 Go 中函数也是一种变量，我们可以通过 type 来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型。

函数当做值和类型在我们写一些通用接口的时候非常有用，一个函数类型类似策略接口，可以接收不同的策略类

声明函数类型 type typeName func(input1 inputType1 [, input2 inputType2 [, ...]) (result1 resultType，result2 resultType)

```go
slice := []int{1, 2, 3, 4, 5, 7}
fmt.Println("slice = ", slice)
//  把函数当做值来传递了，只要奇数
odd := filter(slice, isOdd)
fmt.Println("Odd elements of slice are: ", odd)
//  把函数当做值来传递了，只要偶数
even := filter(slice, isEven)
fmt.Println("Even elements of slice are: ", even)
```
````go

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
````

#### Panic和Recover
Go没有像Java那样的异常机制，它不能抛出异常，而是使用了 panic 和 recover 机制。 一定要记住，你应当把它作为最后的手段来使用，也就是说，你的代码中应当没有，或者很少有 panic 的东西。

##### Panic
是一个内建函数，可以中断原有的控制流程，进入一个令人恐慌的流程中。当函数 F 调用 panic ，函数F的执行被中断，但是 F 中的延迟函数会正常执行，然后F返回到调用它的地方。在调用的地方， F 的行为就像调用了 panic 。这一
过程继续向上，直到发生 panic 的 goroutine 中所有调用的函数返回，此时程序退出。恐慌可以直接调用 panic 产生。也可以由运行时错误产生，例如访问越界的数组。

##### Recover
是一个内建的函数，可以让进入令人恐慌的流程中的 goroutine 恢复过来。 recover 仅在延迟函数中有效。在正常的执行过程中，调用 recover 会返回 nil ，并且没有其它任何效果。如果当前的 goroutine 陷入恐慌，调用
recover 可以捕获到 panic 的输入值，并且恢复正常的执行。


#### main 函数和 init 函数
Go里面有两个保留的函数： init 函数（能够应用于所有的 package ）和 main 函数（只能应用于 package main ）。这两个函数在定义时不能有任何的参数和返回值。虽然一个 package 里面可以写任意多个 init 函数，但这无论是对
于可读性还是以后的可维护性来说，我们都强烈建议用户在一个 package 中每个文件只写一个 init 函数。Go程序会自动调用 init() 和 main() ，所以你不需要在任何地方调用这两个函数。每个 package 中的 init 函数都是
可选的，但 package main 就必须包含一个 main 函数。

程序的初始化和执行都起始于 main 包。如果 main 包还导入了其它的包，那么就会在编译时将它们依次导入。有时一个包会被多个包同时导入，那么它只会被导入一次（例如很多包可能都会用到 fmt 包，但它只会被导入一次，因为没
有必要导入多次）。当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行 init 函数（如果有的话），依次类推。等所有被导入的包都加载完毕了，就会开始对 main 包中的包级常量和变量进行初始化，然后执行 main 包中的 init 函数（如果存在的话），最后执
行 main 函数。

#### import 方式
fmt 是 Go 语言的标准库，其实是去 goroot 下去加载该模块，当然Go的import还支持如下两种方式来加载自己写的模块：
1. 相对路径：import “./model” 当前文件同一目录的model目录，但是不建议这种方式来import
2. 绝对路径：import “shorturl/model” 加载gopath/src/shorturl/model模块

##### 一些特殊的import
- 1.点操作
```go
import(
. "fmt"
)
```
这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调用的fmt.Println("hello world")可以省略的写成Println("hello world")
- 2.别名操作
```go
import(
f "fmt"
)
```
别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("hello world")
- 3._操作
```go
import (
"database/sql"
_ "github.com/ziutek/mymysql/godrv"
)
```
_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数

### struct
Go语言中，也和其他语言一样，我们可以声明新的类型，作为其它类型的属性或字段的容器。

例如，我们可以创建一个自定义类型 person 代表一个人的实体。这个实体拥有属性：姓名和年龄。这样的类型我们称之 struct 。

```go
type Person struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Person
	phone string
}
````
```go
// 定义 person 类型的变量，并初始化值
var P1 Person
P1.name = "g-senlin"
P1.age = 1
fmt.Printf("The person's name is %s \n", P1.name)

//按照顺序提供初始化值
P2 := Person{"gsl", 2, "187-xxxx"}
fmt.Printf("The person's name is %s \n", P2.name)

//通过 field:value 的方式初始化，这样可以任意顺序
P3 := Person{age: 24, name: "senlin"}
fmt.Printf("The person's name is %s \n", P3.name)
```

#### struct 的匿名字段
定义一个 struct 的时候是字段名与其类型一一对应，实际上 Go 支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段。

当匿名字段是一个 struct 的时候，那么这个 struct 所拥有的全部字段都被隐式地引入了当前定义的这个 struct ,所有的内置类型和自定义类型都是可以作为匿名字段的。通过匿名字段能够实现字段的继承。

```go
S1 := Student{Person{"g-senlin", 1, "187xxxx"}, "1888888888"}
// 访问
//Go里面通过最外层的优先访问允许我们去重载通过匿名字段继承的一些字段，当然如果我们想访问重载后对应匿名类型里面的字段，可以通过匿名字段名来访问。

fmt.Printf("S1.name= %s S1.age= %d  S1.phone= %s \n", S1.name, S1.age, S1.phone)
fmt.Printf("S1.Person.name= %s S1.Person.age= %d  S1.Person.phone= %s \n", S1.Person.name, S1.Person.age, S1.Person.phone)

//	修改
S1.name = "xxxxx"
S1.Person.age=222
S1.Person.phone="19999999999999"
S1.phone="00000000"
fmt.Printf("S1.name= %s S1.age= %d  S1.phone= %s \n", S1.name, S1.age, S1.phone)
fmt.Printf("S1.Person.name= %s S1.Person.age= %d  S1.Person.phone= %s \n", S1.Person.name, S1.Person.age, S1.Person.phone)
```


### method
用Rob Pike的话来说就是： `A method is a function with an implicit first argument, called a receiver.`

method的语法如下：

`func (r ReceiverType) funcName(parameters) (results)`

```go
//这段代码可以计算出来长方形的面积，但是area()不是作为Rectangle的方法实现的（类似面向对象里面的方法），而是将Rectangle的对象（如r1,r2）作为参数传入函数计算面积的
func area_Rectangle(r Rectangle) float64 {
	return r.width * r.height
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

//	method 是附属在一个给定的类型上的，他的语法和函数的声明语法几乎一样，只是在 func 后面增加了一个receiver(也就是method所依从的主体)
// Receiver 以值传递
func (r Rectangle) area() float64 {
	return r.height * r.width
}

// Receiver 以值传递
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}


type Person struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Person
	phone string
}

func (p Person) say() {
	fmt.Printf("%s say hello \n", p.name)
}

func (p Person) run() {
	fmt.Printf("%s person run \n", p.name)
}

func (s Student) run() {
	fmt.Printf("%s Student run \n", s.name)
}
```

```go
r := Rectangle{width: 10, height: 20}
fmt.Printf("长方形面积为：%v \n", area_Rectangle(r))

r1 := Rectangle{12, 2}
r2 := Rectangle{9, 4}
c1 := Circle{10}
c2 := Circle{25}

//在使用method的时候重要注意几点
//- 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
//- method里面可以访问接收者的字段
//- 调用method通过 . 访问，就像struct里面访问字段一样
fmt.Println("Area of r1 is: ", r1.area())
fmt.Println("Area of r2 is: ", r2.area())
fmt.Println("Area of c1 is: ", c1.area())
fmt.Println("Area of c2 is: ", c2.area())
````
Receiver可以以值传递，也可以是指针, 两者的差别在于, 指针作为 Receiver 会对实例对象的内容发生操作,而普通类型作为 Receiver 仅仅是以副本作为操作对象,并不对原实例对象发生操作。

method 可以定义在任何你自定义的类型、内置类型、struct 等各种类型上面

#### 指针作为receiver

#### method继承
Go 的一个神奇之处是 method 也是可以继承的。如果匿名字段实现了一个 method ，那么包含这个匿名字段的 struct 也能调用该 method

```go
S1 := Student{Person{"g-senlin", 1, "187xxxx"}, "1888888888"}
S1.say()
```

#### method重写
````go
//调用重写 run 的方法
S1.run()
//调用匿名字段的 run 方法
S1.Person.run()
````
通过以上这些内容，我们可以设计出基本的面向对象的程序了，但是Go里面的面向对象是如此的简单，没有任何的私有、 公有关键字，通过大小写来实现(大写开头的为共有，小写开头的为私有)，方法也同样适用这个原则。

### inteface
interface 是一组 method 的组合，我们通过 interface 来定义对象的一组行为

#### interface 类型
interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。

interface 可以被任意的对象实现。同理，一个对象可以实现任意多个 interface。任意的类型都实现了空 interface (我们这样定义：interface{})，也就是包含0个 method 的 interface

```go
type Men interface {
SayHi()
Sing(lyrics string)
Guzzle(beerStein string)
}
```
#### interface 值
定义了一个 interface 的变量，那么这个变量里面可以存实现这个 interface 的任意类型的对象。

interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现， go 通过interface实现了duck-typing:即"当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那
么这只鸟就可以被称为鸭子"。

#### 空interface
空 interface(interface{}) 不包含任何的 method ，正因为如此，所有的类型都实现了空 interface 。空interface对于描述起不到任何的作用(因为它不包含任何的 method），但是空 interface 在我们需要存储任意类型的数值的时候相当
有用，因为它可以存储任意类型的数值。


#### interface 变量存储的类型

P判断变量里面实际保存了的是哪个类型的对象
- Comma-ok断言
直接判断是否是该类型的变量： value, ok = element.(T)，这里 value 就是变量的值，ok 是一个 bool 类型，element 是 interface 变量，T 是断言的类型。如果 element 里面确实存储了 T 类型的数值，那么 ok 返回 true，否则返回           false

#### 嵌入 interface
Go里面真正吸引人的是他内置的逻辑语法，就像我们在学习Struct时学习的匿名字段，多么的优雅啊，那么相同的逻辑引入到interface里面，那不是更加完美了。如果一个interface1作为interface2的一个嵌入字段，那么
interface2隐式的包含了interface1里面的method。

### reflect
Go语言实现了反射，所谓反射就是动态运行时的状态。我们一般用到的包是reflect包。

要去反射是一个类型的值(这些值都实现了空interface)，首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)


### goroutine
goroutine是Go并行设计的核心。goroutine说到底其实就是线程，但是他比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大
概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。

goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过 go 关键字实现了，其实就是一个普通的函数。

设计上我们要遵循：不要通过共享来通信，而要通过通信来共享

#### channels
goroutine 运行在相同的地址空间，因此访问共享内存必须做好同步。那么 goroutine 之间如何进行数据的通信呢，Go提供了一个很好的通信机制 channel。channel可以与 Unix shell 中的双向管道做类比：可以通过它发送或者接收
值。这些值只能是特定的类型：channel 类型。定义一个channel时，也需要定义发送到channel的值的类型。
```go
//必须使用 make 创建 channel：
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})

//channel通过操作符 <- 来接收和发送数据
//  发送 v 到 channel ch.
ch <- v 
//  从 ch 中接收数据，并赋值给 
v := <-ch 
```
```go
go say("world")
	say("hello")
```
```go
func say(s string) {
	for i := 0; i < 5; i++ {
		//runtime.Gosched() 表示让 CPU 把时间片让给别人,下次某个时候继续恢复执行该 goroutine
		runtime.Gosched()
		fmt.Printf("%d ----%s \n", i, s)
	}
}
```

默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得 Goroutines 同步变的更加的简单，而不需要显式的 lock 。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其
次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲 channel 是在多个goroutine之间同步很棒的工具。
```go
a := []int{7, 2, 8, -9, 4, 0, 10}
	c := make(chan int)
	go sum(c, a)
	go sum(c, a[:len(a)/2])
	x, y := <-c, <-c

	fmt.Printf("x= %d \n", x)
	fmt.Printf("y= %d \n", y)
```
```go
func sum(c chan int, a []int) {
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	c <- sum
}
```

#### Buffered Channels
```go
ch := make(chan type, value)
// 无缓冲（阻塞）
value == 0 ! 
//缓冲（非阻塞，直到 value  个元素）	
value > 0 !  
	
// channel 的缓冲区大小为 2
c1 := make(chan int, 2)
c1 <- 1
c1 <- 2
fmt.Println(<-c1)
fmt.Println(<-c1)

c2 := make(chan int, 1)
c2 <- 11
// channel 的缓存大小为 一个 int ，则执行下一行代码往缓冲区写数据将会阻塞，直到缓存取的数据被读取了，才继续写数据
// 由于此处是单线程的所以被阻塞后就不往下执行了， fatal error: all goroutines are asleep - deadlock!
//c2 <- 22
fmt.Println(<-c2)
//fmt.Println(<-c2)
```

#### Range和Close
可以通过 range，像操作 slice 或者map一样操作缓存类型的 channel

`for i := range c `能够不断的读取channel里面的数据，直到该channel被显式的关闭。可以显 式的关闭channel，生产者通过关键字 close 函数关闭 channel。关闭 channel 之后就无法再发送任何数据了，
在消费方可以通过语法 v, ok := <-ch 测试 channel 是否被关闭。如果 ok 返回 false ，那么说明 channel 已经没有任何数据 并且已经被关闭

应该在生产者的地方关闭 channel，而不是消费的地方去关闭它，这样容易引起 panic ,channel 不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束 range 循环之类的
```go
c3 := make(chan int, 10)
go func() {
    for i := 0; i < cap(c3); i++ {
        c3 <- i
    }
    close(c3)
}()
value, ok := <-c3
//0 true ,ok 为 true 代表已经关闭
fmt.Println(value, ok)

for a := range c3 {
    fmt.Println(a)
}
value, ok = <-c3
//0 false ,ok 为false 代表已经关闭
fmt.Println(value, ok)
```

#### Select
Go 里面提供了一个关键字 select ，通过 select 可以监听channel上的数据流动

select 默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个 channel 都准备好的时候，select 是随机的选择一个执行的

```go
c4 := make(chan int, 10)
quit := make(chan int)
go func() {
    for i := 0; i < cap(c4); i++ {
        fmt.Println(<-c4)
    }
    quit <- 0
}()

func() {
    x := 1
    for {
        select {
        case c4 <- x:
            x += 1
        case <-quit:
            fmt.Println("exit ...")
            return

            //有时候会出现goroutine阻塞的情况，为了避免整个的程序进入阻塞，我们可以利用select来设置超时
            //case <-time.After(5 * time.Second):
            //	fmt.Println("time out")
            //	break

        //	在 select 里面还有 default 语法， select 其实就是类似 switch 的功能，default 就是当监听的 channel 都没有准备好的时候，
        // 默认执行的（select 不再阻塞等待 channel）
        default:
            fmt.Println("execute select default")
        }
    }
    fmt.Println("----------------")
}()
```

### runtime goroutine
runtime包中有几个处理goroutine的函数：
- Goexit 退出当前执行的goroutine，但是defer函数还会继续调用
- Gosched 让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
- NumCPU 返回 CPU 核数量
- NumGoroutine 返回正在执?行和排队的任务总数
- GOMAXPROCS 用来设置可以运行的CPU核数	

### Go 语言的二十五个关键字
|     |      |    |  |  |
| -------- | ----------- | ------ | --------- | :----- |
| break    | default     | func   | interface | struct |
| case     | defer       | go     | map       | select |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

### 特别说明
1、基础类型底层都是分配了一块内存，然后存储了相应的值.
2、Go之所以会那么简洁，是因为它有一些默认的行为：
	2.1、大写字母开头的变量是可导出的，也就是其它包可以读取的，是公用变量；小写字母开头的就是不可导出的，是私有变量
	2.2、大写字母开头的函数也是一样，相当于 class中的带 public 关键词的公有函数；小写字母开头的就是有 private 关键词的私有函数。