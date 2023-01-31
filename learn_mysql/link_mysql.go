package learn_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func LinkSql() {
	dsn := "admin:123321@tcp(192.168.174.131:3306)/learn_sql"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("db err:", err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("fuck something wrong！")
		}
	}(db)
	fmt.Println("connect success!~")
}
