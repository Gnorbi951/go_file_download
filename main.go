package main

import (
	"net/http"

	"download/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
