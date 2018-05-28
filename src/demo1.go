package main

import "fmt"

func main() {

	//testArray()
	testPoint()
}

func testPoint() {
	//var a int = 20 /* 声明实际变量 */
	//var ip *int    /* 声明指针变量 */
	//
	//ip = &a /* 指针变量的存储地址 */
	//
	//fmt.Printf("a 变量的地址是: %x\n", &a)
	//
	///* 指针变量的存储地址 */
	//fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	//
	///* 使用指针访问值 */
	//fmt.Printf("*ip 变量的值: %d\n", *ip)

	//testPointArray()
	//testPointToPoint()

	a := 1
	b := 2

	fmt.Printf("交换前 a = %d, b = %d \n", a, b)
	swap(&a, &b)
	fmt.Printf("交换后 a = %d, b = %d \n", a, b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}

func testPointToPoint() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}

const MAX int = 3

func testPointArray() {
	//a := []int{10, 100, 200}
	//var i int
	//
	//for i = 0; i < MAX; i++ {
	//	fmt.Printf("a[%d] = %d\n", i, a[i])
	//}

	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

func testArray() {
	var n [10]int
	/* n 是一个长度为 10 的数组 */
	var i, j int
	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}
	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}
