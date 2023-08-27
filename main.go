package main

import (
	"database/sql"
	"github.com/NonokaM/Go-API/models"
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
	// if err := db.Ping(); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("connect to DB")
	// }

	const sqlStr = `
		select title, contents, username, nice
		from articles;
	`

	// Queryメソッドに実行したいSQLSQL, 戻り値rowsに実行結果
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	// rowsの戻り値がtrueの間
	// rows に存在するレコードそれぞれに対して、繰り返し処理を実行する
	for rows.Next() {
	// 変数 article の各フィールドに、取得レコードのデータを入れる
	// (SQL クエリの select 句から、タイトル・本文・ユーザー名・いいね数が返ってくることはわかっている)
		var article models.Article
		err := rows.Scan(&article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		if err != nil {
			fmt.Println(err)
		} else {
			// 読み出し結果を格納した変数 article を、配列に追加
			articleArray = append(articleArray, article)
		}
	}

	fmt.Printf("%+v\n", articleArray)
}
