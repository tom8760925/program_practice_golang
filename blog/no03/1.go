package main

import "fmt"

func main() {
	no03_2()
}
func no03_1() {
	var a int = 1
	fmt.Print("p:", a, "\n") //單純顯示
	fmt.Printf("p:%d\n", a)  //使用參數
	fmt.Println("p:", a)     //自動換行
}
func no03_2() {
	var a, b int = 0, 0
	fmt.Print("s:")
	fmt.Scan(&a, &b) //可以連續讀取，並可以用空白符號間隔兩個變數的輸入
	fmt.Println("p:", a)
	fmt.Println("p:", b)
	fmt.Println("-------")
	fmt.Print("s:")
	fmt.Scanf("%d-%d", &a, &b) //可以使用參數，並可以自行設定符號作為間隔
	fmt.Println("p:", a)
	fmt.Println("p:", b)
	fmt.Println("-------")
	fmt.Print("s:")
	fmt.Scanln(&a, &b) //只要到換行符號就停止
	fmt.Println("p:", a)
	fmt.Println("p:", b)
}
