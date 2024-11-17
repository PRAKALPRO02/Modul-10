package main

import "fmt"

type arrBalita [100]float64

var jmlBalita int

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 0; i < jmlBalita; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rataRata(arrBerat arrBalita) float64 {
	var total float64
	for i := 0; i < jmlBalita; i++ {
		total += arrBerat[i]
	}
	return total / float64(jmlBalita)
}

func main() {
	var balitas arrBalita
	var min, max float64

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&jmlBalita)

	for i := 0; i < jmlBalita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&balitas[i])
	}

	hitungMinMax(balitas, &min, &max)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("rataRata berat balita: %.2f kg\n", rataRata(balitas))
}
