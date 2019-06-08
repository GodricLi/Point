package main

import (
	"fmt"
)

func test()  {
	var a int
	a=100
	fmt.Println(a,&a)

	//指针类型，指向的是内存地址
	var b *int
	fmt.Println(b,&b)
	b=&a  //赋值内存地址
	fmt.Println(b,&b,*b)
	//修改指针的值会修改内存地址，a的值也会被修改
	*b=200
	fmt.Println(a)
}

func modify(a*int,b*[3]int)  {
	//指针传参
	*a=66
	(*b)[0]=88//修改指针类型的数字使用的方式
	// fmt.Println(a,b)
}



func main()  {
	// test()
	a:=100
	b:=&a
	arry:=[3]int{1,2,3}
	point_arry:=&arry
	modify(b,point_arry)
	fmt.Println(a,arry)
	// 指针传参在函数内被修改后，被传入的数据也会被修改
}