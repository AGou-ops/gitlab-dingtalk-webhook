package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MethodNotAllowed 仅允许/webhooks URI上的POST请求，其他请求一律405
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
