package vrmp

import (
	"net/http"
	"time"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// get all employee data
	employees := getAllEmployees(getContext(r))

	tmplData := map[string]interface{}{
		"time":      time.Now().Format(time.UnixDate),
		"username":  getUserEmail(r),
		"loginURL":  getLoginURL(r, "/"),
		"logoutURL": getLogoutURL(r, ""),
		"employees": employees,
	}
	renderTemplate(w, "index.html", tmplData)
}
