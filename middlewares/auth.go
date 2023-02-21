package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/kkomissarov/beggar/managers/jwtManager"
	"log"
	"net/http"
)

func AuthRequiredMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		user, err := jwtManager.FindUserByToken(token)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
