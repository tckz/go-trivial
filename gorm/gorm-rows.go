package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type NotExistTable struct {
	ID int       `gorm:"column:id"`
	Dt time.Time `gorm:"column:dt"`
}

func (t NotExistTable) TableName() string {
	return "not_exist"
}

func main() {

	// "root:xxxxxxx@tcp(localhost:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local"
	optDSN := flag.String("dsn", "", "Connection string")
	flag.Parse()

	if *optDSN == "" {
		panic("dsn must be specified")
	}

	db, err := gorm.Open("mysql", *optDSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)

	// 存在しないテーブルを使ってRowsがエラーを返す状況を作る
	// rows == nil, err != nilとなる。
	rows, err := db.Model(&NotExistTable{}).Rows()
	fmt.Fprintf(os.Stderr, "rows.nil?=%t, err=%v\n", rows == nil, err)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}
