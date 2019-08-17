package main

import "fmt"
import "os"
import "flag"

func main() {
	argsRaw := os.Args
	fmt.Printf("-> %#v \n", argsRaw)

	argsRaw = os.Args[1:]
	fmt.Printf("-> %#v \n", argsRaw)

	//	name := flag.String("name","anonymous","type your name")
	var nameVal string
	flag.StringVar(&nameVal, "name", "anonymous", "type your name")
	age := flag.Int64("age", 25, "type your age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", nameVal)

	//	fmt.Printf("name\t: %s\n",*name)
	fmt.Printf("age\t: %d\n", *age)

}
