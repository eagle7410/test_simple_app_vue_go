package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":6060"

var messAppUp = "\n\n === App up on http://localhost" + port + " === \n\n"

func main() {

	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "this log pat")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, messAppUp)
	})
	//dsfds
	log.Println(messAppUp)
	log.Fatal(http.ListenAndServe(port, nil))

}
