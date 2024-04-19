package main

import "fmt"

func main() {
	no05_1()
}
func no05_1() {
	var a int = 1
	var b string = "Apple"
	if a == 1 {
		fmt.Println("p:", a)
	} else if a == 2 {
		fmt.Println("p:", a)
	} else {
		fmt.Println("p:", a)
	}
	fmt.Println("-----")
	switch b {
	case "Apple":
		fmt.Println("p:", b)
	case "banana":
		fmt.Println("p:", b)
	default:
		fmt.Println("p:", b)
	}
}
