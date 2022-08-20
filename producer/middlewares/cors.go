package middlewares

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "600")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		ctx.Writer.Header().Set("Access-Control-Expose-Header", "Content-Length")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTION" {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
