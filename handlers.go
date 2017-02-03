package vrmp

import (
	"fmt"
	"net/http"
	"time"

	"appengine"
	"appengine/user"
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
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Hello, %v!", u)
}
