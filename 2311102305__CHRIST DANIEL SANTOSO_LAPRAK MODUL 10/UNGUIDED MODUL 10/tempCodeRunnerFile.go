package main

		// NAMA : CHRIST DANIEL SANTOSO
		// NIM : 2311102305
		// KELAS :IF-11-02

import "fmt"

const maxData = 100

type weightData [maxData]float64

func getMinMax(weights weightData, n int) {
	min := weights[0]
	max := weights[0]

	for j := 1; j < n; j++ {
		if weights[j] < min {
			min = weights[j]
		}
		if weights[j] > max {
			max = weights[j]
		}
	}

	fmt.Printf("Berat minimum balita: %.2f\n", min)
	fmt.Printf("Berat maksimum balita: %.2f\n", max)
}

func getAverage(weights weightData, n int) float64 {
	var sum float64
	for j := 0; j < n; j++ {
		sum += weights[j]
	}
	return sum / float64(n)
}

func main() {
	var weights weightData
	var n int

	fmt.Print("Masukkan jumlah data berat balita: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan berat masing-masing balita:")
	for j := 0; j < n; j++ {
		fmt.Scan(&weights[j])
	}

	getMinMax(weights, n)
	fmt.Printf("Rata-rata berat balita: %.2f\n", getAverage(weights, n))
}
