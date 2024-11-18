package main

import "fmt"

func main() {
	var Data [1000]float64
	var n int

	fmt.Print("Masukan jumlah anak kelinci : ")
	fmt.Scan(&n)

	fmt.Print("Masukna berat anak kelinci : ")
	for i := 0; i < n; i++ {
		fmt.Scan(&Data[i])
	}
	min := Data[0]
	max := Data[0]

	for i := 0; i < n; i++ {
		if Data[i] < min {
			min = Data[i]
		}

		if Data[i] > max {
			max = Data[i]
		}
	}

	fmt.Printf("Berat kelinci terkecil adalah %.2f dan terbesar adalah %.2f", min, max)
}
