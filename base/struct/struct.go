package main

import "fmt"

type person struct {
	name string
	age int
}

/**
 创建 结构体变量 和 访问 结构体字段
 */

func init()  {
	// 方式1 直接声明一个空结构体
	var p1 person
	fmt.Println(p1) // { 0}

	p1.name = "Tom"

	// 方式2 直接声明一个空结构体
	var p2 person = person{}
	fmt.Println(p2) // { 0}

	// 方式3 直接声明并初始化
	var p3 person = person{"Tom", 18}
	fmt.Println(p3) 		// {Tom 18}
	fmt.Println(p3.name)	// Tom

	// 方式4  结构体指针 声明并初始化
	var p4 *person = &person{"Tom", 18}
	fmt.Println(p4) 		// &{Tom 18}

	/**
	结构体指针访问字段的标准方式应该是：(*结构体指针).字段名 (*p4).name
	Go提供了 隐式解引用 特性，也支持 结构体指针.字段名 p4.name
	 */
	fmt.Println((*p4).name) // Tom
	fmt.Println(p4.name)	// Tom 隐式解引用


	// 方式5  结构体指针 空结构体 方式4一样
	var p5 *person = new(person)
	fmt.Println(p5) 	// &{ 0}
}

// 结构体指针及赋值
func main()  {

	// 结构体指针
	p1 := &person{name: "Tom", age: 18}

	// 普通结构体
	p2 := person{name: "Jerry", age: 18}

	fmt.Println(p1.name, p2.name) // Tom Jerry

	pp1 := p1
	pp2 := p2

	fmt.Println("p1.name=", p1.name, " pp1.name=", pp1.name) // p1.name= Tom    pp1.name= Tom
	fmt.Println("p2.name=", p2.name, " pp2.name=", pp2.name) // p2.name= Jerry  pp2.name= Jerry

	pp1.name = "Jerry"

	pp2.name = "Tom"


	// p1, pp1 共同指向一个结构体地址  修改pp1 会同时改动p1
	fmt.Println("p1.name=", p1.name, " pp1.name=", pp1.name) // p1.name= Jerry  pp1.name= Jerry

	// p2 pp2 都是普通结构体，不是指向同一个结构体地址，修改pp2 不会改p2的值
	fmt.Println("p2.name=", p2.name, " pp2.name=", pp2.name) // p2.name= Jerry  pp2.name= Tom
}