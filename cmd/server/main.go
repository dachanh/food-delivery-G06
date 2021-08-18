package main

import (
	"github.com/dachanh/food-delivery-G06/common"
	"github.com/dachanh/food-delivery-G06/component/appctx"
	"github.com/dachanh/food-delivery-G06/middleware"
	"github.com/dachanh/food-delivery-G06/module/restaurant/transport/ginrestaurant"
	ginrestaurantlike "github.com/dachanh/food-delivery-G06/module/restaurantlikes/transport/gin"
	userstorage "github.com/dachanh/food-delivery-G06/module/user/storage"
	ginuser "github.com/dachanh/food-delivery-G06/module/user/transport/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
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
			restaurant.PUT("/:id", ginrestaurant.UpdateRestaurant(appContext))
			restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
			restaurant.GET("/:id", ginrestaurant.GetRestaurant(appContext))
			restaurant.GET("", ginrestaurant.ListRestaurant(appContext))
			restaurant.POST("/:id/like", ginrestaurantlike.RestaurantLike(appContext))
			restaurant.DELETE("/:id/dislike", ginrestaurantlike.RestaurantDislike(appContext))
		}
	}
	route.GET("/ping", midAuthorize, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: os.Getenv("JAEGER_AGENT_URL"),
		Process:       jaeger.Process{ServiceName: "Food-Delivery"},
	})
	if err != nil {
		panic(common.ErrInternal(err))
	}
	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(0.4)})

	http.ListenAndServe(":8080", &ochttp.Handler{
		Handler: route,
	})

	return nil
}

func main() {
	if err := Activate(); err != nil {
		log.Fatal(err)
	}
}
