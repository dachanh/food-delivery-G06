package userbusiness

import (
	"context"
	"errors"
	"github.com/dachanh/food-delivery-G06/common"
	usermodel "github.com/dachanh/food-delivery-G06/module/user/model"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfor ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type RegisterBusiness struct {
	registerStorage RegisterStorage
	hash            Hasher
}

func NewUserRegister(storage RegisterStorage, hasher Hasher) *RegisterBusiness {
	return &RegisterBusiness{
		registerStorage: storage,
		hash:            hasher,
	}
}

func (r *RegisterBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := r.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return errors.New("Email is Existed")
	}
	sail := common.GenSalt(75)
	data.Password = r.hash.Hash(sail + data.Password)
	data.Salt = sail

	if err := r.registerStorage.CreateUser(ctx, data); err != nil {
		return err
	}
	return nil
}
