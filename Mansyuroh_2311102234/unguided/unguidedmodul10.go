package main

import (
	"fmt"
)

func main() {
	var jumlahBalita234 int
	var beratBalita234 []float64
	// Meminta input jumlah balita
	fmt.Print("Jumlah balita: ")
	fmt.Scanln(&jumlahBalita234)

	// Menginisialisasi array berat balita
	beratBalita234 = make([]float64, jumlahBalita234)

	// Meminta input berat masing-masing balita
	for i := 0; i < jumlahBalita234; i++ {
		fmt.Printf("Berat balita ke-%d: ", i+1)
		fmt.Scanln(&beratBalita234[i])
	}

	beratTerkecil234 := beratBalita234[0]
	beratTerbesar234 := beratBalita234[0]
	totalBerat234 := 0.0
	indeksTerkecil := 0
	indeksTerbesar := 0

	for i, berat := range beratBalita234 {
		if berat < beratTerkecil234 {
			beratTerkecil234 = berat
			indeksTerkecil = i
		}
		if berat > beratTerbesar234 {
			beratTerbesar234 = berat
			indeksTerbesar = i
		}
		totalBerat234 += berat
	}

	rataRata234 := totalBerat234 / float64(jumlahBalita234)

	// Menampilkan hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg (balita ke-%d)\n", beratTerkecil234, indeksTerkecil+1)
	fmt.Printf("Berat balita maksimum: %.2f kg (balita ke-%d)\n", beratTerbesar234, indeksTerbesar+1)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata234)
}
