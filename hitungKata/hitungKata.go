package main

import(
	"fmt"
	"strings"
)

func hitungKata(kalimat string) int {
	return len(strings.Split(kalimat, " "))
}

func main() {
	fmt.Println(hitungKata("Aku Tanpa Kamu Bisa Apa"))
}