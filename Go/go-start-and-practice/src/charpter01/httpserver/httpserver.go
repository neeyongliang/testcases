package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/var/log")))
	http.ListenAndServe(":8080", nil)
}
