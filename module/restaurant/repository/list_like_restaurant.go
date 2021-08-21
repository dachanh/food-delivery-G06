package repositoryrestaurant

import (
	"context"
	"github.com/dachanh/food-delivery-G06/common"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type GetLikedCountStore interface {
	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantLikeCountRepo struct {
	restaurantstore      ListRestaurantStore
	likedrestaurantstore GetLikedCountStore
}

func NewListRestaurntLikeCountRepo(restaurantstore ListRestaurantStore, likedrestaurantstore GetLikedCountStore) *listRestaurantLikeCountRepo {
	return &listRestaurantLikeCountRepo{
		restaurantstore:      restaurantstore,
		likedrestaurantstore: likedrestaurantstore,
	}
}

func (biz *listRestaurantLikeCountRepo) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.restaurantstore.ListDataWithCondition(ctx, filter, paging, moreKeys...)
	if err != nil {
		return nil, err
	}
	resIDs := make([]int, len(result))

	for i := range resIDs {
		resIDs[i] = result[i].ID
	}
	if mapResLiked, err := biz.likedrestaurantstore.GetRestaurantLike(ctx, resIDs); err == nil {
		for i := range result {
			result[i].LikeCount = mapResLiked[result[i].ID]
			result[i].User.Mask(false)
		}
	}
	return result, nil
}
