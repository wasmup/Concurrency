package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"

	"gitlab.com/golang-commonmark/markdown"
)

func main() {
	http.HandleFunc("/render", render)
	log.Println("Serving on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func render(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.RemoteAddr)

	if r.Method != "POST" {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	src, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	md := markdown.New(
		markdown.XHTMLOutput(true),
		markdown.Typographer(true),
		markdown.Linkify(true),
		markdown.Tables(true),
	)

	var buf bytes.Buffer
	if err := md.Render(&buf, src); err != nil {
		log.Printf("error converting markdown: %v", err)
		http.Error(w, "Malformed markdown", http.StatusBadRequest)
		return
	}

	if _, err := io.Copy(w, &buf); err != nil {
		log.Printf("error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
