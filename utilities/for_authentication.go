package utilities

import (
	"errors"
	"github.com/goincremental/negroni-sessions"
	"net/http"
)

func GetUserId(r *http.Request) (string, error) {
	session := sessions.GetSession(r)
	user_id := session.Get("user_id")
	if user_id == nil {
		return "", errors.New("No user")
	} else {
		return user_id.(string), nil
	}
}

func AuthenticationHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := GetUserId(r)
		if err != nil {
			w.WriteHeader(403)

		} else {
			next(w, r)

		}

	}
}
