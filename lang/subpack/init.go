package subpack

import (
	"regexp"
)

var RegVar = regexp.MustCompile("[bad")

func Hello() string {
	return "hello"
}
