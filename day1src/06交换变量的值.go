package main

import "fmt"

func main() {
	a := 10
	b := 20
	//交换变量的值

	//使用第三方变量 进行交换
	//temp := a
	//a = b
	//b = temp

	//通过计算 进行交换
	//a = a + b //30
	//b = a - b //10
	//a = a - b //20

	//通过位运算 进行交换(不需要学习)
	//a = a ^ b
	//b = a ^ b
	//a = a ^ b

	//通过多重赋值 进行交换
	a, b = b, a

	fmt.Println(a) //20
	fmt.Println(b) //10
}
