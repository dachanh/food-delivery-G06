package ginuser

import (
	"github.com/dachanh/food-delivery-G06/component/appctx"
	"github.com/dachanh/food-delivery-G06/component/hasher"
	"github.com/dachanh/food-delivery-G06/component/tokenprovider/jwt"
	userbusiness "github.com/dachanh/food-delivery-G06/module/user/business"
	usermodel "github.com/dachanh/food-delivery-G06/module/user/model"
	userstorage "github.com/dachanh/food-delivery-G06/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appContext appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user usermodel.UserLogin
		if err := c.ShouldBind(&user); err != nil {
			panic(err)
		}
		db := appContext.GetMaiDBConnection()
		tkProvider := jwt.NewTokenJWTProvider(appContext.SecretKey())

		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbusiness.NewLoginBusiness(store, tkProvider, md5, 60*60)
		account, err := biz.Login(c.Request.Context(), &user)
		if err != nil {
			panic(err)
		}
		c.JSONP(http.StatusOK, map[string]string{
			"token":      account.Token,
			"expiry":     string(account.Expiry),
			"created_at": account.Created.String(),
		})
	}
}
