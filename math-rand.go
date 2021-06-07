package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Give 1 to seed means same as default seed of math/globalRand
	optSeed := flag.Int64("seed", time.Now().UnixNano(), "The seed for a sequence of random numbers")
	optFrom := flag.Int("from", 2, "rand.Intn(n)")
	optTo := flag.Int("to", 10, "rand.Intn(n)")
	optCount := flag.Int("count", 10, "How many tries")
	flag.Parse()

	/*
		        Intn(2) Intn(3) Intn(4) Intn(5) Intn(6) Intn(7) Intn(8) Intn(9) Intn(10)
		1       1       2       1       1       5       6       1       5       1
		2       1       0       3       2       3       2       7       6       7
		3       1       2       3       2       5       1       7       2       7
		4       1       2       3       4       5       0       3       2       9
		5       1       1       1       1       1       6       1       4       1
		6       0       0       2       3       0       4       6       6       8
		7       1       1       1       0       1       3       1       7       5
		8       0       2       0       0       2       5       4       8       0
		9       0       1       0       1       4       4       0       4       6
		10      0       0       0       0       0       1       4       6       0
	*/
	series := make([][]int, 0, *optTo-*optFrom+1)
	for n := *optFrom; n <= *optTo; n++ {
		fmt.Printf("\tIntn(%d)", n)
		rand.Seed(*optSeed)
		seq := make([]int, 0, *optCount)
		for i := 0; i < *optCount; i++ {
			seq = append(seq, rand.Intn(n))
		}
		series = append(series, seq)
	}
	fmt.Printf("\n")

	for i := 0; i < *optCount; i++ {
		fmt.Printf("%d", i+1)
		for n := 0; n < *optTo-*optFrom+1; n++ {
			fmt.Printf("\t%d", series[n][i])
		}
		fmt.Printf("\n")
	}
}
