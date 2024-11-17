package main

import "fmt"

const maxData = 100

type datas [maxData] float64

func minMax(data datas, n int) {
	min:= data[0]
	for i:=1; i < n; i++ {
		if min > data[i] {
			min = data[i]
		}

	}

	max := data[0]
	for i:=1; i < n; i++ {
		if max < data[i] {
			max = data[i]
		}

	}

	fmt.Printf("berat balita minimum: %.2f\n",min)
	fmt.Printf("berat balita maximum: %.2f\n",max)

}

func rataRata(data datas, n int) float64 {
	var sum,hasil float64

	for i:=0; i < n; i++ {
		sum += data[i]
	}
	hasil = float64(sum) / float64(n)

	return hasil
}

func main() {
	var data datas
	var n int

	fmt.Scan(&n)

	for i:=0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	minMax(data,n)
	fmt.Printf("%.2f",rataRata(data,n))
}