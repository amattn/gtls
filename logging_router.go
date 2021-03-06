package main

import (
	"log"
	"net/http"
)

type LoggingRouter struct {
	adminResponder *AdminResponder
	linksResponder *LinksResponder
}

func NewLoggingRouter() *LoggingRouter {
	lr := new(LoggingRouter)
	linkstore := NewLinkStore("localhost", 5432, "tutorial", "changeme")
	lr.adminResponder = NewAdminResponder(linkstore)
	lr.linksResponder = NewLinksResponder(linkstore)
	return lr
}

type BaseResponder struct {
	constructorCanary bool
	linkstore         *LinkStore
}

func MakeBaseResponder(linkstore *LinkStore) BaseResponder {
	return BaseResponder{
		constructorCanary: true,
		linkstore:         linkstore,
	}
}

type SimpleRouteHandler interface {
	Respond(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte)
}

// At a high level, a router inspects a request and routes it to an appropriate subcomponent for handling.
// Here, we just look for a simple prefix
func (router *LoggingRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	url := req.URL.Path
	var code int
	responseBytes := []byte{}
	extra_headers := map[string]string{}

	switch {
	case url == "/":
		code = http.StatusOK
		responseBytes = []byte("<html>Welcome to gtls : <a href=\"/admin/list\">list</a> - <a href=\"/admin/add\">add</a></html>")
	case url == "/admin/add":
		code, extra_headers, responseBytes = router.adminResponder.AddShortlinkFormResponse(req)
	case url == "/admin/post":
		code, extra_headers, responseBytes = router.adminResponder.PostResponse(req)
	case url == "/admin/list":
		code, extra_headers, responseBytes = router.adminResponder.ListAllShortlinks(req)
	default:
		// use the shortlink handler
		code, extra_headers, responseBytes = router.linksResponder.Respond(req)
	}

	for k, v := range extra_headers {
		w.Header().Add(k, v)
	}

	w.WriteHeader(code)
	writtenCount, err := w.Write(responseBytes)
	if err != nil {
		log.Println("error writing response", req, err)
	}
	log.Printf("%s", CommonLogFormat(req, code, writtenCount))
}
