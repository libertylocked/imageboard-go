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
	router.HandleFunc("/employeeview", handleEmployeeView).Methods("GET")
	router.HandleFunc("/employeeedit", handleEmployeeEdit).Methods("GET")
	router.HandleFunc("/employeeupdate", handleEmployeeUpdate).Methods("POST")

	return router
}
