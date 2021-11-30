package main

import (
	"fmt"
	"sync"
	"time"
)

func p()  {
	for {
		fmt.Println("I am p method")
	}
}

func main1()  {
	//  子协程
	go p()

	// 主协程

	// 主死随从
	fmt.Println("I am main method")

	time.Sleep(time.Second * 2)
}

func main2()  {

	for i:=0; i< 100; i++ {
		// 匿名
		go func() {
			for {
				fmt.Println("I am p method",  i)
				time.Sleep(time.Second)
			}
		}()
	}


	time.Sleep(time.Second * 2)
}


/**
解决主协程的goroutine 在子协程结束后自动结束

WaitGroup 3个函数
Add()
Done()

Wait()

Add 的数量 和 Done的数量必须相等
 */

var wg sync.WaitGroup

func print(n int)  {
	defer wg.Done() // 减1
	fmt.Println(n)
}

func main()  {

	wg.Add(2) // 总共有多少个

	for i:=0; i< 5; i++ {
		go print(i)
	}

	// 阻塞主协程
	wg.Wait()
}