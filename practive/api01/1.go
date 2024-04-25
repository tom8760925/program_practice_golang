package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type T struct { //資料表欄位
	ID int
	V  string
	I  int
	II int
}

var sqlname string = "mysql"
var conn string = "root:password@tcp(localhost:3306)/test?charset=utf8"
var tables string = "tests"
var field []string = []string{"v", "i", "ii"}

func main() { //http://localhost:4567/test
	e := echo.New()
	e.GET("/test", getdata)
	e.POST("/test", postdata)
	e.PUT("/test/:id", putdata)
	e.DELETE("/test/:id", deletedata)
	e.Logger.Fatal(e.Start(":4567"))
}
func getdata(c echo.Context) error {
	var dt []T
	//連結資料庫
	db, err := sql.Open(sqlname, conn)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	//sql語法
	rows, err := db.Query("SELECT * FROM " + tables)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	for rows.Next() {
		var t T
		//讀資料
		err := rows.Scan(&t.ID, &t.V, &t.I, &t.II)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprint(err))
		}
		dt = append(dt, t)
	}
	rows.Close()
	db.Close()
	//轉json
	j, err := json.Marshal(dt)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	//return c.JSON(http.StatusOK, j)//base64
	return c.JSONBlob(http.StatusOK, j)
}
func postdata(c echo.Context) error {
	post_id := 0
	post_v := c.FormValue(field[0])
	//string to int
	post_i, err := strconv.Atoi(c.FormValue(field[1]))
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	//string to int
	post_ii, err := strconv.Atoi(c.FormValue(field[2]))
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	t := T{ID: post_id, V: post_v, I: post_i, II: post_ii}
	//連結資料庫
	db, err := sql.Open(sqlname, conn)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	//sql語法
	pre, err := db.Prepare("INSERT " + tables + " SET v=?,i=?,ii=?")
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	//寫入的資料
	inst, err := pre.Exec(t.V, t.I, t.II)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	//執行
	id, err := inst.LastInsertId()
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	pre.Close()
	db.Close()
	return c.String(http.StatusOK, fmt.Sprint(id))
}
func putdata(c echo.Context) error {
	//string to int
	put_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	put_v := c.FormValue(field[0])
	//string to int
	put_i, err := strconv.Atoi(c.FormValue(field[1]))
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	//string to int
	put_ii, err := strconv.Atoi(c.FormValue(field[2]))
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	t := T{ID: put_id, V: put_v, I: put_i, II: put_ii}
	//連結資料庫
	db, err := sql.Open(sqlname, conn)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	//sql語法
	pre, err := db.Prepare("UPDATE " + tables + " SET v=?,i=?,ii=? where id=?")
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	//寫入的資料
	inst, err := pre.Exec(t.V, t.I, t.II, t.ID)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	//執行
	id, err := inst.RowsAffected()
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	pre.Close()
	db.Close()
	return c.String(http.StatusOK, fmt.Sprint(id))
}
func deletedata(c echo.Context) error {
	delete_id := c.Param("id")
	//連結資料庫
	db, err := sql.Open(sqlname, conn)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	//sql語法
	pre, err := db.Prepare("delete from " + tables + " where id=?")
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	//寫入的資料
	inst, err := pre.Exec(delete_id)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	//執行
	id, err := inst.RowsAffected()
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprint(err))
	}
	pre.Close()
	db.Close()
	return c.String(http.StatusOK, fmt.Sprint(id))
}
