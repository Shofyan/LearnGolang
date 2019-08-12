package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Worker Pool")

	startTime := time.Now()

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go consumer(1, jobs, results)
	go consumer(2, jobs, results)
	go producer(jobs, 10)
	for i := 1; i <= 10; i++ {
		res := <-results
		fmt.Println("hasil ke-", res)
	}

	fmt.Println("====================")

	elapse := time.Since(startTime)
	fmt.Println("Waktu ", elapse)

}

// fakeHTTPRequest	adalah fungsi palsu http wannabe
func fakeHTTPRequest(x int) int {
	return x * 10
}

func producer(jobs chan<- int, size int) {
	for i := 1; i <= size; i++ {
		jobs <- i
	}
	close(jobs)
}

func consumer(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Consumer ke-", id, " Mulai")
		time.Sleep(2 * time.Second)

		results <- fakeHTTPRequest(job)
	}
}
