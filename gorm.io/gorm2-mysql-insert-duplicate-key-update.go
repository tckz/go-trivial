package main

import (
	"flag"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

/*
CREATE TABLE `ins_dup_update` (
  `some_col` varchar(64) NOT NULL,
  `some_id` int(11) NOT NULL,
  `created_at` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`some_col`)
);
*/

/*
mysqlのinsert on duplicate key updateのrows affectedの確認
このURLによればmysqlのC APIでの接続時パラメータで数値が変わるケースがあるらしく試したもの。
https://stackoverflow.com/questions/55750253/mysql-insert-on-duplicate-key-update-return-rows-affected-1-but-no-chang

gorm(go-sql-driver/mysql)の場合は、変わらないっぽい。
insert時: 1
update時: 2
update時に同値設定: 0
*/

func main() {
	optSomeCol := flag.String("some-col", "some_col_val", "value for some_col")
	optSomeID := flag.Int("some-id", 0, "value for some_id")

	// "root:xxxxxxx@tcp(localhost:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local"
	optDSN := flag.String("dsn", "", "Connection string")
	flag.Parse()

	if *optDSN == "" {
		panic("dsn must be specified")
	}

	if *optSomeID == 0 {
		panic("some-id must be specified")
	}

	db, err := gorm.Open(mysql.Open(*optDSN), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		if d, err := db.DB(); err == nil {
			d.Close()
		}
	}()

	ret := db.Unscoped().Exec(`
insert ins_dup_update (some_col, some_id, created_at, updated_at) values 
(?, ?, now(3), now(3))
on duplicate key update 
updated_at = if(some_id <> values(some_id), now(3), updated_at)
,some_id = if(some_id <> values(some_id), values(some_id), some_id) 
`, *optSomeCol, *optSomeID)

	log.Printf("err=%v, affected=%d", ret.Error, ret.RowsAffected)
}
