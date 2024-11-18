package main

import "fmt"

type arrBalita [100]float64

// Fungsi untuk mencari berat minimum
func cariMin(N int, berat *arrBalita) (min float64) {
	min = berat[0]
	for i := 1; i < N; i++ { // Mulai dari indeks 1 karena min sudah diinisialisasi dengan berat[0]
		if berat[i] < min {
			min = berat[i]
		}
	}
	return
}

// Fungsi untuk mencari berat maksimum
func cariMax(N int, berat *arrBalita) (max float64) {
	max = berat[0]
	for i := 1; i < N; i++ { // Mulai dari indeks 1
		if berat[i] > max {
			max = berat[i]
		}
	}
	return
}

// Fungsi untuk menghitung rata-rata berat
func rataRata(N int, berat *arrBalita) (rata float64) {
	var total float64
	for i := 0; i < N; i++ {
		total += berat[i]
	}
	rata = total / float64(N)
	return
}

func main() {
	var jumlah int
	var berat arrBalita

	fmt.Println("\n~ Program Menyimpan Data Posyandu ~")
	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&jumlah)

	// Memasukkan data berat balita
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	// Pemanggilan fungsi
	minimal := cariMin(jumlah, &berat)
	maksimal := cariMax(jumlah, &berat)
	rata := rataRata(jumlah, &berat)

	// Output hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", minimal)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", maksimal)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n\n", rata)
}

