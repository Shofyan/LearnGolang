//https://dasarpemrogramangolang.novalagung.com/42-regex.html

package main

import "fmt"
import "regexp"

func main() {
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err)
	}

	res1 := regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)

	res2 := regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)

}
