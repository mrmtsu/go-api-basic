package app

import (
	"fmt"
	"net/http"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}
