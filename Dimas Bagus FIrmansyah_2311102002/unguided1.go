package main

import (
	"fmt"
)

// Tipe data array untuk berat balita
type arrBalita [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat []float64, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for _, berat := range arrBerat {
		if berat < *bMin {
			*bMin = berat
		}
		if berat > *bMax {
			*bMax = berat
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func rataRata(arrBerat []float64) float64 {
	total := 0.0
	for _, berat := range arrBerat {
		total += berat
	}
	return total / float64(len(arrBerat))
}

func main() {
	var n int
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	var arrBerat []float64
	for i := 0; i < n; i++ {
		var berat float64
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat)
		arrBerat = append(arrBerat, berat)
	}

	var bMin, bMax float64
	hitungMinMax(arrBerat, &bMin, &bMax)

	rata := rataRata(arrBerat)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Berat rata-rata balita: %.2f kg\n", rata)
}
