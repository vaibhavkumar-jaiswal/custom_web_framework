package framework

import (
	"fmt"
	"time"
)

// LoggingMiddleware logs requests
func LoggingMiddleware(ctx *Context) {
	fmt.Printf("✔ [%s] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), ctx.Method, ctx.Path)
}

// AuthMiddleware checks if "Authorization" header is present
func AuthMiddleware(ctx *Context) {
	if ctx.Headers["Authorization"] != "secret-key" {
		ctx.Response(403, "403 Forbidden: Invalid API Key")
	}
}
