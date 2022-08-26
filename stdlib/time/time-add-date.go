package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	{
		tm := time.Date(2022, 5, 31, 0, 0, 0, 0, time.UTC)
		add := tm.AddDate(0, -1, 0)

		/*
			org=2022-05-31 00:00:00 +0000 UTC
			-1 month=2022-05-01 00:00:00 +0000 UTC
		*/
		// 5/31 -> 4/31 = 5/01 になる
		fmt.Fprintf(os.Stderr, "org=%s\n", tm)
		fmt.Fprintf(os.Stderr, "-1 month=%s\n", add)
	}
	{
		tm := time.Date(2022, 8, 31, 0, 0, 0, 0, time.UTC)
		add := tm.AddDate(0, -2, 0)

		/*
			org=2022-08-31 00:00:00 +0000 UTC
			-2 month=2022-07-01 00:00:00 +0000 UTC
		*/
		// 8/31 -> 6/31 = 7/01 になる
		fmt.Fprintf(os.Stderr, "org=%s\n", tm)
		fmt.Fprintf(os.Stderr, "-2 month=%s\n", add)
	}
	{
		tm := time.Date(2022, 7, 31, 0, 0, 0, 0, time.UTC)
		add := tm.AddDate(0, -2, 0)

		/*
			org=2022-07-31 00:00:00 +0000 UTC
			-2 month=2022-05-31 00:00:00 +0000 UTC
		*/
		// 7/31 -> 5/31 になる
		// 1 monthずつ引くと、 7/31 -> 6/31(=7/01) -> 6/01になりそうにも思えるが逐次繰り返すわけではない
		fmt.Fprintf(os.Stderr, "org=%s\n", tm)
		fmt.Fprintf(os.Stderr, "-2 month=%s\n", add)
	}
}
