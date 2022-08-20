package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := uuid.New().String()

		ctx.Set("x-request-id", id)
		ctx.Header("x-request-id", id)
		ctx.Next()
	}
}
