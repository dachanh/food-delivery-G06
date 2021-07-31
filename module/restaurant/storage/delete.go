package restaurantstorage

import (
	"context"
	"github.com/dachanh/food-delivery-G06/common"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
)

func (s *sqlStore) Delete(cxt context.Context, id int) error {
	if err := s.db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
