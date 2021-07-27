package middleware

import (
	"context"
	"errors"
	"github.com/dachanh/food-delivery-G06/component/appctx"
	"github.com/dachanh/food-delivery-G06/component/tokenprovider/jwt"
	"github.com/gin-gonic/gin"
	"os/user"
	"strings"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, more ...string) (*user.User, error)
}

func RequiredAuth(appCtx appctx.AppContext, authStore AuthenStore) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
	return func(c *gin.Context) {
		token, err := extractTokenHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserID})
		if err != nil {
			panic(err)
		}
		if user.Status == 0 {

		}
	}
}

func extractTokenHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("wrong authen header")
	}
	return parts[1], nil
}
