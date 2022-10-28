package mid

import (
	"context"

	"github.com/google/uuid"
)

const requestIDKey = "request-id"

// const requestIDHeader = "X-Request-Id"

// RequestID read header with key X-Request-Id, if exist that value used to traceID
// if not, generate uuid for traceID
func RequestID(ctx context.Context) context.Context {
	reqID := uuid.New()
	return context.WithValue(ctx, requestIDKey, reqID.String())
}

func GetID(ctx context.Context) string {
	reqID := ctx.Value(requestIDKey)

	if ret, ok := reqID.(string); ok {
		return ret
	}

	return ""
}

// Because echo request context is not included value from echo.Context
// we need to build this method, hiks.
type keyCtx string

func SetIDx(ctx context.Context, requsetID string) context.Context {
	return context.WithValue(ctx, keyCtx(requestIDKey), requsetID)
}

func GetIDx(ctx context.Context) string {
	requestID, ok := ctx.Value(keyCtx(requestIDKey)).(string)
	if !ok {
		return ""
	}
	return requestID
}
