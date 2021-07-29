package userbusiness

import (
	"context"
	"github.com/dachanh/food-delivery-G06/common"
	usermodel "github.com/dachanh/food-delivery-G06/module/user/model"
)

type registerStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfor ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type RegisterBusiness struct {
	registerStorage registerStorage
	hash            Hasher
}

func NewUserRegister(storage registerStorage, hasher Hasher) *RegisterBusiness {
	return &RegisterBusiness{
		registerStorage: storage,
		hash:            hasher,
	}
}

func (r *RegisterBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := r.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return usermodel.ErrEmailExisted
	}
	sail := common.GenSalt(50)
	data.Password = r.hash.Hash(data.Password + sail)
	data.Salt = sail
	data.Role = "user"
	if err := r.registerStorage.CreateUser(ctx, data); err != nil {
		return common.NewUnauthorized(err)
	}
	return nil
}
