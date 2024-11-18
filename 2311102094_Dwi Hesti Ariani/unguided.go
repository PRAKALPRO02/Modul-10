package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	/* I.S. Terdefinisi array dinamis arrBerat dengan panjang n
	   F.S. Menampilkan berat minimum dan maksimum balita */
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < n; i++ { // Hanya proses hingga jumlah input n
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
	fmt.Printf("Berat balita minimum: %.2f kg\n", *bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", *bMax)
}

func rerata(arrBerat arrBalita, n int) float64 {
	/* Menghitung dan mengembalikan rerata berat balita sesuai jumlah input n */
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var arrBalita arrBalita
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scanln(&arrBalita[i])
	}
	var bMin, bMax float64
	hitungMinMax(arrBalita, n, &bMin, &bMax)
	rerataBalita := rerata(arrBalita, n)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerataBalita)
}
