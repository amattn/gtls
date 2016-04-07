package main

import (
	"log"
	"net/http"
	"strings"
)

type LinksResponder struct {
	BaseResponder
}

func NewLinksResponder(linkstore *LinkStore) *LinksResponder {
	responder := new(LinksResponder)
	responder.BaseResponder = MakeBaseResponder(linkstore)

	responder.linkstore.AddShortlink("a", "http://golang.org")
	responder.linkstore.AddShortlink("b", "http://tour.golang.org")
	responder.linkstore.AddShortlink("c", "http://gotutorial.net")

	return responder
}

func (responder *LinksResponder) Respond(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte) {
	shortcode := req.URL.Path

	// remove leading slash if necessary
	if strings.HasPrefix(shortcode, "/") {
		shortcode = shortcode[1:]
	}

	longurl, err := responder.linkstore.GetShortlink(shortcode)
	if err != nil {
		log.Println(2097280714, err)
		return http.StatusInternalServerError, nil, []byte("Internal Server Error")
	}

	if longurl != "" {
		headers := map[string]string{"Location": longurl}
		return http.StatusMovedPermanently, headers, []byte{}
	}
	return http.StatusNotFound, nil, []byte("Not Found")
}
