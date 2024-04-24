package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type t struct {
	id int
	v  string
	i  int
	ii int
}

func main() {
	selectdata()
}
func selectdata() {
	var dt []t
	db, _ := sql.Open("mysql", "root:password@tcp(localhost:3306)/test?charset=utf8")
	rows, _ := db.Query("SELECT * FROM tests")
	for rows.Next() {
		var t t
		_ = rows.Scan(&t.id, &t.v, &t.i, &t.ii)
		dt = append(dt, t)
	}
	fmt.Println("p:" + fmt.Sprint(dt[0].id))
	rows.Close()
	db.Close()
}
func insertdata() {
	db, _ := sql.Open("mysql", "root:password@tcp(localhost:3306)/test?charset=utf8")
	pre, _ := db.Prepare("INSERT tests SET v=?,i=?,ii=?")
	inst, _ := pre.Exec("c", "3", "3")
	id, _ := inst.LastInsertId()
	fmt.Println("p:" + fmt.Sprint(id))
}
func insertdata1() {
	db, _ := sql.Open("mysql", "root:password@tcp(localhost:3306)/test?charset=utf8")
	inst, _ := db.Exec("INSERT INTO tests (v, i, ii) VALUES (?, ?, ?)", "d", 4, 4)
	id, _ := inst.LastInsertId()
	fmt.Println("p:" + fmt.Sprint(id))
}
func deletedata() {
	db, _ := sql.Open("mysql", "root:password@tcp(localhost:3306)/test?charset=utf8")
	pre, _ := db.Prepare("delete from tests where id=?")
	inst, _ := pre.Exec(5)
	id, _ := inst.RowsAffected()
	fmt.Println("p:" + fmt.Sprint(id))
}
