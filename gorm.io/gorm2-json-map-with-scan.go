package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

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

var _ sql.Scanner = (*MapWithScan)(nil)
var _ driver.Valuer = (MapWithScan)(nil)

type MapWithScan map[string]interface{}

func (m *MapWithScan) Scan(src interface{}) error {
	log.Printf("src type:%T, v=%v", src, src)
	if u, ok := src.([]uint8); ok {
		if err := json.Unmarshal(u, &m); err != nil {
			return err
		}
		log.Printf("%v", m)
	} else {
		return fmt.Errorf("MapWithScan.Scan: type=%T not supported", src)
	}

	return nil
}

func (m MapWithScan) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

type JSONSample struct {
	ID int         `gorm:"column:id"`
	JS MapWithScan `gorm:"column:js"`
}

func (t JSONSample) TableName() string {
	return "json_sample"
}

// json型の列（objectが入ってるものとする）をmapで受け取れるか->受け取れない
// unsupported data type: &map[]

// json型の列（objectが入ってるものとする）をtyped mapで受け取れるか
// ValuerとScannerを実装していれば通る

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
		fmt.Fprintf(os.Stderr, "js=%v\n", e.JS)
	}
}
