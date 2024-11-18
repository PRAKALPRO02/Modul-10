package main
import "fmt"

func hitungMinMax(arrBerat207[100]float64, bMin float64, bMax float64,N int){
		for i := 0; i < N; i++{
		if arrBerat207[i]<bMin{
			bMin = arrBerat207[i]
		}
		if arrBerat207[i]>bMax{
			bMax=arrBerat207[i]
		}
	}
	fmt.Printf("Berat Minumum: %.2f\n",bMin)
	fmt.Printf("Berat Maksimum: %.2f\n",bMax)
	fmt.Printf("Rerata berat balita: %.2f",rerata(arrBerat207,N))
}

func rerata(arrBerat207[100]float64,N int) float64{
	var jumlah float64
	for i:=0;i<N;i++{
		jumlah += arrBerat207[i]
	}
	BanyakData := float64(N)
	return jumlah/BanyakData
}
func main (){
	var N int
	var arrBerat207 [100]float64

	fmt.Print("Masukkan Banyak data berat Balita: ")
	fmt.Scan(&N)

	for i:=0; i<N; i++{
	fmt.Printf("Masukkan berat Balita ke-%d: ",i+1)		
	fmt.Scan(&arrBerat207[i])
	}

	bMin := arrBerat207[0]
	bMax := arrBerat207[0]
	hitungMinMax(arrBerat207,bMin,bMax,N)
}
