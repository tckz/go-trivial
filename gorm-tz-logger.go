package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
create table tzsample
(
  id int auto_increment primary key,
  dt datetime not null
);
*/

type TZSample struct {
	ID int       `gorm:"column:id"`
	Dt time.Time `gorm:"column:dt"`
}

func (t TZSample) TableName() string {
	return "tzsample"
}

/*
gormのロガーから出るSQLに含まれるtime.Time部分のTZがdriverのTZ（loc）とは連動してない、という確認
gorm loggerはtime.TimeのTZを気にせず年月日時分秒でformatする。
実際MySQLに投げるのはmysql driverが担当し、driver内でIn(loc)している。
*/

func main() {

	/*
		このモジュールを実行するときはTZ=UTCで実行

		TZ=UTC go run ./gorm-tz-logger.go -dsn "root:xxxxxxx@tcp(localhost:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local&time_zone=UTC"
	*/
	/*
		from=2022-02-09 11:56:14.940925408 +0900 JST
		(/xxxxx/gorm-tz-logger.go:69)
		[2022-02-09 02:56:14]  [0.88ms]  SELECT * FROM `tzsample`  WHERE (dt < '2022-02-09 11:56:14')
		[3 rows affected or returned ]
	*/
	// mysqld側のgeneral log
	/*
		2022-02-09T02:56:14.940588Z        17 Connect   root@somesome on somedb using TCP/IP
		2022-02-09T02:56:14.940704Z        17 Query     SET NAMES utf8mb4
		2022-02-09T02:56:14.940793Z        17 Query     SET time_zone=UTC
		2022-02-09T02:56:14.941216Z        17 Prepare   SELECT * FROM `tzsample`  WHERE (dt < ?)
		2022-02-09T02:56:14.941523Z        17 Execute   SELECT * FROM `tzsample`  WHERE (dt < '2022-02-09 02:56:14.940925408')
		2022-02-09T02:56:14.941869Z        17 Close stmt
		2022-02-09T02:56:14.941967Z        17 Quit
	*/

	// "root:xxxxxxx@tcp(localhost:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local&time_zone=UTC"
	optDSN := flag.String("dsn", "", "Connection string")
	flag.Parse()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	if *optDSN == "" {
		panic("dsn must be specified")
	}

	db, err := gorm.Open("mysql", *optDSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)

	var rec TZSample
	from := time.Now().In(jst)
	fmt.Fprintf(os.Stderr, "from=%s", from)

	err = db.Where("dt < ?", from).Find(&rec).Error
	if err != nil {
		fmt.Fprintf(os.Stderr, "gorm.ErrRecordNotFound? = %t: %T\n", errors.Is(err, gorm.ErrRecordNotFound), err)
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "%+v\n", rec)
}
