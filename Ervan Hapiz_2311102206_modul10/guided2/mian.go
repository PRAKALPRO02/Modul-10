package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukan jumlah ikan dan kapasitas wadah : ")
	fmt.Scan(&x, &y)

	Berat := make([]float64, x)
	fmt.Print("Masukkan berat tiap ikan : ")
	for i := 0; i < x; i++ {
		fmt.Scan(&Berat[i])
	}

	jumlahWadah := (x + y - 1) / y
	totalBeratWadah := make([]float64, jumlahWadah)

	for i := 0; i < x; i++ {
		indekWadah := i / y
		totalBeratWadah[indekWadah] += Berat[i]
	}

	fmt.Println("Total Berat tiap wadah : ")
	for _, total := range totalBeratWadah {
		fmt.Printf("%.2f", total)
	}

	fmt.Println()

	fmt.Println("Rata rata berat tiap wadah : ")
	for _, total := range totalBeratWadah {
		ratarata := total / float64(y)
		fmt.Printf("%.2f", ratarata)
	}
	fmt.Println()
}
