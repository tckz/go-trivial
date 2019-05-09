package main

import (
	"fmt"
	"log"

	"github.com/samclarke/robotstxt"
)

func main() {
	// こういうrobots.txtはすべてを許可する、ということが
	// Google Developersやwikipediaに書いてあるのだが、
	// https://developers.google.com/search/reference/robots_txt?hl=en#disallow
	// https://en.wikipedia.org/wiki/Robots_exclusion_standard
	// samclarke/robotstxtの挙動はこれらと異なっていて参考にできなかった、
	// というもの

	url := "http://www.example.com/robots.txt"
	contents := `
	User-agent: *
	Disallow:
    `

	robots, err := robotstxt.Parse(contents, url)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_ = robots

	allowed, err := robots.IsAllowed("Sams-Bot/1.0", "http://www.example.com/test.html")
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("allowed: %v\n", allowed)
}
