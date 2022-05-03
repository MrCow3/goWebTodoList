package main

import (
	"net/http"
	"api"
)


func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8000", srv)
}
