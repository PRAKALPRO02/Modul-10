package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if *bMax < arrBerat[i] {
			*bMax = arrBerat[i]
		}

		if *bMin > arrBerat[i] {
			*bMin = arrBerat[i]
		}
	}

}

func rerata(arrBerat arrBalita, n int) float64 {
	var sum float64
	for i := 0; i < n; i++ {
		sum += arrBerat[i]
	}
	return sum / float64(n)
}

func main() {
	var dataBalita arrBalita
	var n int
	var max, min float64
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print("Masukan berat balita ke-", i+1, ": ")
		fmt.Scanln(&dataBalita[i])
	}

	hitungMinMax(dataBalita, n, &min, &max)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata(dataBalita, n))
}
