package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	InitRoutes(r)
	err := Connect()
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("server listening")
	http.ListenAndServe(":5000", r)
}
