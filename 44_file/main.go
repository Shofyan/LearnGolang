package main

import (
	"fmt"
	"os"
)

var path = "44_file/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err)
	}

	return err != nil
}

func createFile() {
	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("--> file berhasil dibuat", path)

}

func readFile() {
	// open file
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != nil {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}

	if isError(err) {
		return
	}

	fmt.Println("--> file berhasil dibaca")
	fmt.Println(string(text))

}

func deleteeFile() {
	err := os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("--> file berhasil di delete")

}

func main() {
	//createFile()
	//readFile()
	deleteeFile()
}
