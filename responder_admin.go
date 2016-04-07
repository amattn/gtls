package main

import (
	"fmt"
	"html"
	"net/http"
)

type AdminResponder struct {
}

func (handler *AdminResponder) Respond(req *http.Request) (statusCode int, responseBytes []byte) {
	response := fmt.Sprintf("Hello World, you came from: %q", html.EscapeString(req.URL.Path))
	return http.StatusOK, []byte(response)
}
