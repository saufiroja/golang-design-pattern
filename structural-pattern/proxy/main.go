package main

import "fmt"

func main() {
	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	code, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("URL: %s, Code: %d, Body: %s\n", appStatusURL, code, body)

	code, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("URL: %s, Code: %d, Body: %s\n", appStatusURL, code, body)

	code, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("URL: %s, Code: %d, Body: %s\n", appStatusURL, code, body)

	code, body = nginxServer.handleRequest(createUserURL, "POST")
	fmt.Printf("URL: %s, Code: %d, Body: %s\n", createUserURL, code, body)

	code, body = nginxServer.handleRequest(createUserURL, "POST")
	fmt.Printf("URL: %s, Code: %d, Body: %s\n", createUserURL, code, body)
}

// subject
type server interface {
	handleRequest(string, string) (int, string)
}

// proxy
type Nginx struct {
	application       *application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxServer() *Nginx {
	return &Nginx{
		application:       &application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}

	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}

	n.rateLimiter[url]++

	return true
}

func (n *Nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}

	return n.application.handleRequest(url, method)
}

// real subject
type application struct{}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}

	return 404, "Not Found"
}
