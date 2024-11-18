package main

import "fmt"

func hitungKelipatanEmpat2311102234() int {
	var angka, total int
	for {
		fmt.Print("Masukkan bilangan(negatif untuk berhenti): ")
		fmt.Scan(&angka)

		if angka < 0 {
			break
		}

		if angka > 0 && angka%4 == 0 {
			total += angka
		}
	}
	return total
}

func main() {
	hasil := hitungKelipatanEmpat2311102234()
	fmt.Printf("Jumlah bilangan kelipatan 4 yaitu: %d\n", hasil)
}
