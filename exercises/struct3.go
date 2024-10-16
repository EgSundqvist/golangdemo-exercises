package exercises

import (
	"fmt"

	"github.com/EgSundqvist/golangdemo-exercises/utils"
)

type Dog struct {
	Name   string
	Breed  string
	Age    int
	Weight float64
}

var kennel []Dog

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
			utils.ClearTerminal()
			registerDog()
			utils.Pause()
		case 2:
			utils.ClearTerminal()
			listDogs()
			utils.Pause()
		case 3:
			utils.ClearTerminal()
			deleteDog()
			utils.Pause()
		case 4:
			utils.ClearTerminal()
			fmt.Println("Exiting Program.")
			utils.Pause()
			return
		default:
			fmt.Println("Invalid command.")
		}
	}
}
