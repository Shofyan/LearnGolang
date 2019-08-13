package main

import (
	"crypto/sha1"
	"fmt"

	p "github.com/wuriyanto48/go-pbkdf2"
)

func main() {
	pass := p.NewPassword(sha1.New, 8, 32, 1500)
	hashed := pass.HashPassword("admin")

	fmt.Println(hashed.CipherText)
	fmt.Println(hashed.Salt)
}
