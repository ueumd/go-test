package main

import "fmt"

// 接口
type Programmer interface {
	// 方法只声明
	Coding() string
	Debug() string
}

type Golanger struct {
	Name string
}


func (g Golanger) Coding() string  {
	fmt.Println("golang 开发者")
	return  "golang 开发者"
}

func (g Golanger) Debug() string  {
	fmt.Println("golang 开发者")
	return  "Debug"
}


type FrontEngineer struct {
	Name string
}


func (g FrontEngineer) Coding() string  {
	fmt.Println("前端 开发者")
	return  "前端 开发者"
}

func (g FrontEngineer) Debug() string  {
	fmt.Println("前端 开发者")
	return  "Debug"
}

/**
Java里一种类型继承了一个接口，就必须实现该接口的全部方法

对于Golanger 结构休来说，实现任何方法都可以， 但是只要不全部实现 Coding Debug方法，那个Golanger 就不是一个Progammer类型
1. Golanger 本身就是就一个类型，那何必在意是不是Programer类型？
2. 答案：封装 继承 多态!!! (多态)
 */

// 开发电商网站 支付环节 使用 微信 支付宝 银行卡 系统支持各种类型的支付， 每一种支付都有统一接口
// 定义一个协议 1 创建订单 2支付 3查询支付状态 4退款

type AliPay struct {

}

type WeChat struct {

}
type Bank struct {

}

var a AliPay
var b Bank
var w WeChat

// 如果后期接入一种新的支付类型支付 或者 取消已有的支付
// 但go 语言中不支持继承

// 那这么这时需要一个通用类型 接口的强制性 满足多态特性

//var x Tongyong
//x = Bank{}
//x = AliPay{}
//x = WeChat{}
//
//x.pay
//x.create
//x.query

// 完成多态
func handler(p Programmer)  {
	p.Coding()
}


func main()  {
	// 声明一个Programmer接口类型变量 pro

	var goer = &Golanger{"Tom"}
	goer.Coding()
	handler(goer)

	var jser = FrontEngineer{"Jerry"}
	jser.Coding()
	handler(jser)
}