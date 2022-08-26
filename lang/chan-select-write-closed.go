package main

// closeしたchanを書き込みでselectした状態で書き込む

func main() {

	ch := make(chan int)

	close(ch)

	// panic: send on closed channel
	select {
	case ch <- 1:
	}
}
