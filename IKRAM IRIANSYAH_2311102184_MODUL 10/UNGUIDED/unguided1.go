package main

import "fmt"

type arrBalita [100]float64

func minmax(arrBerat arrBalita, bMin, bMax *float64) {
    *bMin = arrBerat[0]
    *bMax = arrBerat[0]
    
    for i := 1; i < len(arrBerat); i++ {
        if arrBerat[i] != 0 {
            if arrBerat[i] < *bMin {
                *bMin = arrBerat[i]
            }
            if arrBerat[i] > *bMax {
                *bMax = arrBerat[i]
            }
        }
    }
}

func ratarata(arrBerat arrBalita, n int) float64 {
    var total float64
    for i := 0; i < n; i++ {
        total += arrBerat[i]
    }
    return total / float64(n)
}

func main() {
    var balita arrBalita
    var n int
    var min, max float64
    
    fmt.Print("Masukkan jumlah balita : ")
    fmt.Scan(&n)
    
    for i := 0; i < n; i++ {
        fmt.Printf("Masukkan berat balita ke-%d : ", i+1)
        fmt.Scan(&balita[i])
    }
    
    minmax(balita, &min, &max)
    
    rerata := ratarata(balita, n)
    
    fmt.Printf("\n= Hasil Analisis Data Berat Balita =\n")
    fmt.Printf("Berat Terkecil : %.2f kg\n", min)
    fmt.Printf("Berat Terbesar : %.2f kg\n", max)
    fmt.Printf("Berat Rata-rata : %.2f kg\n", rerata)
}