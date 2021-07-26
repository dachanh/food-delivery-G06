package userstorage

import (
	"context"
	"github.com/dachanh/food-delivery-G06/common"
	usermodel "github.com/dachanh/food-delivery-G06/module/user/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfor ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())
	for i := range moreInfor {
		db = db.Preload(moreInfor[i])
	}
	var user usermodel.User
	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, err
	}
	return &user, nil
}
