package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/irvankadhafi/go-boilerplate/pkg/trace"
)

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := uuid.New().String()
		ctx.Set(trace.Key, traceID)
		ctx.Next()
	}
}
