package main 

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah: ")
	fmt.Scan(&x, &y)

	berat := make([]float64, x)
	fmt.Println("Masukan berat tiap ikan: ")
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := (x + y - 1) / y
	totalBeratWadah := make([]float64, jumlahWadah)

	for i := 0; i < x; i++ {
		indeksWadah := i / y
		totalBeratWadah[indeksWadah] += berat[i]
	}

	fmt.Println("Total berat tiap wadah:")
	for _, total := range totalBeratWadah {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()

	fmt.Println("Rata-rata berat tiap wadah:")
	for _, total := range totalBeratWadah {
		rataRata := total / float64(y)
		fmt.Printf("%.2f ", rataRata)
	}
	fmt.Println()
}