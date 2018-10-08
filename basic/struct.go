package main

import "fmt"

func main() {

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

	//struct 的匿名字段
	//定义一个 struct 的时候是字段名与其类型一一对应，实际上 Go 支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段
	//	当匿名字段是一个 struct 的时候，那么这个 struct 所拥有的全部字段都被隐式地引入了当前定义的这个 struct,所有的内置类型和自定义类型都是可以作为匿名 字段的。

	//匿名字段就是这样，能够实现字段的继承。

	S1 := Student{Person{"g-senlin", 1, "187xxxx"}, "1888888888"}
	// 访问
	//Go里面通过最外层的优先访问允许我们去重载通过匿名字段继承的一些字段，当然如果我们想访问重载后对应匿名类型里面的字段，可以通过匿名字段名来访问。

	fmt.Printf("S1.name= %s S1.age= %d  S1.phone= %s \n", S1.name, S1.age, S1.phone)
	fmt.Printf("S1.Person.name= %s S1.Person.age= %d  S1.Person.phone= %s \n", S1.Person.name, S1.Person.age, S1.Person.phone)

	//	修改
	S1.name = "xxxxx"
	S1.Person.age = 222
	S1.Person.phone = "19999999999999"
	S1.phone = "00000000"
	fmt.Printf("S1.name= %s S1.age= %d  S1.phone= %s \n", S1.name, S1.age, S1.phone)
	fmt.Printf("S1.Person.name= %s S1.Person.age= %d  S1.Person.phone= %s \n", S1.Person.name, S1.Person.age, S1.Person.phone)

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
