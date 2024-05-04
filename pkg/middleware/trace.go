package middleware

import (
	"github.com/binus-thesis-team/product-service/pkg/trace"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := uuid.New().String()
		ctx.Set(trace.Key, traceID)
		ctx.Next()
	}
}
