package main

import (
	"fmt"
)

func hitungMinMax(arrBalita []float64) (float64, float64) {
	min := arrBalita[0]
	max := arrBalita[0]

	for _, berat := range arrBalita {
		if berat < min {
			min = berat
		}
		if berat > max {
			max = berat
		}
	}

	return min, max
}

func rerata(arrBalita []float64) float64 {
	total := 0.0

	for _, berat := range arrBalita {
		total += berat
	}

	return total / float64(len(arrBalita))
}

func main() {
	var n int
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	arrBalita := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBalita[i])
	}

	min, max := hitungMinMax(arrBalita)
	rata := rerata(arrBalita)

	fmt.Printf("\nBerat balita terkecil: %.2f kg\n", min)
	fmt.Printf("Berat balita terbesar: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rata)
}
