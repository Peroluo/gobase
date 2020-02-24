package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() (db *sql.DB, err error) {
	dsn := "root:ectripaes2013@tcp(202.96.155.121:3306)/golist"
	db, err = sql.Open("mysql", dsn)
	return
}

func transcations(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	// 执行多个sql操作
	sqlStr1 := `update user set age = age -2 where id = 19`
	sqlStr2 := `update user set age = age +2 where id = 17`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		fmt.Println("sqlStr1执行报错")
		tx.Rollback()
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		fmt.Println("sqlStr2执行报错")
		tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("sqlStr2执行报错")
		tx.Rollback()
		return
	}
	fmt.Println("事务执行成功！")
}

func main() {
	db, err := initDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("数据库已连接！")
	transcations(db)
}
