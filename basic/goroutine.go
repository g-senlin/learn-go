package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		//runtime.Gosched() 表示让 CPU 把时间片让给别人,下次某个时候继续恢复执行该 goroutine
		runtime.Gosched()
		fmt.Printf("%d ----%s \n", i, s)
	}
}

func sum(c chan int, a []int) {
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	c <- sum
}
func main() {
	go say("world")
	say("hello")

	//	channel
	a := []int{7, 2, 8, -9, 4, 0, 10}
	c := make(chan int)
	go sum(c, a)
	go sum(c, a[:len(a)/2])
	x, y := <-c, <-c

	fmt.Printf("x= %d \n", x)
	fmt.Printf("y= %d \n", y)

	//	默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的
	//简单，而不需要显式的lock。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其
	//次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具。

	//	Buffered Channels
	//ch := make(chan type, value)
	//value == 0 !  无缓冲（阻塞）
	//value > 0 !  缓冲（非阻塞，直到 value  个元素）

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

	//	Range和Close
	//读取两次c，这样不是很方便，Go 考虑到了这一点，所以也可以通过 range，像操作 slice 或者map一样操作缓存类型的 channel
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

	//	for i := range c 能够不断的读取channel里面的数据，直到该channel被显式的关闭。
	// 上面代码我们看到可以显 式的关闭channel，生产者通过关键字 close 函数关闭channel。关闭channel之后就无法再发送任何数据了，
	// 在消费 方可以通过语法 v, ok := <-ch 测试channel是否被关闭。如果ok返回false，那么说明channel已经没有任何数据 并且已经被关闭

	// 应该在生产者的地方关闭 channel，而不是消费的地方去关闭它，这样容易引起 panic
	// channel 不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束 range 循环之类的

	//	Select
	// Go 里面提供了一个关键字 select ，通过 select 可以监听channel上的数据流动

	//	select 默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个 channel 都准备好的时候，select 是随机的选择一个执行的

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

}
