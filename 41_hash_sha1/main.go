package main

import (
	"fmt"
	"time"
)
import "crypto/sha1"

func main() {
	text := "this is secrect "
	sha := sha1.New()
	sha.Write([]byte(text))
	encripted := sha.Sum(nil)
	encriptedString := fmt.Sprintf("%x", encripted)
	fmt.Println(encriptedString)

	// USING salt for  encription
	fmt.Println("original %s\n\n", text)

	var hashed, salt1 = doHashUsingSalt(text)
	fmt.Printf(" hashed 1 : %s \n \n", hashed)
	_ = salt1

}

func doHashUsingSalt(text string) (string, string) {
	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	saltedText := fmt.Sprintf("text :'%s', salt: %s ", text, salt)
	fmt.Println(saltedText)
	sha := sha1.New()
	sha.Write([]byte(saltedText))
	encripted := sha.Sum(nil)

	return fmt.Sprintf("%x", encripted), salt

}
