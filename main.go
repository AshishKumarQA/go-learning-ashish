package main

import (
	"fmt"
	"hello_world/types"
)

func main() {
	ratings := make(map[string]types.RatingsAndReviews)
	ratings["Ashish"] = types.RatingsAndReviews{
		ProductDetails: types.Product{
			ProductId:            1,
			ProductName:          "Iphone",
			ProductSpecForRating: "Camera",
		},
		Rating:   4,
		Comments: []string{"Camera quality is better"},
	}
	ratings["Vignesh"] = types.RatingsAndReviews{
		ProductDetails: types.Product{
			ProductId:            1,
			ProductName:          "Iphone",
			ProductSpecForRating: "Camera",
		},
		Rating:   2,
		Comments: []string{"Camera resolution is not good", "Camera quality compared to samsung latest series is poor"},
	}
	ratings["Akshat"] = types.RatingsAndReviews{
		ProductDetails: types.Product{
			ProductId:            2,
			ProductName:          "MAC M4",
			ProductSpecForRating: "Operating System",
		},
		Rating:   5,
		Comments: []string{"Operating System is very fast"},
	}
	fmt.Println("---------ALL Ratings------------")
	for custmerName, ratingInfo := range ratings {
		fmt.Printf("-------Ratings done by customer %s----------\n", custmerName)
		fmt.Printf("All info about ratings: %v\n", ratingInfo)
		ratingInfo.PrintMessageBasedOnRatings()
		ratingInfo.PrintStarsBasedOnRatings()
	}
}
