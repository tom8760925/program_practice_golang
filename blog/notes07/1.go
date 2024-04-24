package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	put_main()
}
func put_main() { //http://localhost:4567
	e := echo.New()
	e.PUT("/test/:id", postdata)      //覆蓋
	e.PATCH("/test/:id", patchdata)   //修改
	e.DELETE("/test/:id", deletedata) //刪除
	e.Logger.Fatal(e.Start(":4567"))
}
func postdata(c echo.Context) error {

	return c.String(http.StatusOK, "p:"+c.Param("id")+"(p)")
}
func patchdata(c echo.Context) error {

	return c.String(http.StatusOK, "p:"+c.Param("id")+"(p)")
}
func deletedata(c echo.Context) error {

	return c.String(http.StatusOK, "p:"+c.Param("id")+"(d)")
}
