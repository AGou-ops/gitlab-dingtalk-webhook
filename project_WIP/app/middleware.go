package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MethodNotAllowed() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method != http.MethodPost ||
			ctx.Request.RequestURI != "/webhooks" {
			ctx.JSON(http.StatusMethodNotAllowed, gin.H{
				"message": "Method Not Allowed",
			})
			ctx.Abort()
			return
		}
	}
}
