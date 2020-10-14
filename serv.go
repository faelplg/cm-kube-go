package main

import (
	"io"
	"log"
	"net/http"
)

func greeting(msg string) string {
	return "<b>" + msg + "</b>"
}

func main() {
	cerocks := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, greeting("Code.Education Rocks!"))
	}

	http.HandleFunc("/", cerocks)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
