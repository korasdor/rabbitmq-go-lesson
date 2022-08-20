package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request_id := ctx.GetString("x-request-id")
		client_ip := ctx.ClientIP()
		user_agent := ctx.Request.UserAgent()
		method := ctx.Request.Method
		path := ctx.Request.URL.Path

		t := time.Now()

		ctx.Next()

		latency := float32(time.Since(t).Seconds())
		status := ctx.Writer.Status()

		log.Info().Str("request_id", request_id).Str("client_ip", client_ip).
			Str("user_agent", user_agent).Str("method", method).Str("path", path).Float32("latency", latency).Int("status", status).Msg("")
	}
}
