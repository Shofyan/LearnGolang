// reference : https://dasarpemrogramangolang.novalagung.com/38-time.html
package main

import (
	"fmt"
)
import "time"

func main() {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)

	var time2 = time.Date(2011, 12, 24, 120, 37, 98, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)

	// method time, which their own
	var now = time.Now()
	fmt.Println("year :", now.Year(), "month :", now.Month(), "\n")

	// parse string to time
	var date, _ = time.Parse("2006-01-02 15:04:05", "2015-09-02 00:00:00")
	fmt.Println(date.String())

	// parsing time
	date, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	var dateS1 = date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dates1", dateS1)

	var dateS2 = date.Format(time.RFC3339)
	fmt.Println(dateS2, dateS2)

	date, err := time.Parse("06 jan 15", "02 Sep 15 08:00 WIB")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(date)
}
