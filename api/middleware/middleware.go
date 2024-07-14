package middleware

import (
	"api_gateway/api/handlers/tokens"
	"api_gateway/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context){
		auth := ctx.GetHeader("Authorization")

		if len(auth) == 0{
			log.Println("Empty Authorization header")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Error{
				Message: "No Authorization header",
			})
			return
		}
		_, err := tokens.ExtractClaims(auth, false)
		if err != nil {
			log.Println("Invalid access token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Error{
				Message: "Invalid access token",
			})
			return
		}
		ctx.Next()
	}
}