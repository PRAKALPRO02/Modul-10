package main

import ("fmt")

func main() {
	var N int
	var berat [1000]float64

	fmt.Print("Masukan Jumlah Anak Kelinci: ")
	fmt.Scan(&N)

	fmt.Println("Masukan Berat Anak Kelinci: ")
	for i := 0; i < N; i++{
		fmt.Scan(&berat[i])
	}

	min := berat [0]
	max := berat [0]

	for i := 1; i <N; i++ {
		if berat [i] <min {
			min = berat [i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}
	fmt.Printf("Berat Kelonci: %.2f\n", min)
	fmt.Printf("Berat Terbesar: %.2f\n", max)
}