package main

import (
	"fmt"
	"os/exec"
)

func main() {
	data, _ := exec.Command("ls").Output()
	fmt.Printf("-> ls \n %s \n", string(data))

}
