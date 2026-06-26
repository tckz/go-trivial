package main

import (
	"flag"
	"log"
	"time"
)

var optDuration = flag.Duration("dur", 42*time.Millisecond, "dur")

func main() {
	flag.Parse()

	if *optDuration == 0 {
		log.Printf("zero!!")
		return
	}

	log.Printf("%s\n", *optDuration)
}
