package main

import (
	"fmt"
	"math"
)

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

//这段代码可以计算出来长方形的面积，但是area()不是作为Rectangle的方法实现的（类似面向对象里面的方法），而是将Rectangle的对象（如r1,r2）作为参数传入函数计算面积的
func area_Rectangle(r Rectangle) float64 {
	return r.width * r.height
}

type Person struct {
	name  string
	age   int
	phone string
}

func (p Person) say() {
	fmt.Printf("%s say hello \n", p.name)
}

func (p Person) run() {
	fmt.Printf("%s person run \n", p.name)
}

type Student struct {
	Person
	phone string
}

func (s Student) run() {
	fmt.Printf("%s Student run \n", s.name)
}

func main() {
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

	//	Receiver可以以值传递，也可以是指针, 两者的差别在于, 指针作为 Receiver 会对实例对象的内容发生操作,而普通类型作为 Receiver 仅仅是以副本作为操作对象,并不对原实例对象发生操作。
	// method 可以定义在任何你自定义的类型、内置类型、struct 等各种类型上面

	//指针作为receiver

	//method继承
	//Go 的一个神奇之处是 method 也是可以继承的。如果匿名字段实现了一个 method ，那么包含这个匿名字段的 struct 也能调用该 method
	S1 := Student{Person{"g-senlin", 1, "187xxxx"}, "1888888888"}
	S1.say()

	//method重写
	//调用重写 run 的方法
	S1.run()
	//调用匿名字段的 run 方法
	S1.Person.run()

	//	过这些内容，我们可以设计出基本的面向对象的程序了，但是Go里面的面向对象是如此的简单，没有任何的私有、 公有关键字，通过大小写来实现(大写开头的为共有，小写开头的为私有)，方法也同样适用这个原则。
}
