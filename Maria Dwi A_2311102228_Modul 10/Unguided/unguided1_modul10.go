package main

import "fmt"

type arrBalita [100]float64

func cariMin(N int, berat *arrBalita) float64{
	min := berat[0]

	for i := 0; i < N; i++ {
		if min > berat[i]{
			min = berat[i]
		}
	}
	return min
}

func cariMax(N int, berat *arrBalita)float64 {
	max := berat[0]

	for i := 0; i < N; i++ {
		if max < berat[i] {
			max = berat[i]
		}
	}
	return max
}

func rataRata(N int, berat *arrBalita) float64 {
	totalBerat := 0.0

	for i := 0; i < N; i++ {
		totalBerat += berat[i]
	}
	return totalBerat / float64(N)
}

func main() {

	var jumlah int
	var berat arrBalita

	fmt.Print("\n~Program Menyimpan Data Posyandu~")
	fmt.Print("\n\nMasukkan jumlah balita : ")
	fmt.Scan(&jumlah)

	for i := 0; i < jumlah; i++ {
	fmt.Print("Masukkan berat balita ke- ", i+1,  ": ")
	fmt.Scan(&berat[i])
	}

	minimal := cariMin(jumlah, &berat)
	maksimal := cariMax(jumlah, &berat)
	rata := rataRata(jumlah, &berat)

	fmt.Printf("\nBerat balita maksimum : %.2f kg\n", minimal)
	fmt.Printf("Berat balita maksimum : %.2f kg\n", maksimal)
	fmt.Printf("Rata-rata berat balita : %.2f kg\n\n", rata)



}