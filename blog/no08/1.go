package main

import "fmt"

func main() {
	no08_1()
}
func no08_1() {
	var a int = 1
	var b *int
	var c **int
	var d [1]*int
	b = &a
	c = &b
	d[0] = &a
	fmt.Println("p:", b)
	fmt.Println("p:", *b)
	fmt.Println("-----")
	fmt.Println("p:", *c)
	fmt.Println("p:", **c)
	fmt.Println("-----")
	fmt.Println("p:", d[0])
	fmt.Println("p:", *d[0])
}
