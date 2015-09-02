package main

import (
	"fmt"
	"log"
	"net/http"
)

func whoa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🎉  Whoa! 🎉")
}

func main() {
	http.HandleFunc("/", whoa)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
