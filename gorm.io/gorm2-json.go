package main

import (
	"flag"
	"fmt"
	"os"

	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

/*
CREATE TABLE `json_sample` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `js` json DEFAULT NULL,
  PRIMARY KEY (`id`));
*/

type JSONSample struct {
	ID int            `gorm:"column:id"`
	JS datatypes.JSON `gorm:"column:js"`
}

func (t JSONSample) TableName() string {
	return "json_sample"
}

// json型の列はdatatypes.JSONで受け取れる

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

	var recs []JSONSample
	result := db.Find(&recs)
	if result.Error != nil {
		panic(result.Error)
	}

	for i, e := range recs {
		fmt.Fprintf(os.Stderr, "%d: %+v\n", i, e)
		fmt.Fprintf(os.Stderr, "js=%s\n", e.JS.String())
	}
}
