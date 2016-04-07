package main

import (
	"log"
	"net/http"
	"strings"
)

type LoggingRouter struct {
	adminResponder *AdminResponder
	// userResponder, photosResponder, etc.
}

func NewLoggingRouter() *LoggingRouter {
	lr := new(LoggingRouter)
	lr.adminResponder = new(AdminResponder)
	return lr
}

type SimpleRouteResponder interface {
	Respond(req *http.Request) (statusCode int, responseBytes []byte)
}

// At a high level, a router inspects a request and routes it to an appropriate subcomponent for handling.
// Here, we just look for a simple prefix
func (router *LoggingRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	url := req.URL.Path
	var code int
	responseBytes := []byte{}

	if strings.HasPrefix(url, "/admin/") {
		// use the admin responder
		code, responseBytes = router.adminResponder.Respond(req)
	} else {
		code = http.StatusBadRequest
		responseBytes = []byte("Bad Request")
	}

	w.WriteHeader(code)
	writtenCount, err := w.Write(responseBytes)
	if err != nil {
		log.Println("error writing response", req, err)
	}
	log.Printf("%s", CommonLogFormat(req, code, writtenCount))
}
