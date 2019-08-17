package main

import (
	"fmt"
	"log"
)
import "encoding/base64"

func main() {
	var data = "jhon wick"
	var encodingString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded :", encodingString)

	var decodingByte, _ = base64.StdEncoding.DecodeString(encodingString)
	fmt.Println(string(decodingByte))

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodingString2 = string(encoded)
	fmt.Println(encodingString2)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var val, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		log.Print(err)
	}
	var decodedString = string(decoded)
	fmt.Println(decodedString)
	fmt.Println(val)

	// encode decode url

	var urlString = "https://dasarpemrogramangolang.novalagung.com/43-encoding-base64.html"
	encodeUrl := base64.URLEncoding.EncodeToString([]byte(urlString))
	fmt.Println(encodeUrl)

	decodeUrl, err := base64.URLEncoding.DecodeString(encodeUrl)
}
