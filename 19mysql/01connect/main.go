package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	age  int
	name string
	id   int
}

func initDB() (db *sql.DB, err error) {
	dsn := "root:ectripaes2013@tcp(202.96.155.121:3306)/golist"
	db, err = sql.Open("mysql", dsn)
	return
}

func find(db *sql.DB) (user1 user, err error) {
	sqlStr := "select id, `name`, age from user where id=?"
	var u user
	err = db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Println(err)
		return u, err
	}
	return u, nil
}

func search(db *sql.DB) []user {
	sqlStr := "select id, name, age from user where id>?"
	users := []user{}
	rows, err := db.Query(sqlStr, 0)
	defer rows.Close()
	if err != nil {
		return users
	}
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, u)
	}
	return users
}

func insert(db *sql.DB) {
	sqlstr := `insert into user(name, age) values("啊歐陽",28)`
	ret, err := db.Exec(sqlstr)
	if err != nil {
		fmt.Println(err)
	}
	id, err := ret.LastInsertId()
	fmt.Println(id)
}

func update(db *sql.DB) {
	sqlstr := `update user set age = ? where id > ?`
	ret, err := db.Exec(sqlstr, 80, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret.RowsAffected())

}

func delete(db *sql.DB) {
	sqlstr := `delete from user where id = ?`
	ret, err := db.Exec(sqlstr, 16)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret.RowsAffected())
}

// 预处理插入多条数据
// 适用于同一条sql语句，执行多次的情况
func prepareInsert(db *sql.DB) {
	sqlStr := `insert into user(name,age) values(?,?)`
	smt, err := db.Prepare(sqlStr)
	defer smt.Close()
	if err != nil {
		fmt.Println(err)
	}
	var users = make(map[string]int, 8)
	users["小王1"] = 10
	users["小王2"] = 11
	users["小王3"] = 12
	for k, s := range users {
		result, err := smt.Exec(k, s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result.LastInsertId())
	}
}

func main() {
	db, err := initDB()
	defer db.Close()
	if err != nil {
		fmt.Println(db)
	}
	db.SetMaxOpenConns(10) // 设置数据库连接池最大连接数
	db.SetMaxIdleConns(5)  // 设置最大空闲连接数
	fmt.Println("数据库已连接！")
	// 查询单条
	// u, err := find(db)
	// fmt.Println(u)

	// 查询多条数据
	users := search(db)
	fmt.Println(users)

	// 添加一条数据
	// insert(db)

	// 更新数据
	// update(db)

	// 删除数据
	// delete(db)

	// 插入多条预处理
	// prepareInsert(db)
}
