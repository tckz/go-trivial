package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
create table timestamp_sample
(
  id int auto_increment primary key,
  -- no default value. to be set by gorm
  created_at datetime not null
);
*/

type TimestampSample struct {
	ID        int       `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (t TimestampSample) TableName() string {
	return "timestamp_sample"
}

// スライスで受けてレコードがないときはerrにならず、空のスライスが設定される

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

	var recs []TimestampSample

	// id=-1は存在しない前提。
	err = db.Where("id = ?", -1).Find(&recs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// ここには来ない
			log.Printf("err=%#v", err)
			return
		} else {
			panic(err)
		}
	}
	// nil?=false, len=0, []
	fmt.Fprintf(os.Stderr, "nil?=%t, len=%d, %+v\n", recs == nil, len(recs), recs)
}
