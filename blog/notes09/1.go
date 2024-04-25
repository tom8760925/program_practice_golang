package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	getmain()
}
func getmain() {
	url := "http://localhost:4567/test"
	req, _ := http.Get(url)
	defer req.Body.Close() //defer 延後執行在結束時執行
	s, _ := io.ReadAll(req.Body)
	fmt.Println(string(s))
}
func postmain() {
	v := url.Values{"v": {"d"}, "i": {"4"}, "ii": {"4"}}
	url := "http://localhost:4567/test"
	req, _ := http.PostForm(url, v)
	defer req.Body.Close()
	s, _ := io.ReadAll(req.Body)
	fmt.Println(string(s))
}
