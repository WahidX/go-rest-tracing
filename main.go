package main

import (
	"fmt"
	"net/http"

	"github.com/wahidx/go-rest-sample/web/rest"
)

func main() {
	router := rest.NewRouter()
	port := "8000"

	fmt.Println("Listening to port:", port)
	http.ListenAndServe(":"+port, router)

}