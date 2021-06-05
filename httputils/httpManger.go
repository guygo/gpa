package httputils

import (
	"log"
	"net/http"
)

type server struct {
	port string
}

func Serve(adress string) {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(adress, http.FileServer(http.Dir("."))))
}
