package main

import "fmt"

type arrBalita [100]float64

func main() {

	var banyakData int
	var arrBerat arrBalita
	var bMax, bMin float64

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&banyakData)

	for i := 0; i < banyakData; i++ {
		fmt.Printf("Masukkan berat balita ke-%d : ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	hitungMinMax(arrBerat, banyakData, &bMin, &bMax)
	rata := rerata(arrBerat, banyakData)

	fmt.Printf("\nBerat balita terkecil: %.2f kg\n", bMin)
	fmt.Printf("Berat balita terbesar: %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rata)

}

func hitungMinMax(arrBerat arrBalita, banyakData int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < banyakData; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}

	}
}

func rerata(arrBerat arrBalita, banyakData int) float64 {
	var totalBerat float64

	for i := 0; i < banyakData; i++ {
		totalBerat += arrBerat[i]
	}
	return totalBerat / float64(banyakData)

}
