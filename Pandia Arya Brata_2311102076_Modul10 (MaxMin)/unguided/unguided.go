package main

import (
	"fmt"
)

type arrBalita [100]float64

func balitaminmax(arrBerat arrBalita, n int, min, max *float64) {
	*min = arrBerat[0]
	*max = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *min {
			*min = arrBerat[i]
		}
		if arrBerat[i] > *max {
			*max = arrBerat[i]
		}
	}
}

func brata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var arr arrBalita
	var n int
	var min, max float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scanln(&arr[i])
	}

	balitaminmax(arr, n, &min, &max)

	avg := brata(arr, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rata' berat balita: %.2f kg\n", avg)
}
