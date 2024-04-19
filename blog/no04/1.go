package main

import "fmt"

func main() {
	no04_1()
}
func no04_1() {
	a := no04_1r
	fmt.Println("p:", no04_1r(2, "A"))
	fmt.Println("------")
	fmt.Println("p:", a(2, "B"))
}
func no04_1r(n int, i string) int {
	fmt.Println("p:", i)
	fmt.Println("p:", n)
	fmt.Println("------")
	return n
}
func no04_2() {
	a, b := no04_2r(1, "A")
	fmt.Println("p:", a)
	fmt.Println("p:", b)
	fmt.Println("------")
	var c int
	var d string
	c, d = no04_2r2(2, "B")
	fmt.Println("p:", c)
	fmt.Println("p:", d)
}
func no04_2r(n int, i string) (int, string) {
	return n, i
}
func no04_2r2(n int, i string) (rn int, ri string) {
	rn = n
	ri = i
	return
}
