package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// format付きじゃない関数で複数arg指定するとどうなるか
	// ->連結される。
	// {"level":"warning","msg":"aaaabbbcccc","time":"2019-11-29T17:23:35+09:00"}

	logrus.Warn("aaaa", "bbb", "cccc")
}
