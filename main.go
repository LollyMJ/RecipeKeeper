package main

import (
	"net/http"

	"RecipeKeeper/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.ListenAndServe(":8000", nil)
}
