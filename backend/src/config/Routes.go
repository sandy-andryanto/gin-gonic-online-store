/**
 * This file is part of the Sandy Andryanto Online Store Website.
 *
 * @author     Sandy Andryanto <sandy.andryanto.official@gmail.com>
 * @copyright  2025
 *
 * For the full copyright and license information,
 * please view the LICENSE.md file that was distributed
 * with this source code.
 */

package config

import (
	controllers "backend/src/controllers"
	"backend/src/middleware"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type RouteSource struct {
	Name   string
	Method string
	Auth   bool
	Result func(c *gin.Context)
}

func SetupRoutes(db *gorm.DB) *gin.Engine {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("api/home/ping", controllers.HomePing)
	r.GET("api/home/component", controllers.HomeComponent)
	r.GET("api/home/page", controllers.HomePage)
	r.POST("api/home/newsletter", controllers.HomeNewsletter)

	r.POST("api/auth/login", controllers.AuthLogin)
	r.POST("api/auth/register", controllers.AuthRegister)
	r.GET("api/auth/confirm/:token", controllers.AuthConfirm)
	r.POST("api/auth/email/forgot", controllers.AuthEmailForgot)
	r.POST("api/auth/email/reset/:token", controllers.AuthEmailReset)

	r.GET("api/profile/detail", middleware.AuthorizeJWT(), controllers.ProfileDetail)
	r.GET("api/profile/activity", middleware.AuthorizeJWT(), controllers.ProfileActivity)
	r.GET("api/profile/refresh", middleware.AuthorizeJWT(), controllers.ProfileRefresh)
	r.POST("api/profile/update", middleware.AuthorizeJWT(), controllers.ProfileUpdate)
	r.POST("api/profile/password", middleware.AuthorizeJWT(), controllers.ProfilePassword)
	r.POST("api/profile/upload", middleware.AuthorizeJWT(), controllers.ProfileUpload)

	r.GET("api/shop/filter", controllers.ShopFilter)
	r.GET("api/shop/list", controllers.ShopList)

	r.GET("api/order/list", middleware.AuthorizeJWT(), controllers.OrderList)
	r.GET("api/order/detail/:id", middleware.AuthorizeJWT(), controllers.OrderDetail)
	r.GET("api/order/cancel/:id", middleware.AuthorizeJWT(), controllers.OrderCancel)
	r.GET("api/order/wishlist/:id", middleware.AuthorizeJWT(), controllers.OrderWishlist)
	r.GET("api/order/session", middleware.AuthorizeJWT(), controllers.OrderGetSession)
	r.GET("api/order/cart/:id", middleware.AuthorizeJWT(), controllers.OrderCart)
	r.GET("api/order/review/:id", middleware.AuthorizeJWT(), controllers.OrderListReview)
	r.POST("api/order/review/:id", middleware.AuthorizeJWT(), controllers.OrderCreateReview)
	r.POST("api/order/create/cart/:id", middleware.AuthorizeJWT(), controllers.OrderCreateCart)
	r.GET("api/order/checkout/initial", middleware.AuthorizeJWT(), controllers.OrderCheckoutInitial)
	r.POST("api/order/checkout/submit", middleware.AuthorizeJWT(), controllers.OrderCheckout)

	r.MaxMultipartMemory = 8 << 20
	r.Static("uploads", os.Getenv("UPLOAD_PATH"))
	return r
}
