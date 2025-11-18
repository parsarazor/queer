package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func tarkibat(coats, shirts, pants, caps, jackets []string, season string) {
	season = strings.ToLower(season)

	switch season { // switch case for different seasons combinations

	case "spring", "fall":
		// Combinations without coat/jacket
		for _, shirt := range shirts {
			for _, pant := range pants {
				fmt.Printf("SHIRT: %s PANT: %s\n", shirt, pant) // Added \n for clarity
			}
		}
		// Combinations with cap
		for _, shirt := range shirts {
			for _, pant := range pants {
				for _, cap := range caps {
					fmt.Printf("SHIRT: %s PANT: %s CAP: %s\n", shirt, pant, cap)
				}
			}
		}

		// Combinations with coat (excluding specific colors in fall)
		for _, coat := range coats {
			if season == "fall" && (coat == "yellow" || coat == "orange") {
				continue
			}
			for _, shirt := range shirts {
				for _, pant := range pants {
					fmt.Printf("COAT: %s SHIRT: %s PANT: %s\n", coat, shirt, pant)
				}
			}
		}

		// Combinations with coat and cap (excluding specific colors in fall)
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
		// Combinations with cap (no coat/jacket/jacket in summer)
		for _, shirt := range shirts {
			for _, pant := range pants {
				for _, cap := range caps {
					fmt.Printf("SHIRT: %s PANT: %s CAP: %s\n", shirt, pant, cap)
				}
			}
		}

	case "winter":
		// Combinations with coats
		for _, coat := range coats {
			for _, shirt := range shirts {
				for _, pant := range pants {
					fmt.Printf("COAT: %s SHIRT: %s PANT: %s\n", coat, shirt, pant)
				}
			}
		}
		// Combinations with jackets
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

	// Read the full line for each category
	scanner.Scan()
	coatInputLine := scanner.Text()
	scanner.Scan()
	shirtInputLine := scanner.Text()
	scanner.Scan()
	pantsInputLine := scanner.Text()
	scanner.Scan()
	capInputLine := scanner.Text()
	scanner.Scan()
	jacketInputLine := scanner.Text()
	scanner.Scan()
	season := scanner.Text()


	coatParts := strings.SplitN(coatInputLine, ": ", 2)
	coatInput := ""
	if len(coatParts) > 1 {
		coatInput = coatParts[1] // Take the part after "PREFIX: "
	}
	coats := strings.Fields(coatInput) // Split the remaining values

	shirtParts := strings.SplitN(shirtInputLine, ": ", 2)
	shirtInput := ""
	if len(shirtParts) > 1 {
		shirtInput = shirtParts[1]
	}
	shirts := strings.Fields(shirtInput)

	pantsParts := strings.SplitN(pantsInputLine, ": ", 2)
	pantsInput := ""
	if len(pantsParts) > 1 {
		pantsInput = pantsParts[1]
	}
	pants := strings.Fields(pantsInput)

	capParts := strings.SplitN(capInputLine, ": ", 2)
	capInput := ""
	if len(capParts) > 1 {
		capInput = capParts[1]
	}
	caps := strings.Fields(capInput)

	jacketParts := strings.SplitN(jacketInputLine, ": ", 2)
	jacketInput := ""
	if len(jacketParts) > 1 {
		jacketInput = jacketParts[1]
	}
	jackets := strings.Fields(jacketInput)

	tarkibat(coats, shirts, pants, caps, jackets, season)
}
