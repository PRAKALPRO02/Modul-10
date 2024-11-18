package main

import (
	"fmt"
)

type arrBalita [100]float64

var jml_balita int

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 0; i < jml_balita; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func Rerata(arrBerat arrBalita) float64 {
	var jumlah float64
	for _, value := range arrBerat {
		jumlah += value
	}
	return jumlah / float64(jml_balita)
}

func main() {
	var balitas arrBalita
	var min, max float64

	fmt.Print("Masukan Banyak data berat balita: ")
	fmt.Scan(&jml_balita)
	for i := 0; i < jml_balita; i++ {
		fmt.Print("Masukan berat balita ke-", i+1, ": ")
		fmt.Scan(&balitas[i])
	}
	hitungMinMax(balitas, &min, &max)
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", Rerata(balitas))
}
