package businessrestaurant

import (
	"context"
	"github.com/dachanh/food-delivery-G06/common"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
)

type ListRestaurantRepo interface {
	ListRestaurant(ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantbiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantbiz(repo ListRestaurantRepo) *listRestaurantbiz {
	return &listRestaurantbiz{
		repo: repo,
	}
}
func (biz *listRestaurantbiz) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.repo.ListRestaurant(ctx, filter, paging, "User", "LikedCount")
	if err != nil {
		return nil, err
	}
	return result, nil
}
