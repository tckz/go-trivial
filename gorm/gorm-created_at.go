package main

import (
	"flag"
	"fmt"
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

// gormの仕様でcreated_at列に勝手に現在時刻を入れる（DB側のdefaultではない）
// このとき、現在時刻をどこから持ってくるか？
// -> アプリケーションプロセス側の現在時刻が使われる。Now()とかではない。

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

	rec := &TimestampSample{}

	fmt.Fprintf(os.Stderr, "%+v\n", rec)
	/*
		(/some/path/gorm-created_at.go:52)
		[2020-09-02 14:22:34]  [0.69ms]  INSERT INTO `timestamp_sample` (`created_at`) VALUES ('2020-09-02 14:22:34')
		[1 rows affected or returned ]
	*/
	err = db.Create(rec).Error
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "%+v\n", rec)
}
