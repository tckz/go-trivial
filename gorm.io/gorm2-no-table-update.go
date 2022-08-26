package main

import (
	"flag"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// update対象テーブル名のヒントがない状態でupdate
// ->普通にエラーに

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

	// unsupported data type: map[created_at:2021-12-15 16:33:43.798453006 +0900 JST m=+0.007009833]: Table not set, please set it like: db.Model(&user) or db.Table("users")
	err = db.Unscoped().Where("id = ?", 1).Update("created_at", time.Now()).Error
	if err != nil {
		panic(err)
	}
}
