package userbusiness

import (
	"context"
	"github.com/dachanh/food-delivery-G06/component/tokenprovider"
	usermodel "github.com/dachanh/food-delivery-G06/module/user/model"
	"log"
)

type loginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfor ...string) (*usermodel.User, error)
}
type LoginBusiness struct {
	store  loginStorage
	token  tokenprovider.Provider
	hasher Hasher
	expiry int
}

func NewLoginBusiness(st loginStorage, tk tokenprovider.Provider, hash Hasher, expiry int) *LoginBusiness {
	return &LoginBusiness{
		store:  st,
		token:  tk,
		hasher: hash,
		expiry: expiry,
	}
}

func (l *LoginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := l.store.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrEmailorPasswordInvalid
	}
	passHashed := l.hasher.Hash(data.Password + user.Salt)
	if passHashed != user.Password {
		return nil, usermodel.ErrEmailorPasswordInvalid
	}
	payload := tokenprovider.TokenPayLoad{
		UserID: user.ID,
		Role:   user.Role,
	}

	accessToken, err := l.token.Generate(payload, l.expiry)
	log.Println(accessToken)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
