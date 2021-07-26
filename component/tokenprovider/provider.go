package tokenprovider

import "time"

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
	Generate(data Token)
}
