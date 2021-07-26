package tokenprovider

import (
	"errors"
	"time"
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}
type TokenPayLoad struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
}
type Provider interface {
	Generate(data TokenPayLoad, expiry int) (*Token, error)
	Validate(token string) (*TokenPayLoad, error)
}

var (
	ErrNotFound      = errors.New("token not found")
	ErrEncodingTokne = errors.New("error encoding the token")
	ErrInvalidToken  = errors.New("invalid token provided")
)
