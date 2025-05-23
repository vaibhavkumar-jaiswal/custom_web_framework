package framework

// MiddlewareFunc defines the function signature for middleware
type MiddlewareFunc func(*Context)

// HandlerFunc defines a function type for request handling
type HandlerFunc func(*Context)

// Router holds route mappings and middleware
type router struct {
	routes     map[string]map[string]HandlerFunc
	middleware []MiddlewareFunc
}

// NewRouter initializes the router
func newRouter() *router {
	return &router{
		routes:     make(map[string]map[string]HandlerFunc),
		middleware: []MiddlewareFunc{},
	}
}

// Use adds middleware to the router
func (rt *router) use(middleware MiddlewareFunc) {
	rt.middleware = append(rt.middleware, middleware)
}

// AddRoute maps a path to a handler
func (rt *router) addRoute(method string, path string, handler HandlerFunc) {
	if rt.routes[method] == nil {
		rt.routes[method] = make(map[string]HandlerFunc)
	}
	rt.routes[method][path] = handler
}

// Handle incoming requests
func (rt *router) handle(ctx *Context) {
	// Execute middleware before handling the request
	for _, middleware := range rt.middleware {
		middleware(ctx)
	}

	// Route the request to the correct handler
	if handlers, exists := rt.routes[ctx.Method]; exists {
		if handler, found := handlers[ctx.Path]; found {
			handler(ctx)
			return
		}
	}
	ctx.Response(404, "404 Not Found")
}
