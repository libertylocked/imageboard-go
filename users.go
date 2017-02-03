package vrmp

import (
	"net/http"

	"appengine"
	"appengine/user"
)

func getUserEmail(r *http.Request) string {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u != nil {
		return u.Email
	}
	return ""
}

func getLoginURL(r *http.Request, redirURL string) string {
	c := appengine.NewContext(r)
	url, err := user.LoginURL(c, redirURL)
	if err != nil {
		return ""
	}
	return url
}

func getLogoutURL(r *http.Request, redirURL string) string {
	c := appengine.NewContext(r)
	url, err := user.LogoutURL(c, redirURL)
	if err != nil {
		return ""
	}
	return url
}

func getContext(r *http.Request) appengine.Context {
	return appengine.NewContext(r)
}

// Redirects to login page if not logged in. Otherwise do nothing
func redirToLogin(w http.ResponseWriter, r *http.Request) {
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
}
