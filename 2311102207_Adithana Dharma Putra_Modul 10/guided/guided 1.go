package main
import "fmt"

func main (){
	var N int
	var berat [1000]float64

	fmt.Print("masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	for i:=0; i<N; i++{	
		fmt.Print("masukkan berat anak kelinci: ")
		fmt.Scan(&berat[i])
	}
	
	min := berat[0]
	max := berat[0]

	for i := 1; i < N; i++{
		if berat[i]<min{
			min = berat[i]
		}
		if berat[i]>max{
			max=berat[i]
		}
	}
	fmt.Printf("Berat terkecil: %.2f\n",min)
	fmt.Printf("Berat terbesar: %.2f",max)
}
