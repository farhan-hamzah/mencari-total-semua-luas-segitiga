package main
import "fmt"
func luasSegitigas(alas, tinggi int, luas *float64){
	*luas = (0.5)*(float64(alas*tinggi))
	fmt.Println(*luas)
}

func main(){
	var a, t int
	var hasil, totalLuas float64
	var i int
	totalLuas = 0
	fmt.Scan(&i)
	for i >= 1{
		fmt.Scan(&a, &t)
		luasSegitigas(a, t, &hasil)
		totalLuas += (hasil)
		i--
	}
	fmt.Print(totalLuas)
}