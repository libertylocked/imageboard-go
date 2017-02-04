package vrmp

import (
	"fmt"
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

func handleEmployeeUpdate(w http.ResponseWriter, r *http.Request) {
	email := getUserEmail(r)
	if email == "" {
		redirToLogin(w, r)
		return
	}
	name := r.FormValue("name")
	bio := r.FormValue("bio")
	updateEmployee(getContext(r), name, bio, email)
	http.Redirect(w, r, "/", http.StatusFound)
}

func handleEmployeeView(w http.ResponseWriter, r *http.Request) {
	employeeEmail := r.FormValue("email")
	employee, err := getEmployee(getContext(r), employeeEmail)
	if err != nil {
		fmt.Fprintf(w, "error getting employee info. %v", err)
	} else {
		tmplData := map[string]string{
			"name":        employee.Name,
			"bio":         employee.Bio,
			"email":       employee.Email,
			"lastUpdated": employee.LastUpdated.Format(time.UnixDate),
		}
		renderTemplate(w, "employee_view.html", tmplData)
	}
}

func handleEmployeeEdit(w http.ResponseWriter, r *http.Request) {
	email := getUserEmail(r)
	if email == "" {
		redirToLogin(w, r)
		return
	}
	// pre-populate input fields if record exists
	employee, err := getEmployee(getContext(r), email)
	var tmplData map[string]string

	if err != nil {
		tmplData = map[string]string{
			"name":        "<enter your name>",
			"bio":         "<enter your bio>",
			"lastUpdated": "not found",
		}
	} else {
		tmplData = map[string]string{
			"name":        employee.Name,
			"bio":         employee.Bio,
			"lastUpdated": employee.LastUpdated.Format(time.UnixDate),
		}
	}
	tmplData["email"] = email
	renderTemplate(w, "employee_edit.html", tmplData)
}
