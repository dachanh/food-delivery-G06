package main

import (
	"github.com/dachanh/food-delivery-G06/component/appctx"
	"github.com/dachanh/food-delivery-G06/middleware"
	ginrestaurant "github.com/dachanh/food-delivery-G06/module/restaurant/transport/ginrestaurant"
	userstorage "github.com/dachanh/food-delivery-G06/module/user/storage"
	ginuser "github.com/dachanh/food-delivery-G06/module/user/transport/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func Activate() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
		return err
	}
	dsn := os.Getenv("DNS")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug()
	appContext := appctx.NewAppContext(db, "helloworld")
	route := gin.Default()
	route.Use(middleware.Recover(appContext))
	userStore := userstorage.NewSqlStore(appContext.GetMaiDBConnection())
	midAuthorize := middleware.RequiredAuth(appContext, userStore)
	v1 := route.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("", ginuser.Register(appContext))
		}
		login := v1.Group("/login")
		{
			login.POST("", ginuser.Login(appContext))
		}
		restaurant := v1.Group("/restaurant", midAuthorize)
		{
			restaurant.POST("", ginrestaurant.CreateRestaurant(appContext))
			restaurant.DELETE("")
		}
	}
	route.GET("/ping", midAuthorize, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	route.Run()
	return nil
}

func main() {
	if err := Activate(); err != nil {
		log.Fatal(err)
	}
}
