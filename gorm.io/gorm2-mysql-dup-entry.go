package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	dmysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

/*
create table duplicate_entry (id int ,primary key(id));
*/

type DupSample struct {
	ID int `gorm:"column:id"`
}

func (t DupSample) TableName() string {
	return "duplicate_entry"
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

	// 二回実行で
	// Error 1062: Duplicate entry '1' for key 'PRIMARY'
	err = db.Unscoped().Create(&DupSample{ID: 1}).Error
	if err != nil {
		var me *dmysql.MySQLError
		if errors.As(err, &me) {
			// Asで識別できる
			fmt.Fprintf(os.Stderr, "mes=%s, code=%d\n", me.Message, me.Number)
		}
		panic(err)
	}
}
