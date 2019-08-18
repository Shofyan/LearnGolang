package main

import "fmt"
import "net/url"

func main() {
	var urlString = "http://kalipare.com:80/hello?name=shofyan&age=23"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Printf("url: %s\n", urlString)
	fmt.Printf("protocol: %s\n", u.Scheme) // http
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s \n ", name, age)

}
