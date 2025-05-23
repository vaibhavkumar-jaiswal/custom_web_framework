package framework

import (
	"fmt"
	"net"
	"strings"
)

// this will handle request & response
type Context struct {
	connection  net.Conn
	Method      string
	Path        string
	Headers     map[string]string
	QueryParams map[string]string
	Body        string
	HttpVer     string
}

func newContext(conn net.Conn) *Context {
	buffer := make([]byte, 1024)

	count, _ := conn.Read(buffer)

	rawRequest := string(buffer[:count])

	// every part in raw http request ends with \r\n
	lines := strings.Split(rawRequest, "\r\n")

	var firstLine []string
	if len(lines) > 0 {
		firstLine = strings.Fields(lines[0]) // First line: GET /path HTTP/1.1
	}

	var method, path, httpVer string
	if len(firstLine) == 3 {
		method, path, httpVer = firstLine[0], firstLine[1], firstLine[2]
	}

	// every header ends with \r\n and there will be a line after header part ends.
	headers := make(map[string]string)

	for _, headerLine := range lines[1:] {
		if headerLine == "" {
			break
		}

		parts := strings.SplitN(headerLine, ": ", 2)
		if len(parts) == 2 {
			headers[parts[0]] = parts[1]
		}
	}

	// QueryParams
	queryParams := make(map[string]string)

	if strings.Contains(path, "?") {
		pathParts := strings.SplitN(path, "?", 2)
		path = pathParts[0]

		queryStr := pathParts[1]
		for _, query := range strings.Split(queryStr, "&") {
			kv := strings.SplitN(query, "=", 2)
			if len(kv) == 2 {
				queryParams[kv[0]] = kv[1]
			}
		}
	}

	// Extract Body (for POST/PUT/PATCH requests)
	body := ""
	if method == POST || method == PUT || method == PATCH {
		body = lines[len(lines)-1] // Last part is the body
	}

	return &Context{
		connection:  conn,
		Method:      method,
		Path:        path,
		Headers:     headers,
		QueryParams: queryParams,
		Body:        body,
		HttpVer:     httpVer,
	}
}

// Send plain text response
func (ctx *Context) Response(status int, message string) {
	response := fmt.Sprintf("HTTP/1.1 %d OK\r\nContent-Length: %d\r\n\r\n%s", status, len(message), message)
	ctx.connection.Write([]byte(response))
	ctx.connection.Close()
}
