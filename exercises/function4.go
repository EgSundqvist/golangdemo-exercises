package exercises

import (
	"errors"
	"fmt"
)

func CalculateTaxsOnSalary(salary float64) (float64, string, error) {
	if salary < 0 {
		return 0, "", errors.New("negative salary, really?")
	}

	var taxRate float64
	var taxDesc string

	if salary < 15000 {
		taxRate = 0.12
		taxDesc = "12%"
	} else if salary >= 15000 && salary < 30000 {
		taxRate = 0.28
		taxDesc = "28%"
	} else {
		taxRate = 0.50
		taxDesc = "50%"
	}

	tax := salary * taxRate
	return tax, taxDesc, nil

}

func Functions4() {
	var salary float64

	fmt.Println("Enter your salary: ")
	fmt.Scanln(&salary)

	tax, taxDesc, err := CalculateTaxsOnSalary(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The tax on your salary is %.2f (%s)\n", tax, taxDesc)
	}

}
