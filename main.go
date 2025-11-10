package main

import (
	"fmt"
	"hello_world/types"
)

func main() {
	product1 := types.AddProductForRating(1, "Samsung", "Camera")
	product2 := types.AddProductForRating(2, "Iphone", "Operating System")
	product1.AddRating(1, 3, "Average camera quality")
	product2.AddRating(2, 5, "Best OS")
	product1.AddRating(1, 2, "Poor camera quality")
	fmt.Println(product1.String())
	fmt.Println(product2.String())

}
