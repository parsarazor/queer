package main

import "fmt"

func main() {
	var (
		w      int     = 0   // w is income of shaban for a day
		sh     int     = 0   // sh is the price of a sheep
		p      float64 = 0.0 // p is the probability of a sheep being eaten by a wolf in a day
		income float64 = 0.0 // total income of shaban in a year could be: 0 or negative value
	)
	fmt.Scan(&w, &sh)
	fmt.Scan(&p)

	income = float64(w*365) - ( 365.0 * p * float64(sh))

	fmt.Printf("%.2f\n", income)

}
