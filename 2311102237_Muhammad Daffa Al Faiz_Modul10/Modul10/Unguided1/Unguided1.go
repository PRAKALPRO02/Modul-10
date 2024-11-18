package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int) (float64, float64) {
	min, max := arrBerat[0], arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < min {
			min = arrBerat[i]
		}
		if arrBerat[i] > max {
			max = arrBerat[i]
		}
	}
	return min, max
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var banyakBalita int
	var beratBalita arrBalita

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scanln(&banyakBalita)

	for i := 0; i < banyakBalita; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scanln(&beratBalita[i])
	}

	min, max := hitungMinMax(beratBalita, banyakBalita)
	avg := rerata(beratBalita, banyakBalita)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", avg)
}
