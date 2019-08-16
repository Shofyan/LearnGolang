//https://dasarpemrogramangolang.novalagung.com/39-timer.html
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	/*	fmt.Println("start")
		time.Sleep(time.Second * 4)
		fmt.Println("after 4 second")

		for true{
			fmt.Println("hai dunia")
			time.Sleep(1 * time.Second)
		}*/

	/*	var timer = time.NewTimer(4 * time.Second)
		fmt.Println("start")
		<-timer.C
		fmt.Print("finish")*/

	/*	var ch = make(chan bool)
		time.AfterFunc(4*time.Second,func(){
			fmt.Println("expired")
			ch <- true
		})

		fmt.Println("start")
		<-ch
		fmt.Println("end")

	*/
	//  time.After function
	//<-time.After(4* time.Second)
	//fmt.Println("Expired")

	//  scheduler using ticker

	// timer and Goroutine combination

	var timeout = 5
	var ch = make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Println("what is 725/25? ?")
	_, _ = fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right!")

	} else {
		fmt.Println("the answer is wrong !")
	}

}

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second,
		func() {
			ch <- true
		})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\n time out! no answer more than", timeout, "seconds")
	os.Exit(0)

}
