package main

import "fmt"

const maxData = 100

type datas [maxData]float64

func minMax(data datas, n int) (float64, float64) {
	min := data[0]
	max := data[0]

	for i := 1; i < n; i++ {
		if data[i] < min {
			min = data[i]
		}
		if data[i] > max {
			max = data[i]
		}
	}

	return min, max
}


func rataRata(data datas, n int) float64 {
	var sum float64

	for i := 0; i < n; i++ {
		sum += data[i]
	}

	return sum / float64(n)
}

func main() {
	var data datas
	var n int

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	min, max := minMax(data, n)
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata(data, n))
}
