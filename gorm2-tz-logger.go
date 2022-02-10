package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
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

		TZ=UTC go run ./gorm2-tz-logger.go -dsn "root:xxxxxxx@tcp(localhost:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local&time_zone=UTC"
	*/
	/*
		from=2022-02-10 17:23:08.137268635 +0900 JST
		2022/02/10 08:23:08 /xxxxx/gorm2-tz-logger.go:92
		[3.212ms] [rows:1] SELECT * FROM `tzsample` WHERE dt < '2022-02-10 17:23:08.137'
		{ID:1 Dt:2022-02-09 02:06:15 +0000 UTC}
	*/
	// mysqld側のgeneral log
	/*
		2022-02-10T08:23:08.136047Z         2 Connect   root@xxxxx on somedb using TCP/IP
		2022-02-10T08:23:08.136327Z         2 Query     SET NAMES utf8mb4
		2022-02-10T08:23:08.136670Z         2 Query     SET time_zone=UTC
		2022-02-10T08:23:08.137022Z         2 Query     SELECT VERSION()
		2022-02-10T08:23:08.140245Z         2 Prepare   SELECT * FROM `tzsample` WHERE dt < ?
		2022-02-10T08:23:08.140326Z         2 Execute   SELECT * FROM `tzsample` WHERE dt < '2022-02-10 08:23:08.137268635'
		2022-02-10T08:23:08.140567Z         2 Close stmt
		2022-02-10T08:23:08.140738Z         2 Quit
	*/

	// "root:xxxxxxx@tcp(localhost:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local"
	optDSN := flag.String("dsn", "", "Connection string")
	flag.Parse()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	if *optDSN == "" {
		panic("dsn must be specified")
	}

	gl := gormlogger.Default.LogMode(gormlogger.Info)
	db, err := gorm.Open(mysql.Open(*optDSN), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 gl,
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		if d, err := db.DB(); err == nil {
			d.Close()
		}
	}()

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
