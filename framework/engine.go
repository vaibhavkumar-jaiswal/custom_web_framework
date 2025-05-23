package framework

import (
	"fmt"
	"net"
)

// Engine is the core framework structure
type Engine struct {
	router *router
}

// New creates a new instance of the framework
func New() *Engine {
	return &Engine{router: newRouter()}
}

// Use adds middleware to the router
func (engine *Engine) Use(middleware MiddlewareFunc) {
	engine.router.use(middleware)
}

// GET defines a route for GET requests
func (engine *Engine) GET(path string, handler HandlerFunc) {
	engine.router.addRoute(GET, path, handler)
}

// POST defines a route for POST requests
func (engine *Engine) POST(path string, handler HandlerFunc) {
	engine.router.addRoute(POST, path, handler)
}

// POST defines a route for POST requests
func (engine *Engine) PUT(path string, handler HandlerFunc) {
	engine.router.addRoute(PUT, path, handler)
}

// POST defines a route for POST requests
func (engine *Engine) PATCH(path string, handler HandlerFunc) {
	engine.router.addRoute(PATCH, path, handler)
}

// Start the server (TCP Listener)
func (engine *Engine) Run(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("❌ failed to start server: %v", err)
	}
	fmt.Println("✅ Server running on", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("❌ Failed to accept connection:", err)
			continue
		}

		go func(conn net.Conn) {
			ctx := newContext(conn)
			engine.router.handle(ctx)
		}(conn)
	}
}
