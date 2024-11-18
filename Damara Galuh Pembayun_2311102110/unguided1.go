package main

import "fmt"

// arrBalita adalah tipe data array dengan ukuran maksimal 100 elemen,
// yang menyimpan nilai berat balita dalam tipe float64.
type arrBalita [100]float64

// hitungMinMax menghitung nilai minimum dan maksimum dari array berat balita.
// Parameter:
//   arrBerat: array yang berisi data berat balita.
//   bMin: pointer ke variabel yang akan menyimpan nilai minimum.
//   bMax: pointer ke variabel yang akan menyimpan nilai maksimum.
func hitungMinMax(arrBerat arrBalita, bMin *float64, bMax *float64) {
    *bMin = arrBerat[0]
    *bMax = arrBerat[0]

    for _, berat := range arrBerat {
        if berat < *bMin {
            *bMin = berat
        }
        if berat > *bMax {
            *bMax = berat
        }
    }
}

// rerata menghitung nilai rata-rata dari array berat balita.
// Parameter:
//   arrBerat: array yang berisi data berat balita.
// Return:
//   nilai rata-rata dari array berat balita.
func rerata(arrBerat arrBalita) float64 {
    var total float64
    for _, berat := range arrBerat {
        total += berat
    }
    return total / float64(len(arrBerat))
}

func main() {
    var n int
    fmt.Print("Masukkan banyak data berat balita: ")
    fmt.Scan(&n)

    var beratBalita arrBalita
    for i := 0; i < n; i++ {
        fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
        fmt.Scan(&beratBalita[i])
    }

    var min, max float64
    hitungMinMax(beratBalita, &min, &max)

    rata := rerata(beratBalita)

    fmt.Printf("Berat balita minimum: %.2f kg\n", min)
    fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
    fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}