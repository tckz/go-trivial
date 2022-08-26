package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// "root:xxxxxxx@tcp(localhost:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local"
	// "root:xxxxxxx@tcp(localhost:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"
	optDSN := flag.String("dsn", "", "Connection string")
	flag.Parse()

	if *optDSN == "" {
		panic("dsn must be specified")
	}

	db, err := sql.Open("mysql", *optDSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// dsnにmultiStatements=trueがないと複数SQLはエラーになる
	rows, err := db.Query(`
select now();
SET time_zone = 'Asia/Tokyo';
select now(), now();
`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	w := csv.NewWriter(os.Stdout)
	w.Comma = '\t'
	defer w.Flush()

	for {
		cols, err := rows.Columns()
		if err != nil {
			panic(err)
		}
		w.Write(cols)

		// TSV出力のように変数の型を気にしない状況であればSQLの結果型がなんであってもstringとして受け取ることはできる
		rowCols := make([]string, len(cols))
		rowColsPointer := make([]interface{}, len(cols))
		for i, _ := range rowCols {
			rowColsPointer[i] = &rowCols[i]
		}

		for rows.Next() {
			if err := rows.Scan(rowColsPointer...); err != nil {
				panic(err)
			}

			w.Write(rowCols)
		}

		if err = rows.Err(); err != nil {
			panic(err)
		}

		if !rows.NextResultSet() {
			if err := rows.Err(); err != nil {
				panic(err)
			}

			break
		}

		w.Write([]string{})
	}
}
