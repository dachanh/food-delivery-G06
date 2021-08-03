package restaurantstorage

import (
	"context"
	"github.com/dachanh/food-delivery-G06/common"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
)

func (s *sqlStore) GetDataWithCondition(ctx context.Context, cond map[string]interface{}) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
