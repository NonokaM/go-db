package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"

	// DBに接続するためのアドレス文を定義
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// Open関数を用いてDBに接続
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	// プログラムが終了するときに、コネクションがCloseされるようにする
	defer db.Close()

	// sql.DB型のPingメソッドで疎通確認
	if err := db.Ping(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect to DB")
	}
}
