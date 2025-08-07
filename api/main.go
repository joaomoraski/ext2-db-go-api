package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {

	r := chi.NewRouter()

	r.Post("/users", CreateUserHandler)
	r.Get("/users", SelectUserHandler)

	fmt.Println("Api rodando na porta 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
