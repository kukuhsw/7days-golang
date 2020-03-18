package main

import (
	"fmt"
	"strconv"
)

func konversiMenit(menit int) string {
	var hour int = menit / 60
	var minute int = menit % 60
	var waktu string

	if hour >= 10 {
		waktu = waktu + strconv.Itoa(hour)
	} else {
		waktu = waktu + "0" + strconv.Itoa(hour)
	}

	if minute >= 10 {
		waktu = waktu + ":" + strconv.Itoa(minute)
	} else {
		waktu = waktu + ":" + "0" + strconv.Itoa(minute)
	}

	return waktu

}

func main() {
	fmt.Println(konversiMenit(120))
}
