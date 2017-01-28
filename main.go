package main

import (
	"log"
	"net/http"

	"github.com/adriangonz/timbrelite-server/timbrelite"
)

func main() {

	router := timbrelite.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
