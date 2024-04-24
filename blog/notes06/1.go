package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	post_main()
}
func post_main() { //http://localhost:4567
	e := echo.New()
	e.POST("/", postdata)
	e.Logger.Fatal(e.Start(":4567"))
}
func postdata(c echo.Context) error {
	a := c.FormValue("a")
	return c.String(http.StatusOK, "p:"+a)
}
