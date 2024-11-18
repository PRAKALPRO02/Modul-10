package main

import (
	"fmt"
)

func main() {
	var N int
	var berat [1000]float64

	fmt.Print("masukkan jumlah anak kelinci :")
	fmt.Scan(&N)

	fmt.Println("masukkan berat anak kelinci :")
	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}

	// inisialisasi nilai min dan max dg elemen pertama
	min := berat[0]
	max := berat[0]

	for i := 1; i < N; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}
	fmt.Printf(" berat terkecil : %.2f\n", min)
	fmt.Printf(" berat terbesar : %.2f\n", max)
}
