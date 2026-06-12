package middleware

import (
	"time"

	"user-api/internal/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := uuid.New().String()

		start := time.Now()

		c.Set("X-Request-ID", requestID)

		err := c.Next()

		duration := time.Since(start)

		logger.Log.Info(
			"request",
			zap.String("requestId", requestID),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
		)

		return err
	}
}
