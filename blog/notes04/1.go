package main

import (
	"fmt"
)

func main() {
	a, b, _ := notes04_1() //忽略回傳值
	fmt.Println("p:", a)
	fmt.Println("p:", b)
	fmt.Println("------")
	var c = [...]int{1, 2} //忽略大小
	fmt.Println("p:", c)
	fmt.Println("------")
	notes04_2(1, 2)
	notes04_2(1, 2, 3, 4)
}
func notes04_1() (int, string, bool) {
	return 1, "A", false
}
func notes04_2(n ...int) { //可變的input
	fmt.Println("p:", n)
}
