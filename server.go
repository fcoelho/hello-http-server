package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "v204\n")
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
