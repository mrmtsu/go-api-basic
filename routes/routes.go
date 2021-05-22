package routes

import (
	"go-api-basic/app"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/", app.Hello)

	http.ListenAndServe(":8000", router)
}
