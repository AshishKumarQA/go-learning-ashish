package types

import (
	"fmt"

	"github.com/fatih/color"
)

type RatingsAndReviews struct {
	ProductDetails Product
	Rating         int
	Comments       []string
}

type Product struct {
	ProductId            int
	ProductName          string
	ProductSpecForRating string
}

func (r RatingsAndReviews) PrintMessageBasedOnRatings() {
	if r.Rating == 5 || r.Rating == 4 {
		color.Green("Thanks for rating its best")
	} else {
		color.Red("Sorry, we will improve the rating")
	}

}

func (r RatingsAndReviews) PrintStarsBasedOnRatings() {
	switch r.Rating {
	case 1:
		fmt.Println("Ratings: *")
	case 2:
		fmt.Println("Ratings: * *")
	case 3:
		fmt.Println("Ratings: * * *")
	case 4:
		fmt.Println("Ratings: * * * *")
	case 5:
		fmt.Println("Ratings: * * * * *")
	default:
		fmt.Println("Ratings out of range")
	}

}
