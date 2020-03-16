package main

import (
	"fmt"
	"unicode"
)

func urutanKecilBesar(str string) string {
	var resStr string
	for _, v := range str {
		if v == unicode.ToLower(rune(v)) {
			resStr = resStr + string(unicode.ToUpper(rune(v)))
		} else {
			resStr = resStr + string(unicode.ToLower(rune(v)))
		}
	}
	return resStr
}

func main() {
	fmt.Println(urutanKecilBesar("R3d D3VIl"))
	fmt.Println(urutanKecilBesar("MakaN nAsI gOrEnG"))
	fmt.Println(urutanKecilBesar("kOmAR iS mE"))
	fmt.Println(urutanKecilBesar("k4mU J4h4t"))
}
