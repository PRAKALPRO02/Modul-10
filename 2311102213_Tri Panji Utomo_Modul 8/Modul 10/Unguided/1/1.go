package main

import "fmt"

type arrBalita [100]float64

func MinMax(arrBerat arrBalita, n int) (float64, float64) {
    min := arrBerat[0]
    max := arrBerat[0]

    for i := 1; i < n; i++ {
        if arrBerat[i] < min {
            min = arrBerat[i]
        }
        if arrBerat[i] > max {
            max = arrBerat[i]
        }
    }
    return min, max
}

func rerata(arrBerat arrBalita, n int) float64 {
    var total213 float64 = 0

    for i := 0; i < n; i++ {
        total213 += arrBerat[i]
    }
    return total213 / float64(n)
}

func main() {
    var banyak213 int
    var berat213 arrBalita

    fmt.Print("Masukkan banyak data balita: ")
    fmt.Scanln(&banyak213)

    for i := 0; i < banyak213; i++ {
        fmt.Printf("Masukan berat balita ke-%d: ", i+1)
        fmt.Scanln(&berat213[i])
    }

    min, max := MinMax(berat213, banyak213)
    avg := rerata(berat213, banyak213)

    fmt.Printf("Berat balita minimum: %.2f\n", min)
    fmt.Printf("Berat balita maximum: %.2f\n", max)
    fmt.Printf("Rerata berat balita: %.2f\n", avg)
}
