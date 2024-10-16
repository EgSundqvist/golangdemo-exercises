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
