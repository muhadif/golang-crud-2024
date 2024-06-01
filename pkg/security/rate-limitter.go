package security

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-limiter"
	coreErr "golang-crud-2024/core/error"
	"golang-crud-2024/pkg/api"
	"golang-crud-2024/pkg/fault"
)

func RateLimiterMiddleware(store limiter.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()

		_, limit, _, _, err := store.Take(context.Background(), key)

		if err != nil {
			api.ResponseFailed(c, fault.ErrorDictionary(fault.HTTPInternalServiceError, err.Error()))
			c.Abort()
			return
		}

		if limit == 0 {
			api.ResponseFailed(c, fault.ErrorDictionary(fault.HTTPToManyRequest, coreErr.ErrToManyRequest))
			c.Abort()
			return
		}

		c.Next()
	}
}
