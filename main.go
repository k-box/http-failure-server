package main

import (
	"flag"
	"net/http"
)

var (
	listenAddr  string
	contentType string
)

func init() {
	flag.StringVar(&listenAddr, "http", ":8000", "address to listen on")
	flag.StringVar(&contentType, "mime", "application/pdf", "content type to fake")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", failureHandler)

	http.ListenAndServe(listenAddr, nil)
}

func failureHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "HEAD":
		// fake some content
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Content-Length", "10000001")
		w.WriteHeader(http.StatusOK)
		return
	case "GET":
		// just write a 404 header
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
}
