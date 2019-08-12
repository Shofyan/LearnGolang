package main

import "fmt"

// unbuffer channel
// done <- true (sender)
// <-done (receiver)
func main() {

	// channel buffer
	hello := make(chan string, 2)
	hello <- "hello"
	hello <- "golang"
	close(hello)

	for v := range hello {
		fmt.Println(v)
	}

	// kosong karena sudah di print di sebelumnya
	fmt.Println(<-hello)
	fmt.Println(<-hello)

}
