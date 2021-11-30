package main

import (
	"fmt"
	"time"
)

func init()  {

	/**
	select
	作用于channel之上， 多路复用
	select 会随机公平的选择一个case语句执行
	 */

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)

	ch1 <- 1
	ch2 <- 2

	// 下面会随机执行一个case
	select {
	case data:= <-ch1 :
		fmt.Println(data)
	case data:= <-ch2 :
		fmt.Println(data)

	}
}

// select的作用 超时处理

func main()  {
	timeout := false

	go func() {
		time.Sleep(time.Second*2)
		timeout = true
	}()

	for {
		if timeout {
			fmt.Println("Over")
			break
		}

		time.Sleep(time.Millisecond*10)
	}
}