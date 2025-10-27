package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckAdminMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context){
		role, exists := ctx.Get("user_role")
		if !exists || strings.ToLower(role.(string))  != "admin" {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			ctx.Abort()
			return
		}
	}
}