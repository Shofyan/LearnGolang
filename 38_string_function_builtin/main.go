//https://dasarpemrogramangolang.novalagung.com/41-strings.html

package main

import "fmt"
import "strings"

func main() {
	// contain
	var isExist = strings.Contains("tamu indo", "indo")
	fmt.Println(isExist)

	// prefix or start with
	isExist2 := strings.Contains("tamu", "ta")
	fmt.Println(isExist2)

	// suffix or end with word
	isExist3 := strings.HasSuffix("tamu", "mu")
	fmt.Println(isExist3)

	// count char on word
	numbCount := strings.Count("tamu tata", "t")
	fmt.Println(numbCount)

	sarina := strings.Index("tomat", "a")
	fmt.Println(sarina)

	// string replace
	rep := strings.Replace("aku kere", "aku", "shofyan", -1)
	fmt.Println(rep)

	repe := strings.Repeat("nanao", 10)
	fmt.Println(repe)

	split := strings.Split("shofyan,jangan,tukang,bangunan", ",")
	fmt.Println(split)

	arrString := []string{"batman", "spiderman", "cave man"}
	joinstr := strings.Join(arrString, ",")
	fmt.Println(joinstr)

	fmt.Println(strings.ToUpper("alAy"))
	fmt.Println(strings.ToLower("kecil semua"))
}
