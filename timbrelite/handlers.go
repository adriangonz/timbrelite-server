package timbrelite

import (
	"fmt"
	"net/http"
)

func Ring(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
