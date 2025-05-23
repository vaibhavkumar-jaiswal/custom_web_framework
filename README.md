
# 🌐 Custom HTTP Web Framework in Go

This is a lightweight custom HTTP web framework built entirely using Go's `net` package (without `net/http`). It demonstrates how to create a basic yet functional web server from scratch, complete with routing, middleware, and context handling.

---

## 🚀 Key Features

- 📦 Written using pure Go (`net` package) – no external HTTP libraries
- 🔁 Custom routing system for HTTP methods: `GET`, `POST`, `PUT`, `PATCH`
- 🧩 Middleware support (e.g., logging, authentication)
- 📥 Custom context handling for query parameters, body parsing, etc.
- 🛡️ Basic error handling and panic recovery
- 🔐 TLS-ready with self-signed certificates

---

## 📁 Project Structure

```
.
├── main.go                         # Entry point with route definitions
├── go.mod                          # Go module definition
├── cert.pem                        # Self-signed TLS certificate
├── key.pem                         # Private key for TLS
├── framework/                      # Core of the web framework
│   ├── engine.go                   # Engine that handles routing and TCP listener
│   ├── context.go                  # Custom context for handling requests
│   ├── middlewares.go             # Built-in middleware like logging
│   ├── constants.go               # Constants used throughout the framework
```

---

## 🛠️ Getting Started

### ✅ Prerequisites

- Go 1.18 or higher installed
- Terminal access

### 🧪 Running the Project

```bash
go run main.go
```

You’ll see output like:

```
✅ Server running on :8080
```

### 🔒 To run with TLS:

Uncomment the TLS logic and run:

```bash
go run main.go
```

Ensure `cert.pem` and `key.pem` are in the same directory.

---

## 📚 Example Routes

The framework supports defining routes using methods like `.GET()`, `.POST()`:

```go
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
```

Visit:
- `http://localhost:8080/` → Welcome page
- `http://localhost:8080/hello?name=Vaibhav` → Personalized greeting
- `POST http://localhost:8080/post` → Echoes back posted body

---

## 🧱 Middleware

You can register middleware functions globally using `app.Use()`:

```go
app.Use(framework.LoggingMiddleware)
// app.Use(framework.AuthMiddleware) // example auth middleware (commented)
```

Middleware runs before hitting the final handler, giving you control over logging, authorization, etc.

---

## ⚙️ How It Works

- The framework starts a **TCP server** on the defined port.
- Incoming connections are parsed manually into HTTP requests.
- Custom `Context` object simplifies handler development.
- Middleware functions are chained and executed in order.
- Router maps method + path to the correct handler.

---

## 🧪 Sample Output

```bash
✅ Server running on :8080
📥 [GET] /hello?name=Vaibhav
📤 Response: 200 OK
```

---

## 🧾 License

MIT License. You are free to modify, use, and distribute.

---

## 👨‍💻 Author

**Vaibhavkumar V. Jaiswal**  
Email: [jaiswal.vaibhavkumar45@gmail.com](mailto:jaiswal.vaibhavkumar45@gmail.com)

---

## 📌 TODOs / Enhancements

- [ ] Add support for static file serving
- [ ] Improve error handling with proper HTTP response codes
- [ ] Add unit tests for router and middleware
- [ ] Add request body parsing support for JSON
