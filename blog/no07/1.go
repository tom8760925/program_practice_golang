package main

import "fmt"

func main() {
	no07_5()
}
func no07_1() {
	var a [3]int = [3]int{1, 2, 3}
	var b []int = []int{1, 2, 3, 4, 5, 6}
	var c []int = make([]int, 5)     //宣告5個長度
	var d []int = make([]int, 5, 10) //宣告5個長度，預留10個長度
	var e []int
	f := [3]int{1, 2, 3}
	var g []int = a[0:3]
	fmt.Println("p:", a[0])
	fmt.Println("-----")
	fmt.Println("p:", b)
	fmt.Println("-----")
	fmt.Println("p:", c)
	fmt.Println("-----")
	fmt.Println("p:", d)
	fmt.Println("-----")
	fmt.Println("p:", e)
	fmt.Println("-----")
	fmt.Println("p:", f)
	fmt.Println("-----")
	fmt.Println("p:", g)
}
func no07_2() {
	var a [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var b [][]int = [][]int{{1, 2, 3}, {4, 5, 6}}
	var c [][]int = [][]int{make([]int, 3), make([]int, 3)}
	var d [][]int = make([][]int, 2)
	for i := range d {
		d[i] = make([]int, 3)
	}
	var e [][]int
	f := [2][3]int{}
	fmt.Println("p:", a)
	fmt.Println("-----")
	fmt.Println("p:", b)
	fmt.Println("-----")
	fmt.Println("p:", c)
	fmt.Println("-----")
	fmt.Println("p:", d)
	fmt.Println("-----")
	fmt.Println("p:", e)
	fmt.Println("-----")
	fmt.Println("p:", f)
}
func no07_3() {
	var a []int = []int{1, 2, 3}
	var b []int
	var c [][]int = [][]int{{1, 2, 3}, {4, 5, 6}}
	var d []int = make([]int, 3)
	b = append(b, 1, 2, 3, 4)
	c[0] = append(c[0], 4)
	copy(d, a) //複製進去的容量要足夠才能複製
	fmt.Println("p:", a)
	fmt.Println("-----")
	fmt.Println("p:", b)
	fmt.Println("-----")
	fmt.Println("p:", c)
}
func no07_4() {
	var a []int = []int{1, 2, 3}
	var b [][]int = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("p:", len(a))
	fmt.Println("-----")
	fmt.Println("p:", len(b))
	fmt.Println("-----")
	fmt.Println("p:", len(b[0]))
}
func no07_5() {
	var a []int = []int{1, 2, 3}
	for i := range a { //顯示key
		fmt.Println("p:", i)
	}
	fmt.Println("-----")
	for _, i := range a { //顯示值
		fmt.Println("p:", i)
	}
}
