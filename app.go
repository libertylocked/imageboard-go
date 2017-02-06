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

	router.HandleFunc("/imageupload", handleImageUpload).Methods("GET")
	router.HandleFunc("/imageupload_complete", handleImageUploadComplete).Methods("POST")
	router.HandleFunc("/imageview", handleImageView).Methods("GET")
	router.HandleFunc("/serveimage", handleImageServe).Methods("GET")
	router.HandleFunc("/imagedelete", handleImageDelete).Methods("GET")

	return router
}
