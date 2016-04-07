package main

import "log"

func main() {
	log.Println("gtls", Version(), "build", BuildNumber())

	// we are not actually a web server yet.
	log.Println("Hello World")
}
