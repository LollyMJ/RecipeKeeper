package main

import (
	"fmt"
	"net/http"

	"RecipeKeeper/data"
	"RecipeKeeper/handlers"
)

func main() {
	data.FetchAllRecipes()
	fmt.Println(data.AllRecipes)
	http.HandleFunc("/", handlers.HomePage)
	http.ListenAndServe(":8000", nil)
}
