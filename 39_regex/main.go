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

	res2 := regex.FindStringIndex(text)
	fmt.Println(res2)

	// replace all string
	fmt.Println(regex.ReplaceAllString(text, "potato"))

	// replace all string func
	fmt.Println(regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	}))

	str := regex.Split(text, -1)
	fmt.Println(str)

}
