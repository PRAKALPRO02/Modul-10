package main

import "fmt"



func main() {
	var N int
	var berat [1000]float64

	fmt.Print("\n\nMasukkan jumlah anak kelici : ")
	fmt.Scan(&N)

	fmt.Print("Masukkan berat anak kelinci: ")
	for i := 0; i < N; i++ {
	fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 0; i < N; i++ {
		if min > berat[i] {
			min = berat[i]

		}
		if max < berat[i] {
			max = berat[i]
		}
	}
	fmt.Printf("\nBerat terkecil : %.2f\n", min)
	fmt.Printf("Berat terbesar : %.2f\n", max)
}


