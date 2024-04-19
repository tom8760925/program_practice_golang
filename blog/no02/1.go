package main

import "fmt"

func main() {
	no02_3()
}
func no02_1() {
	var a = 1
	var a1 int = 2
	var (
		a2 int    = 3
		a3 string = "4"
		a4 bool
	)
	var a5, a6 int = 5, 6
	a7 := 7 //簡化寫法
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(a6)
	fmt.Println(a7)
	a5, a6 = a6, a5 //交換
	fmt.Println(a5)
	fmt.Println(a6)
}
func no02_2() {
	const a = 1
	const a1 int = 2
	const (
		a2 int    = 3
		a3 string = "4"
	)
	const a4, a5 int = 5, 6
	const ( //沒有宣告值會使用前一個值
		a6 = 6
		a7
	)
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(a6)
	fmt.Println(a7)
}
func no02_3() {
	const ( //+1
		a = iota
		a1
		a2
	)
	const ( //二進位往左移
		a3 = 1 << iota //1
		a4             //10
		a5             //100
	)
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
}
