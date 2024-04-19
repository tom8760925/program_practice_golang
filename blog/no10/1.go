package main

import (
	"errors"
	"fmt"
)

type a struct {
	n int
	i int
}

func main() {
	no10_3()
}
func no10_1() {
	err := errors.New("p:err")
	if err != nil {
		fmt.Println(err)
	}
}
func no10_2() {
	if n, err := no10_2r(0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("p:", n)
	}
	if n, err := no10_2r(1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("p:", n)
	}
}
func no10_2r(n int) (int, error) {
	if n == 0 {
		return n, fmt.Errorf("p:err")
	}
	return n, nil
}
func no10_3() {
	if n, err := no10_3r(1, 2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("p:", n)
	}
	if n, err := no10_3r(1, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("p:", n)
	}
}
func no10_3r(n, i int) (int, error) {
	if n <= 0 || i <= 0 {
		return 0, &a{n: n, i: i}
	}
	return n + i, nil
}
func (e *a) Error() string {
	if e.n <= 0 {
		return "err(n=" + fmt.Sprint(e.n) + ")"
	} else if e.i <= 0 {
		return "err(i=" + fmt.Sprint(e.i) + ")"
	} else {
		return "p:err"
	}
}
