package middlewares

import (
	"fmt"
	"net/http"
	"github.com/houssybadr/lawyermanagement/backend/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		jwtKey:=[]byte(utils.GetJwtSecret())
		stringToken:=ctx.GetHeader("Authorization")
		if stringToken==""{
			ctx.JSON(http.StatusUnauthorized,gin.H{"error":"Missing Authorization header"})
			ctx.Abort()
			return
		}
		stringToken=stringToken[len("Bearer "):]
		token,err:=utils.ParseJwtToken(stringToken,jwtKey)
		fmt.Println(token)
		if err!=nil||!token.Valid{
			ctx.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid token"})
			ctx.Abort()
			return
		}
		claims,ok:=token.Claims.(jwt.MapClaims)
		if !ok{
			ctx.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid token claims"})
			ctx.Abort()
			return
		}
		ctx.Set("user_role",claims["user_role"])
		ctx.Next()
		
	}
}