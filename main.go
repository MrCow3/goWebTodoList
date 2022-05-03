package main

import (
	"net/http"
	"github.com/MrCow3/goWebTodoList/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8000", srv)
}
