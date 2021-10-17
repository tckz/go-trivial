package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	optCount := flag.Int("count", 1000, "number of generation")
	optStdDev := flag.Float64("stddev", 1, "standard deviation")
	optMean := flag.Float64("mean", 0, "desired mean")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < *optCount; i++ {
		n := rand.NormFloat64()**optStdDev + *optMean
		fmt.Printf("%f\n", n)
	}
}
