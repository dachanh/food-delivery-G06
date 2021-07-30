package middleware

import (
	"context"
	"errors"
	"github.com/dachanh/food-delivery-G06/common"
	"github.com/dachanh/food-delivery-G06/component/appctx"
	"github.com/dachanh/food-delivery-G06/component/tokenprovider/jwt"
	usermodel "github.com/dachanh/food-delivery-G06/module/user/model"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"
	"strings"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, more ...string) (*usermodel.User, error)
}

func RequiredAuth(appCtx appctx.AppContext, authStore AuthenStore) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
	return func(c *gin.Context) {
		_, span := trace.StartSpan(c.Request.Context(), "middleware.authorize")
		defer span.End()
		token, err := extractTokenHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}
		// extract payload get info user
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		//  search info user use userID
		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserID})
		if err != nil {
			panic(err)
		}
		// this method help check status and
		if user.Status == 0 {
			panic(errors.New("user has beeb delete or banned"))
		}
		user.Mask(2)
		c.Set(common.CurrentUser, user)
		c.Next()

	}
}

func extractTokenHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("wrong authen header")
	}
	return parts[1], nil
}
