package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	h := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v %v", time.Now(), r.Method, r.Host)
		io.WriteString(w, "<h1>running</h1>")
	}

	http.HandleFunc("/", h)
	log.Printf("listening 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
