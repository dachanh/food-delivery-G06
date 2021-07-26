package appctx

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	SecretKey() string
}
type appCtx struct {
	db        *gorm.DB
	secretkey string
}

func NewAppContext(db *gorm.DB, secretKey string) *appCtx {
	return &appCtx{
		db:        db,
		secretkey: secretKey,
	}
}

func (ctx *appCtx) GetMaiDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretkey
}
