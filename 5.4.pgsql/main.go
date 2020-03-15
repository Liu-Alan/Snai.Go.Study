package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres  password=810618 dbname=test sslmode=disable")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("Insert Into userinfo(username,department,created) Values($1,$2,$3) Returning uid")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	var lastInsertId int
	err = db.QueryRow("Insert Into userinfo(username,department,created) Values($1,$2,$3) Returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("最后插入id=", lastInsertId)

	//更新数据
	stmt, err = db.Prepare("Update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("Select * From userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time

		err = rows.Scan(&uid, &username, &department, &created)

		checkErr(err)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("Delete From userinfo Where uid=$1")
	checkErr(err)

	res, err = stmt.Exec(1)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
