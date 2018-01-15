package main

import (
	"bytes"
	"io"
	"net/http"
)

//go:generate go-bindata -prefix "frontend/" -pkg main -o bindata.go frontend/...

func static_handler(rw http.ResponseWriter, req *http.Request) {
	var path string = req.URL.Path
	if path == "" {
		path = "frontend/index.html"
	}
	if bs, err := Asset(path); err != nil {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		var reader = bytes.NewBuffer(bs)
		io.Copy(rw, reader)
	}
}

func main() {
	http.Handle("/", http.StripPrefix("/", http.HandlerFunc(static_handler)))
	http.ListenAndServe(":3000", nil)
}
