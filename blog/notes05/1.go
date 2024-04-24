package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"encoding/json"
)

func main() {
	get_main()
}

// -----
func get_main() { //http://localhost:4567
	e := echo.New()
	e.GET("/", getdata)
	e.Logger.Fatal(e.Start(":4567"))
}
func getdata(c echo.Context) error {
	return c.String(http.StatusOK, "test") //test
}

// -----
func get_main1() { //http://localhost:4567/test
	e := echo.New()
	e.GET("/test", getdata1)
	e.Logger.Fatal(e.Start(":4567"))
}
func getdata1(c echo.Context) error {
	return c.String(http.StatusOK, "test") //test
}

// -----
func get_main2() { //http://localhost:4567/test/a/b
	e := echo.New()
	e.GET("/test/:a/:b", getdata2)
	e.Logger.Fatal(e.Start(":4567"))
}
func getdata2(c echo.Context) error {
	a := c.Param("a")
	b := c.Param("b")
	return c.String(http.StatusOK, a+"|"+b) //a|b
}

// -----
func get_main3() { //http://localhost:4567/test/a?ty=ty
	e := echo.New()
	e.GET("/test/:a", getdata3)
	e.Logger.Fatal(e.Start(":4567"))
}
func getdata3(c echo.Context) error {
	a := c.Param("a")
	ty := c.QueryParam("ty")
	return c.String(http.StatusOK, a+"|"+ty) //a|ty
}

// -----
type datajson struct {
	A int
	B string
	C datajson2
	D []string
	E map[string]int
}
type datajson2 struct {
	A int
	B string
}

func get_main4() { //http://localhost:4567
	e := echo.New()
	e.GET("/", getdata4)
	e.Logger.Fatal(e.Start(":4567"))
}
func getdata4(c echo.Context) error {
	d := datajson{
		A: 1,
		B: "a",
		C: datajson2{
			A: 2,
			B: "b",
		},
		D: []string{"c", "d"},
		E: map[string]int{
			"E": 3,
			"F": 4,
		},
	}
	j, _ := json.Marshal(d)
	return c.JSON(http.StatusOK, j) //"eyJBIjoxLCJCIjoiYSIsIkMiOnsiQSI6MiwiQiI6ImIifSwiRCI6WyJjIiwiZCJdLCJFIjp7IkUiOjMsIkYiOjR9fQ=="
	//base64:{"A":1,"B":"a","C":{"A":2,"B":"b"},"D":["c","d"],"E":{"E":3,"F":4}}
}
