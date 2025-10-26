package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get funfando")
}

func HandleRequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	HandleRequest()
}
