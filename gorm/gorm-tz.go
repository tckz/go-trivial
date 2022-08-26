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

func main() {

	// "root:xxxxxxx@tcp(localhost:3306)/somedb?charset=utf8mb4&parseTime=True&loc=Local"
	optDSN := flag.String("dsn", "", "Connection string")
	optID := flag.Int("id", 0, "ID of record")
	optUpdate := flag.Bool("update", false, "Update mode")
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

	if *optID != 0 {
		var rec TZSample
		err = db.Where("id = ?", *optID).Find(&rec).Error
		if err != nil {
			fmt.Fprintf(os.Stderr, "gorm.ErrRecordNotFound? = %t: %T\n", err == gorm.ErrRecordNotFound, err)
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "%+v\n", rec)

		if *optUpdate {
			rec.Dt = time.Now()
			err := db.Save(rec).Error
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(os.Stderr, "%+v\n", rec)
		}
	} else {
		rec := &TZSample{
			Dt: time.Now(),
		}

		fmt.Fprintf(os.Stderr, "%+v\n", rec)
		err = db.Create(rec).Error
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "%+v\n", rec)
	}
}
