package restaurantstorage

import (
	"context"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
)

func (s *sqlStore) Update(ctx context.Context,
	id int,
	updateData *restaurantmodel.RestaurantUpdate) error {
	var data restaurantmodel.Restaurant
	s.db.Where("id = ?", id).First(&data)
	if err := s.db.Where("id = ?", id).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}
