package controllers

import (
	"fmt"
	"log"
	"net/http"
)

// Home controller maneja la ruta principal del servidor
func HomeController(w http.ResponseWriter, r *http.Request) {
	log.Printf("HomeController accessed: %s", r.URL.Path)
	fmt.Fprintln(w, "Welcome to the chatbot API")
}