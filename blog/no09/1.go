package main

import "fmt"

func main() {
	no09_1()
}

type no09 struct {
	i int
	n string
}
type no092 struct {
	a struct {
		b int
		c string
	}
	d bool
}
type no093 struct {
	no no09
	d  bool
}
type no094 struct {
	no09 //隱藏
	d    bool
}

func no09_1() {
	a := no09{i: 1, n: "A"}
	var b no09 = no09{2, "B"}
	fmt.Println("p:", a)
	fmt.Println("p:", a.i)
	fmt.Println("-----")
	fmt.Println("p:", b)
	fmt.Println("p:", b.i)
	fmt.Println("-----")
	c := no09_1r(a)
	fmt.Println("p:", c.i)
	fmt.Println("p:", c.n)
	fmt.Println("-----")
	d := no092{}
	d.a.b = 3
	d.a.c = "C"
	d.d = true
	fmt.Println("p:", d.a)
	fmt.Println("p:", d.a.b)
	fmt.Println("p:", d.a.c)
	fmt.Println("p:", d.d)
	fmt.Println("-----")
	e := no093{no09{4, "D"}, false}
	fmt.Println("p:", e.no)
	fmt.Println("p:", e.no.i)
	fmt.Println("p:", e.no.n)
	fmt.Println("p:", e.d)
	fmt.Println("-----")
	f := no094{no09{4, "D"}, false}
	fmt.Println("p:", f)
	fmt.Println("p:", f.i)
	fmt.Println("p:", f.n)
	fmt.Println("p:", f.d)
}
func no09_1r(no no09) no09 {
	no.i += 2
	no.n = "c"
	return no
}
