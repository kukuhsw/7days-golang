package main

import (
	"fmt"
	"strconv"
)

func lanjutPalindrome(angka int) int {
	for {
		angka = angka + 1
		if angkaPalindrome(angka) {
			return angka
		}
	}
}

func angkaPalindrome(angka int) bool {
	return strconv.Itoa(angka) == balikKata(strconv.Itoa(angka))
}

func balikKata(kata string) string {
	var balik string

	for i := 0; i < len(kata); i++ {
		balik = balik + kata[len(kata)-1-i:len(kata)-i]
	}

	return balik
}

func main() {
	fmt.Println(lanjutPalindrome(4))
	fmt.Println(lanjutPalindrome(34))
	fmt.Println(lanjutPalindrome(145))
	fmt.Println(lanjutPalindrome(278))
	fmt.Println(lanjutPalindrome(1567))
}
