package main

import "fmt"


/**
指针
空指针： 没有初始化的指针。var p *int
野指针： 被一片无效的地址空间初始化
 */

func p1()  {
	// 空指针
	var p *int;

	fmt.Println(*p) // invalid memory address or nil pointer dereference 无效的内存地址 或空指针引用
}

func p2()  {
	var p int = 10
	fmt.Println("&p = ", &p) // &p =  0xc00000a098

	// 野指针
	//var p2 *int = 0xc00000a098
	//fmt.Println(*p)
}

func main()  {

	// heap上申请一片内存空间

	var p *string

	p = new(string)

	*p = "Abcdef"

	fmt.Printf("%s\n", *p)
	fmt.Printf("%q\n", *p)
}