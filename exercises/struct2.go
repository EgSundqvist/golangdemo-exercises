package exercises

import (
	"fmt"
)

type Person struct {
	BirthDate     string
	MiddleName    string
	Name          string
	StreetAddress string
	PostalCode    string
	City          string
}

// Person methods
func (p *Person) AddMiddleName(middlename string) {
	p.MiddleName = middlename
}

func (p *Person) MoveAdress(streetAddress, postalCode, city string) {
	p.StreetAddress = streetAddress
	p.PostalCode = postalCode
	p.City = city
}

func (p *Person) MoveTo(otherPerson Person) {
	p.MoveAdress(otherPerson.StreetAddress, otherPerson.PostalCode, otherPerson.City)
}

func Structs2() {
	person1 := Person{
		BirthDate:     "1990-01-01",
		Name:          "Petter Niklas",
		MiddleName:    "",
		StreetAddress: "Broadway 1",
		PostalCode:    "12345",
		City:          "Rågsved",
	}

	person2 := Person{
		BirthDate:     "1992-02-02",
		Name:          "Christer Pettersson",
		MiddleName:    "",
		StreetAddress: "Lortgränd 2",
		PostalCode:    "67890",
		City:          "Bandhagen",
	}

	var name1, name2 string
	fmt.Println("Add middle name for person 1:")
	fmt.Scanln(&name1)
	person1.AddMiddleName(name1)

	fmt.Println("Add middle name for person 2:")
	fmt.Scanln(&name2)
	person2.AddMiddleName(name2)

	fmt.Printf("Person 1: %s, Middle name: %s, Birthdate: %s, Adress: %s, %s %s\n", person1.Name, person1.MiddleName, person1.BirthDate, person1.StreetAddress, person1.PostalCode, person1.City)
	fmt.Printf("Person 2: %s, Middle name: %s, Birthdate: %s, Adress: %s, %s %s\n", person2.Name, person2.MiddleName, person2.BirthDate, person2.StreetAddress, person2.PostalCode, person2.City)

	var moveChoice int
	fmt.Println("Who is moving, person (1 or 2):")
	fmt.Scanln(&moveChoice)

	if moveChoice == 1 {
		person1.MoveTo(person2)
	} else if moveChoice == 2 {
		person2.MoveTo(person1)
	} else {
		fmt.Println("Invalid choice")
		return
	}

	fmt.Println("\nAfter the move:")
	fmt.Printf("Person 1: %s, Birthdate: %s, Adress: %s, %s %s\n", person1.Name, person1.BirthDate, person1.StreetAddress, person1.PostalCode, person1.City)
	fmt.Printf("Person 2: %s, Birthdate: %s, Adress: %s, %s %s\n", person2.Name, person2.BirthDate, person2.StreetAddress, person2.PostalCode, person2.City)
}
