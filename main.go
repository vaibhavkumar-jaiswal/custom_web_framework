package main

import (
	"fmt"

	"gowebframework.vaibhavjaiswal.net/framework"
)

func main() {
	app := framework.New()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("🛑 Recovered from panic:", r)
		}
	}()

	// use framework middleware
	// app.Use(framework.AuthMiddleware)
	app.Use(framework.LoggingMiddleware)

	// Define routes
	app.GET("/", func(ctx *framework.Context) {
		ctx.Response(200, "Welcome to my raw HTTP web framework!")
	})

	app.GET("/hello", func(ctx *framework.Context) {
		name := ctx.QueryParams["name"]
		if name == "" {
			name = "Guest"
		}
		ctx.Response(200, fmt.Sprintf("Hello, %s!", name))
	})

	app.POST("/post", func(ctx *framework.Context) {
		ctx.Response(200, "Received: "+ctx.Body)
	})

	// Start server
	app.Run(":8080")
}
