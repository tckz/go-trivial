package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
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

	var rec TimestampSample

	// id=-1は存在しない前提。
	// gorm v2はFindが保存先が非sliceであってもErrRecordNotFoundを返さない
	// Take/First/LastはErrRecordNotFoundを返す
	err = db.Where("id = ?", -1).Find(&rec).Error
	// err == nilになる
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 通らない
			log.Printf("err=%#v", err)
			return
		} else {
			panic(err)
		}
	}
	fmt.Fprintf(os.Stderr, "%+v\n", rec)
}
