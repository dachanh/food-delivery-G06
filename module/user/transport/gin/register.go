package ginuser

import (
	"github.com/dachanh/food-delivery-G06/component/appctx"
	hasher2 "github.com/dachanh/food-delivery-G06/component/hasher"
	userbusiness "github.com/dachanh/food-delivery-G06/module/user/business"
	usermodel "github.com/dachanh/food-delivery-G06/module/user/model"
	userstorage "github.com/dachanh/food-delivery-G06/module/user/storage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Register(appContext appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appContext.GetMaiDBConnection()
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := userstorage.NewSqlStore(db)
		hasher := hasher2.NewMd5Hash()
		biz := userbusiness.NewUserRegister(store, hasher)
		log.Println(data)
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(1)
		c.JSONP(http.StatusOK, map[string]string{
			"data": data.FakeId.String(),
		})
	}
}
