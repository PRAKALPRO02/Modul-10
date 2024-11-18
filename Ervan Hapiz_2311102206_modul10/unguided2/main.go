package main

import "fmt"

type arrBalita [100]float64

func hitungMInMax(balita arrBalita, min, max *float64, n int) {

	*min = balita[0]
	*max = balita[0]

	for i := 0; i < n; i++ {
		if balita[i] < *min {
			*min = balita[i]
		}

		if balita[i] > *max {
			*max = balita[i]
		}
	}
	fmt.Printf("Berat Balita Minimum  : %.2f kg\n", *min)
	fmt.Printf("Berat Balita Maksimum : %.2f kg\n", *max)
}

func rerata(balita arrBalita, n int) float64 {
	sum := 0.0

	for i := 0; i < n; i++ {
		sum += balita[i]
	}

	rata_rata := sum / float64(n)
	return rata_rata
}
func main() {
	var balita arrBalita
	var min, max float64
	var n int
	fmt.Print("Masukan Banyak Data Berat Balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan Berat Balita Ke-%v: ", i+1)
		fmt.Scan(&balita[i])
	}

	hitungMInMax(balita, &min, &max, n)
	fmt.Printf("Rerata Berat balita : %.2f kg", rerata(balita, n))
}
