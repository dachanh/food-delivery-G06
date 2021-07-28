package jwt

import (
	"github.com/dachanh/food-delivery-G06/component/tokenprovider"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayLoad `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayLoad, expiry int) (*tokenprovider.Token, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})
	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}
func (j *jwtProvider) Validate(token string) (*tokenprovider.TokenPayLoad, error) {
	res, err := jwt.ParseWithClaims(token, &myClaims{}, func(tk *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}
	// validate token
	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}
	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}
	return &claims.Payload, nil
}
