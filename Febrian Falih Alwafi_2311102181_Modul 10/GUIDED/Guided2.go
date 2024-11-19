package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah: ")
	fmt.Scan(&x, &y)

	berat := make([]float64, x)
	fmt.Println("Masukkan berat tiap ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := (x + y - 1) / y
	totalBeratWadah := make([]float64, jumlahWadah)
	jumlahIkanDiWadah := make([]int, jumlahWadah)

	for i := 0; i < x; i++ {
		indeksWadah := i / y
		totalBeratWadah[indeksWadah] += berat[i]
		jumlahIkanDiWadah[indeksWadah]++
	}

	// Output total berat tiap wadah
	fmt.Println("Total berat tiap wadah:")
	for _, total := range totalBeratWadah {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()

	// Output rata-rata berat tiap wadah
	fmt.Println("Rata-rata berat tiap wadah:")
	for i := 0; i < jumlahWadah; i++ {
		if jumlahIkanDiWadah[i] > 0 {
			rataRata := totalBeratWadah[i] / float64(jumlahIkanDiWadah[i])
			fmt.Printf("%.2f ", rataRata)
		} else {
			fmt.Printf("0.00 ") // jika wadah kosong
		}
	}
	fmt.Println()
}