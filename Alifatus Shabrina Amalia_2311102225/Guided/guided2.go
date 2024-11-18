package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int
	fmt.Printf("Masukkan jumlah ikan dan kapasitas wadah: ")
	fmt.Scan(&x, &y)

	// Inisialisasi slice berat
	berat := make([]float64, x)
	fmt.Println("Masukkan berat tiap ikan: ")
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	// Hitung jumlah wadah dengan pembulatan ke atas
	jumlahWadah := int(math.Ceil(float64(x) / float64(y)))
	totalBeratWadah := make([]float64, jumlahWadah)

	// Distribusi berat ikan ke wadah
	for i := 0; i < x; i++ {
		indeksWadah := i / y
		totalBeratWadah[indeksWadah] += berat[i]
	}

	// Output total berat tiap wadah
	fmt.Println("Total berat tiap wadah:")
	for i, total := range totalBeratWadah {
		fmt.Printf("Wadah %d: %.2f\n", i+1, total)
	}

	// Output rata-rata berat tiap wadah
	fmt.Println("Rata-rata berat tiap wadah:")
	for i, total := range totalBeratWadah {
		// Jumlah ikan dalam wadah terakhir bisa kurang dari kapasitas maksimal
		jumlahIkan := y
		if i == jumlahWadah-1 && x%y != 0 {
			jumlahIkan = x % y
		}
		rataRata := total / float64(jumlahIkan)
		fmt.Printf("Wadah %d: %.2f\n", i+1, rataRata)
	}
}
