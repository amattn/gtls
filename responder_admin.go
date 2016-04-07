package main

import (
	"fmt"
	"html"
	"net/http"
)

type AdminResponder struct {
	BaseResponder
}

func NewAdminResponder(linksDB map[string]string) *AdminResponder {
	handler := new(AdminResponder)
	handler.BaseResponder = MakeBaseResponder(linksDB)
	return handler
}

func (responder *AdminResponder) Respond(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte) {
	url_path := req.URL.Path
	switch url_path {
	case "/admin/post":
		if req.Method == "POST" {
			responder.linksDB[req.FormValue("code")] = req.FormValue("url")
			return http.StatusOK, nil, []byte("ok")
		} else {
			return http.StatusMethodNotAllowed, nil, []byte("Method not allowed")
		}
	}

	response := fmt.Sprintf("Hello World, you came from: %q", html.EscapeString(req.URL.Path))
	return http.StatusOK, nil, []byte(response)
}
