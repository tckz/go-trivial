package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	jst, _ := time.LoadLocation("Asia/Tokyo")

	birthDay := 19720101
	f := time.Date(2019, 1, 1, 0, 0, 0, 0, jst)
	t := time.Date(2021, 1, 1, 0, 0, 1, 0, jst)

	d := f
	for d.Before(t) {
		s := d.Format("20060102")
		nowInt, _ := strconv.Atoi(s)

		age := (nowInt - birthDay) / 10000
		fmt.Fprintf(os.Stdout, "%s -> %d\n", s, age)

		d = d.AddDate(0, 0, 1)
	}

}
