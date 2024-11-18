package main

import "fmt"

const maxData = 100

type arrBalita [maxData]float64

var jml_balita int

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rataRata(arrBerat arrBalita, n int) float64 {
	var jumlah float64
	for i := 0; i < n; i++ {
		jumlah += arrBerat[i]
	}
	return jumlah / float64(n)
}

func main() {
	var balitas arrBalita
	var min, max float64

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&jml_balita)

	for i := 0; i < jml_balita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&balitas[i])
	}

	hitungMinMax(balitas, jml_balita, &min, &max)
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata(balitas, jml_balita))
}
