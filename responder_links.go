package main

import (
	"net/http"
	"strings"
)

type LinksResponder struct {
	constructorCanary bool
	linksDB           map[string]string
}

func NewLinksResponder() *LinksResponder {
	handler := new(LinksResponder)
	handler.constructorCanary = true
	handler.linksDB = make(map[string]string)

	handler.linksDB["a"] = "http://golang.org"
	handler.linksDB["b"] = "http://tour.golang.org"
	handler.linksDB["c"] = "http://gotutorial.net"

	return handler
}

func (responder *LinksResponder) Respond(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte) {
	shortcode := req.URL.Path

	// remove leading slash if necessary
	if strings.HasPrefix(shortcode, "/") {
		shortcode = shortcode[1:]
	}

	longurl := responder.linksDB[shortcode]

	if longurl != "" {
		headers := map[string]string{"Location": longurl}
		return http.StatusMovedPermanently, headers, []byte{}
	}
	return http.StatusNotFound, nil, []byte("Not Found")
}
