package api

import (
	"fmt"
	"net/http"
)

//Handler untuk "/api/echo"
//Contoh : http://localhost:PORT/api/echo?message=aaaas
func Echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
