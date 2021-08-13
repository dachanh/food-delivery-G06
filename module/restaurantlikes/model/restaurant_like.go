package restaurantlikesmodel

import (
	"fmt"
	"github.com/dachanh/food-delivery-G06/common"
	"time"
)

type Like struct {
	RestaurantID int                `json:"-" gorm:"column:restaurant_id;"`
	UserID       int                `json:""  gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Like) TableName() string { return "restaurant_likes" }

func (l *Like) GetUserID() int {
	return l.UserID
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"))
}
func ErrCannotDislikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot dislike this restaurant"),
		fmt.Sprintf("ErrCannotDislikeRestaurant"))
}
