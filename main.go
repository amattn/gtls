package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// command line flag variables
var (
	show_h       bool
	show_help    bool
	show_version bool
	listen_addr  string
)

// sensible defaults
const (
	DEFAULT_LISTEN_ADDRESS = ":8080"
)

func init() {
	flag.BoolVar(&show_h, "h", false, "show help message and exit(0)")
	flag.BoolVar(&show_help, "help", false, "show help message and exit(0)")
	flag.BoolVar(&show_version, "version", false, "show version info and exit(0)")
	flag.StringVar(&listen_addr, "addr", DEFAULT_LISTEN_ADDRESS, "addr that our webserver listens on")
}

func main() {
	log.Println("gtls", Version(), "build", BuildNumber())

	flag.Parse()

	if show_version {
		os.Exit(0)
	}

	if show_h || show_help {
		flag.Usage()
		os.Exit(0)
	}

	router := NewLoggingRouter()

	log.Println("Listening from:", listen_addr)
	log.Fatal(http.ListenAndServe(listen_addr, router))
}
