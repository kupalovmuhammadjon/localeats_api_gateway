package api

import (
	_ "api_gateway/api/docs"
	v1 "api_gateway/api/handlers/v1"
	"api_gateway/api/middleware"
	"api_gateway/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title LocalEats API Gateway
// @version 1.0
// @description LocalEats is a program to order local and homemade food with quality and precise delivery.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8888
// @BasePath /localeats.uz

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func NewRouter(sysConfig *models.SystemConfig) *gin.Engine {

	handlerV1 := v1.NewHandlerV1(sysConfig)

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	main := router.Group("/localeats.uz")

	main.Use(middleware.JWTMiddleware())

	users := main.Group("/users")

	users.GET("/:id/profile", handlerV1.GetProfile)
	users.PUT("/:id/update", handlerV1.UpdateProfile)
	users.DELETE("/:id/delete", handlerV1.DeleteUser)

	kitchens := main.Group("/kitchens")

	kitchens.POST("/create", handlerV1.CreateKitchen)
	kitchens.PUT("/:id/update", handlerV1.UpdateKitchen)
	kitchens.GET("/:id", handlerV1.GetKitchenById)
	kitchens.GET("/", handlerV1.GetKitchens)
	kitchens.GET("/search", handlerV1.SearchKitchens)
	kitchens.DELETE("/:id/delete", handlerV1.DeleteKitchen)

	dishes := main.Group("/dishes")

	dishes.POST("/create", handlerV1.CreateDish)
	dishes.PUT("/:id/update", handlerV1.UpdateDish)
	dishes.PUT("/:id/updatenutritions", handlerV1.UpdateDishNutritionInfo)
	dishes.GET("/:id", handlerV1.GetDishById)
	dishes.GET("/all/:id", handlerV1.GetDishesByKitchenId)
	dishes.DELETE("/:id/delete", handlerV1.DeleteDish)

	orders := main.Group("/orders")

	orders.POST("/create", handlerV1.CreateOrder)
	orders.PUT("/:id/update", handlerV1.UpdateOrderStatus)
	orders.GET("/:id", handlerV1.GetOrderById)
	orders.GET("/user/:id", handlerV1.GetOrdersForUser)
	orders.GET("/chef/:id", handlerV1.GetOrdersForChef)
	orders.DELETE("/:id/delete", handlerV1.DeleteOrder)

	payments := main.Group("/payments")

	payments.POST("/create", handlerV1.CreatePayment)

	reviews := main.Group("/reviews")

	reviews.POST("/create", handlerV1.CreateReview)
	reviews.GET("/:id", handlerV1.GetReviewsByKitchenId)
	reviews.DELETE("/:id", handlerV1.DeleteReview)

	return router
}
