package main

import (
	"fmt"
	"log"
	"net/http"
)

func whoa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🎉  Whoa! New text!  🎉")
}

func main() {
	http.HandleFunc("/", whoa)
	log.Println("Away we go!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
