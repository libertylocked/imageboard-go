package vrmp

import (
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	http.Handle("/", getHandlers())
}

func getHandlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handleIndex).Methods("GET")
	router.HandleFunc("/user", handleUser)

	return router
}
