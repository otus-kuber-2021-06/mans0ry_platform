package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/app"))))
	http.ListenAndServe(":8000", nil)
}
