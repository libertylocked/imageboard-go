package vrmp

import (
	"fmt"
	"net/http"
	"time"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmplData := map[string]string{
		"time":      time.Now().Format(time.UnixDate),
		"username":  getUserEmail(r),
		"loginURL":  getLoginURL(r, "/"),
		"logoutURL": getLogoutURL(r, ""),
	}
	renderTemplate(w, "index.html", tmplData)
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	email := getUserEmail(r)
	if email == "" {
		redirToLogin(w, r)
		return
	}
	fmt.Fprintf(w, "Hello, %v!", email)
}

func handleEmployeeUpdate(w http.ResponseWriter, r *http.Request) {
	email := getUserEmail(r)
	if email == "" {
		redirToLogin(w, r)
		return
	}
	name := r.FormValue("name")
	bio := r.FormValue("bio")
	updateEmployee(getContext(r), name, bio, email)
}

func handleEmployeeGet(w http.ResponseWriter, r *http.Request) {
	employeeEmail := r.FormValue("email")
	employee, err := getEmployee(getContext(r), employeeEmail)
	if err != nil {
		fmt.Fprintf(w, "error %v", err)
	} else {
		fmt.Fprintf(w, "%v", employee)
	}
}
