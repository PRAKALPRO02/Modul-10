package main

import "fmt"

const maxEntries = 100

type weights [maxEntries]float64

func findMinMax(data weights, count int) {
	minVal := data[0]
	maxVal := data[0]

	for i := 1; i < count; i++ {
		if data[i] < minVal {
			minVal = data[i]
		}
		if data[i] > maxVal {
			maxVal = data[i]
		}
	}

	fmt.Printf("Berat minimum balita: %.2f\n", minVal)
	fmt.Printf("Berat maksimum balita: %.2f\n", maxVal)
}

func calculateAverage(data weights, count int) float64 {
	var total float64
	for i := 0; i < count; i++ {
		total += data[i]
	}
	return total / float64(count)
}

func main() {
	var data weights
	var count int

	fmt.Print("Masukkan jumlah data berat balita: ")
	fmt.Scan(&count)

	fmt.Println("Masukkan berat masing-masing balita:")
	for i := 0; i < count; i++ {
		fmt.Scan(&data[i])
	}

	findMinMax(data, count)
	fmt.Printf("Rata-rata berat balita: %.2f\n", calculateAverage(data, count))
}
