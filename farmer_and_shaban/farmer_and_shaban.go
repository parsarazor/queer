package main
import "fmt"

func main(){
	var(
		w int = 0
		sh int = 0
		p float64 = 0.0
		s float64 = 0.0
		income float64 = 0.0
	)
	fmt.Scan(&w,&sh)
	fmt.Scan(&p)

	s = 365.0 / (1.0 / p)
	income = float64(w * 365) - (s * float64(sh))

	fmt.Printf("%.2f\n", income)
	

}
