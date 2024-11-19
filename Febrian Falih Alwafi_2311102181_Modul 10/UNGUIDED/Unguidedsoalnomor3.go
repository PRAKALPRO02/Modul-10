package main

import "fmt"

func hitungMinMax_2311102181(arrBerat []float64) (float64, float64) {
	min, max := arrBerat[0], arrBerat[0]
	for _, berat := range arrBerat {
		if berat < min {
			min = berat
		}
		if berat > max {
			max = berat
		}
	}
	return min, max
}

func rataRata(arrBerat []float64) float64 {
	var total float64
	for _, berat := range arrBerat {
		total += berat
	}
	return total / float64(len(arrBerat))
}

func main() {
	var jmlBalita int
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&jmlBalita)

	balitas := make([]float64, jmlBalita)
	for i := 0; i < jmlBalita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&balitas[i])
	}

	min, max := hitungMinMax_2311102181(balitas)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata(balitas))
}
