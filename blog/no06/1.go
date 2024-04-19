package main

import "fmt"

func main() {
	no06_2()
}
func no06_1() {
	var a int = 1
	for i := 1; i < 3; i++ {
		fmt.Println("p:", i)
	}
	fmt.Println("----")
	for a < 3 {
		fmt.Println("p:", a)
		a++
	}
}
func no06_2() {
	for i := 1; i < 4; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("p:", i)
	}
	fmt.Println("----")
	for i := 1; i < 4; i++ {
		if i == 2 {
			break
		}
		fmt.Println("p:", i)
	}
	fmt.Println("----")
	var a int = 1
test: //從這裡重新開始
	for a < 4 {
		if a == 2 {
			a++
			goto test //會返到test
		}
		fmt.Println("p:", a)
		a++
	}
}
