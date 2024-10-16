package exercises

import (
	"fmt"
)

type Dish struct {
	Name     string
	Price    float64
	Type     string
	Calories int
}

type Person struct {
	BirthDate     string
	MiddleName    string
	Name          string
	StreetAddress string
	PostalCode    string
	City          string
}

type Dog struct {
	Name   string
	Breed  string
	Age    int
	Weight float64
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

var kennel []Dog

// Dog methods
func registerDog() {
	var name, breed string
	var age int
	var weight float64

	fmt.Println("Enter dog name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter dog breed: ")
	fmt.Scanln(&breed)
	fmt.Println("Enter dog age (in years): ")
	fmt.Scanln(&age)
	fmt.Println("Enter dog weight (in kg): ")
	fmt.Scanln(&weight)

	dog := Dog{
		Name:   name,
		Breed:  breed,
		Age:    age,
		Weight: weight,
	}

	kennel = append(kennel, dog)
	fmt.Println("Dog has been registred.")
}

func calculateTailLength(dog Dog) float64 {
	if dog.Breed == "Tax" {
		return 3.7
	}
	return float64(dog.Age) * dog.Weight / 10
}

func listDogs() {
	var minTailLength float64
	fmt.Println("Enter minimum tale length: ")
	fmt.Scanln(&minTailLength)

	for _, dog := range kennel {
		tailLength := calculateTailLength(dog)
		if tailLength >= minTailLength {
			fmt.Printf("%s %s %d age %g kg tale=%g\n", dog.Name, dog.Breed, dog.Age, dog.Weight, tailLength)
		}
	}
}

func deleteDog() {
	var name string
	fmt.Println("Enter the name of the dog to be deleted: ")
	fmt.Scanln(&name)

	for i, dog := range kennel {
		if dog.Name == name {
			kennel = append(kennel[:i], kennel[i+1:]...)
			fmt.Printf("Dog with the name %s is deleted.\n", name)
			return
		}
	}
	fmt.Printf("Dog with the name %s could not be found.\n", name)
}

func Structs1() {
	dishes := []Dish{
		{"Pizza", 100, "Yummy", 800},
		{"Salad", 120, "Veg", 300},
		{"Burger", 90, "Meat", 700},
	}
	for _, dish := range dishes {
		fmt.Printf("Name: %s, Price: %.2f, Type: %s, Calories: %d\n", dish.Name, dish.Price, dish.Type, dish.Calories)
	}
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

func Structs3() {
	for {
		var command int
		fmt.Println("Choose a command:")
		fmt.Println("1. Register a dog")
		fmt.Println("2. List all dogs")
		fmt.Println("3. Delete a dog")
		fmt.Println("4. Exit")
		fmt.Scanln(&command)

		switch command {
		case 1:
			registerDog()
		case 2:
			listDogs()
		case 3:
			deleteDog()
		case 4:
			fmt.Println("Exiting Program.")
			return
		default:
			fmt.Println("Invalid command.")
		}
	}
}
