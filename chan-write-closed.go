package main

// closeしたchanに書き込むとpanic

func main() {

	ch := make(chan int)

	close(ch)

	// panic: send on closed channel
	ch <- 1
}
