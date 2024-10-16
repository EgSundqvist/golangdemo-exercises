package exercises

import (
	"fmt"
)

func CalculateVAT(price float64) float64 {
	return price * 0.25
}

func Functions2() {
	var price float64

	fmt.Println("Enter the price: ")
	fmt.Scanln(&price)

	vat := CalculateVAT(price)
	fmt.Println("The VAT is: ", vat)

}
