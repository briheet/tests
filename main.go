package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const port = ":5001"

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello from, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "briheet")
}

func main() {
	fmt.Printf("testing writer, checkout %s", port)
	log.Fatal(http.ListenAndServe(port, http.HandlerFunc(MyGreetHandler)))
}
