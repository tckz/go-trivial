package main

// closeしたchanをcloseするとpanic

func main() {

	ch := make(chan int)

	close(ch)

	// panic: close of closed channel
	close(ch)
}
