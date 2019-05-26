package main

import "fmt"

func main0201() {
	//变量的定义和使用
	//变量定义格式
	//var 变量名 数据类型（声明） 默认值为0
	//var 变量名 数据类型 = 值 （定义）

	//int 表示的是整型 可以存储整数
	//var a int = 100
	//fmt.Println(a)
	//a = 123
	//fmt.Println(a)

	//声明
	var a int
	a = 10
	fmt.Println(a)
}
func main0202() {
	//变量计算
	var a int = 10

	a = a + 10
	fmt.Println(a)
}
