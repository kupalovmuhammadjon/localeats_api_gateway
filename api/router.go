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

// @securityDefinitions.apikey BearerAuth
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


	return router
}
