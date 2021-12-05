package main

import (
	"fmt"
	"time"
)

func main9()  {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)

	//ch1 <- 1
	//ch2 <- 2

	// 下面会随机执行一个case
	select {
	case data:= <-ch1 :
		fmt.Println("random ", data)
	case data:= <-ch2 :
		fmt.Println("random ", data)
	default:
		fmt.Println("exit")
	}
}


// select的作用 超时处理

func main22()  {
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

func main3() {
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(2 * time.Second) //2秒
		timeout <- true
	}()

	ch := make(chan int)
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("timeout 01")   // 超时2秒打印
	case <-time.After(time.Second * 1): // 超时1秒打印
		fmt.Println("timeout 02")
	}
}

// 检查通道是否已满
func main7() {
	ch := make(chan int, 2)
	ch <- 1

	select {
	case ch <- 2:
		fmt.Println("channel value is", <-ch)
		fmt.Println("channel value is", <-ch)
	default:
		fmt.Println("channel blocking")
	}
}

// 多个select
func main() {
	i := 0
	ch := make(chan string, 0)

	defer func() {
		// 关闭通道
		close(ch)
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Unix())
			i++

			select {
			case m := <-ch:
				println(m)
				break
			default:
			}
		}
	}()

	time.Sleep(time.Second * 4)
	ch <- "stop"
}