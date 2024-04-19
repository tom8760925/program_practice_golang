package main

import "fmt"

type a interface {
	is()
}
type b interface {
	isi(int)
}
type c interface {
	isir(int) string
}
type d interface {
	is2()
}
type e interface {
	is()
	isi(int)
	isir(int) string
	is2()
}
type nint struct {
	n int
}

type nstring struct {
	n string
}

func (n nint) is() {
	fmt.Println("P:", n.n, "(int)")
}

func (n nstring) is() {
	fmt.Println("P:", n.n, "(string)")
}
func (n nint) isi(i int) {
	fmt.Println("P:", n.n, "(int)", i)
}

func (n nstring) isi(i int) {
	fmt.Println("P:", n.n, "(string)", i)
}
func (n nint) isir(i int) string {
	return "P:" + fmt.Sprint(n.n) + "(int)" + fmt.Sprint(i)
}

func (n nstring) isir(i int) string {
	return "P:" + n.n + "(string)" + fmt.Sprint(i)
}
func (n *nint) is2() {
	fmt.Println("P:", n.n, "(int)")
}

func (n *nstring) is2() {
	fmt.Println("P:", n.n, "(string)")
}
func is(n e) {
	n.is()
	n.is2()
}
func is2(n e, i int) string {
	return n.isir(i)
}
func main() {
	var aa a
	var bb b
	var cc c
	var dd d
	aa = nint{n: 1}
	aa.is()
	fmt.Println("------")
	aa = nstring{n: "A"}
	aa.is()
	fmt.Println("======")
	bb = nint{n: 2}
	bb.isi(1)
	fmt.Println("------")
	bb = nstring{n: "B"}
	bb.isi(1)
	fmt.Println("======")
	cc = nint{n: 3}
	fmt.Println(cc.isir(1))
	fmt.Println("------")
	cc = nstring{n: "C"}
	fmt.Println(cc.isir(1))
	fmt.Println("=====")
	dd = &nint{n: 4}
	dd.is2()
	fmt.Println("------")
	dd = &nstring{n: "D"}
	dd.is2()
	fmt.Println("======")
	ee := nint{n: 5}
	ee.is()
	fmt.Println("------")
	ee2 := nstring{n: "E"}
	ee2.is()
	fmt.Println("=====")
	ff := nint{n: 6}
	is(&ff)
	fmt.Println("------")
	ff2 := nstring{n: "F"}
	is(&ff2)
	fmt.Println("=====")
	gg := nint{n: 7}
	fmt.Println(is2(&gg, 1))
	fmt.Println("------")
	gg2 := nstring{n: "G"}
	fmt.Println(is2(&gg2, 1))
}
