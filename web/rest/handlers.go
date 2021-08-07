package rest

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Println(ctx)
	fmt.Printf("%u", ctx)

	w.Write([]byte("Pong"))
}