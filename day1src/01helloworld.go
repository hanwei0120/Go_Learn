package main //文件所属包 主函数默认所在的包是main

//行注释  可以注释一行
/*
块注释
可以注释多行
注释的作用是为代码进行提示 注释不参与程序的编译
 */

import "fmt"//format 包 包含格式化输入输出 goland 会自动导入需要的包


//func function  函数格式
//main 程序的主入口  有且只有一个
//() 函数参数  如果没有为空
//{} 代码体  函数体
func main(){
	//包名.函数名()
	//打印输出  一个字符串 ""引起来成为字符串 （字符串是一种数据类型）
	fmt.Println("hello world")
	fmt.Println("欢迎来到传智播客，黑马程序员，学习go语言与区块链")
}