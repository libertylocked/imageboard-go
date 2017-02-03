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
