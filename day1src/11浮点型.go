package main

import "fmt"

func main() {
	//单精度浮点型  小数位数有效为7位
	var a float32 = 3.14
	//双精度浮点型  小数位数有效为15位
	var b float64 = 3.14
	//建议在开发中使用float64代替float32

	//自动推导类型创建浮点型 默认为float64
	c := 3.14//float64

	//fmt.Println(a)
	//fmt.Println(b)

	fmt.Printf("%T\n", c)

	fmt.Printf("%.30f\n", a)
	fmt.Printf("%.30f\n", b)
}
