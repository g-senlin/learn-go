package main

import (
	"fmt"
	"time"
)

func main() {
	c4 := make(chan int, 10)
	quit := make(chan int)
	go func() {
		for i := 0; i < cap(c4); i++ {
			//fmt.Println(<-c4)
			c4 <- 1
		}
		c4 <- 1
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
			case <-time.After(5 * time.Second):
				fmt.Println("time out")
				break
				//default:
				//	fmt.Println("execute select default")
			}
		}
		fmt.Println("----------------")
	}()
}
