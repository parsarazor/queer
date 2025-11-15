package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func tarkibat(coats, shirts, pants, caps, jackets []string, season string) {
	season = strings.ToLower(season)

	switch season {		// switch case for different seasons combinations

	case "spring", "fall":
		for _, shirt := range shirts {
			for _, pant := range pants {
				fmt.Printf("SHIRT: %s PANT: %s ", shirt, pant)
			}

		}
		for _, shirt := range shirts {
			for _, pant := range pants {
				for _, cap := range caps {
					fmt.Printf("SHIRT: %s PANT: %s CAP: %s\n", shirt, pant, cap)
				}
			}
		}
		
		for _, coat := range coats {
			if season == "fall" && (coat == "yellow" || coat == "orange") {
				continue
			}
			for _, shirt := range shirts {
				for _, pant := range pants {
					fmt.Printf("COAT: %s SHIRT: %s PANTS: %s\n", coat, shirt, pant)
				}
			}
		}

		for _, coat := range coats {
			if season == "fall" && (coat == "yellow" || coat == "orange") {
				continue
			}
			for _, shirt := range shirts {
				for _, pant := range pants {
					for _, cap := range caps {
						fmt.Printf("COAT: %s SHIRT: %s PANT: %s CAP: %s\n", coat, shirt, pant, cap)
					}
				}
			}
		}
		

	case "summer":
		for _, shirt := range shirts {
			for _, pant := range pants {
				for _, cap := range caps {
					fmt.Printf("SHIRT: %s PANT: %s CAP: %s\n", shirt, pant, cap)
				}
			}
		}
	

	case "winter":
		for _, coat := range coats {
			for _, shirt := range shirts {
				for _, pant := range pants {
					fmt.Printf("COAT: %s SHIRT: %s PANT: %s\n", coat, shirt, pant)
				}
			}
		}
		for _, jacket := range jackets {
			for _, shirt := range shirts {
				for _, pant := range pants {
				
					fmt.Printf("JACKET: %s SHIRT: %s PANT: %s\n", jacket, shirt, pant)
				}
			} 
		}

	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
		
		fmt.Print("COAT: ")
		scanner.Scan()
		coatInput := scanner.Text()
		
		fmt.Print("SHIRT: ")
		scanner.Scan()
		shirtInput := scanner.Text()
		
		fmt.Print("PANTS: ")
		scanner.Scan()
		pantsInput := scanner.Text()
		
		fmt.Print("CAP: ")
		scanner.Scan()
		capInput := scanner.Text()
		
		fmt.Print("JACKET: ")
		scanner.Scan()
		jacketInput := scanner.Text()
		
		scanner.Scan()
		season := scanner.Text()
		
		coats := strings.Fields(coatInput)
		shirts := strings.Fields(shirtInput)
		pants := strings.Fields(pantsInput)
		caps := strings.Fields(capInput)
		jackets := strings.Fields(jacketInput)
		
		tarkibat(coats, shirts, pants, caps, jackets, season)
	

}
