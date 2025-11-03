package main

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type Auth struct {
	Password string
	Store *sessions.CookieStore
}

func NewAuth(sessionKey, password string) *Auth {
	return &Auth{
		Password: password,
		Store: sessions.NewCookieStore([]byte(sessionKey)),
	}
}

func (a *Auth) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := a.Store.Get(r, "session-name")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
      return
		}
		authed, ok := session.Values["authed"].(bool)
		if !ok || !authed {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := a.Store.Get(r, "session-name")
	session.Values["authed"] = true
	session.Save(r, w)
}

func (a *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := a.Store.Get(r, "session-name")
	session.Values["authed"] = false
	session.Save(r, w)
}
