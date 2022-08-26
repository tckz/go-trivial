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

	/*
		[2020-10-23 15:54:54]  [0.64ms]  SELECT count(distinct(id)) FROM `timestamp_sample`
		[0 rows affected or returned ]

		uniqueCount=2
	*/
	/* mysqlで同じSQL実行すると1行扱いだけど。この差はどこから。
	1 row in set (0.00 sec)
	*/
	var uniqueCount int64
	err = db.Select("count(distinct(id))").
		Model(TimestampSample{}).
		Count(&uniqueCount).Error
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "uniqueCount=%d\n", uniqueCount)
}
