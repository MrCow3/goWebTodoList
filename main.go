package main

import (
	"net/http"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8000", srv)
}
