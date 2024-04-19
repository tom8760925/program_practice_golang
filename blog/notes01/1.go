package main

import "fmt"

func main() {
	a := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	b := make(map[string]int)
	c := a
	fmt.Println("p:", a)
	fmt.Println("p:", a["A"])
	fmt.Println("-----")
	fmt.Println("p:", b)
	b["A"] = 1 //新增
	fmt.Println("p:", b)
	fmt.Println("-----")
	fmt.Println("p:", c)
	fmt.Println("-----")
	fmt.Println("p:", len(a))
	fmt.Println("-----")
	delete(a, "C") //刪除
	fmt.Println("p:", a)
}
