package main

import (
	"database/sql"
	"flag"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// mysqlの接続文字列でセッションのタイムゾーンを変更する
// now()の結果が変わる

func main() {

	// "root:xxxxxxx@tcp(localhost:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local"
	// "root:xxxxxxx@tcp(localhost:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local&time_zone=%27Asia%2FTokyo%27"
	optDSN := flag.String("dsn", "", "Connection string")
	flag.Parse()

	if *optDSN == "" {
		panic("dsn must be specified")
	}

	db, err := sql.Open("mysql", *optDSN)
	if err != nil {
		panic(err)
	}

	var t time.Time
	err = db.QueryRow(`select now()`).Scan(&t)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", t)
}
