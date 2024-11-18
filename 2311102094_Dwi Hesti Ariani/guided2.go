package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("masukkan jumlah ikan dan kapasitas wadah : ")
	fmt.Scan(&x, &y)

	berat := make([]float64, x)
	fmt.Println("masukkan berat tiap ikan: ")
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := (x + y - 1) / y
	totalBeratWadah := make([]float64, jumlahWadah)

	for i := 0; i < x; i++ {
		indeksWadah := i / y
		totalBeratWadah[indeksWadah] += berat[i]
	}
	//output total berat tiap wadah
	fmt.Println("total berat tiap wadah :")
	for _, total := range totalBeratWadah {
		fmt.Printf("%.2f", total)
	}
	fmt.Println()

	// output rata2
	fmt.Println("rata -rata berat tiap wadah: ")
	for _, total := range totalBeratWadah {
		rataRata := total / float64(y)
		fmt.Printf("%.2f", rataRata)
	}
	fmt.Println()
}