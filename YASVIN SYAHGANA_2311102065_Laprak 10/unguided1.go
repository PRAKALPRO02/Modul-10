package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin *float64, bMax *float64, n int) {
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			fmt.Print(arrBerat[i])
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var arrBerat arrBalita
	fmt.Print("Masukkan banyak data berat balita : ")
	fmt.Scan(&n)

	if n > 100 {
		fmt.Print("Ukuran terlalu besar. Harus kurang dari atau sama dengan 100.")
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("Masukkan berat balita ke-%v: ", i)
		fmt.Scan(&arrBerat[i-1])
	}

	var min, max float64 = arrBerat[0], arrBerat[0]
	hitungMinMax(arrBerat, &min, &max, n)
	fmt.Printf("Berat balita minimum: %.2f KG\n", min)
	fmt.Printf("Berat balita maksimum: %.2f KG\n", max)
	fmt.Printf("Rerata berat balita: %.2f KG\n", rerata(arrBerat, n))
}
