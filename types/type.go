package types

import (
	"fmt"
	"time"
)

type Rating struct {
	productId            int
	productName          string
	productSpecForRating string
	avgRating            float64
	userRatings          []UserRating
	comments             []Comment
}
type UserRating struct {
	userId    int
	rating    int
	createdAt time.Time
}
type Comment struct {
	comment string
}

func (r *Rating) AddRating(userId int, rating int, comment string) {
	r.userRatings = append(r.userRatings, UserRating{userId, rating, time.Now()})
	r.comments = append(r.comments, Comment{comment})
	r.avgRating = (float64(rating) + r.avgRating) / float64(len(r.comments))
}

func (r *Rating) String() string {
	return fmt.Sprintf("ProductId: %v, ProductName: %v, ProductSpecForRating: %s, Avg Rating: %v, User Info's: %v, Comments: %v",
		r.productId, r.productName, r.productSpecForRating, r.avgRating, r.userRatings, r.comments)
}

func (u UserRating) String() string {
	return fmt.Sprintf("userId: %v, rating: %v, createdAt: %v", u.userId, u.rating, u.createdAt)
}
func (c Comment) String() string {
	return fmt.Sprintf("comment: %v", c.comment)
}

func AddProductForRating(productId int, productName string, productSpecForRating string) Rating {
	return Rating{productId: productId, productName: productName, productSpecForRating: productSpecForRating}
}
