package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/saucon/belajar_go_lyndacom/api"
)

func main() {

	http.HandleFunc("/", index)

	http.HandleFunc("/api/echo", api.Echo)

	http.HandleFunc("/api/books", api.BookHandlerFunc)

	//port() menyesuaikan pada Environment Variable
	http.ListenAndServe(port(), nil)
}

//Handler untuk "/"
//Contoh : http://localhost:PORT/
func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ucon'S Cloud Native Go")
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
