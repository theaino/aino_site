package server

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

func (a *Auth) Authed(w http.ResponseWriter, r *http.Request) bool {
	session, err := a.Store.Get(r, "auth-session")
	if err != nil {
		return false
	}
	authed, ok := session.Values["authed"].(bool)
	return authed && ok
}

func (a *Auth) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !a.Authed(w, r) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := a.Store.Get(r, "auth-session")
	session.Values["authed"] = true
	session.Save(r, w)
}

func (a *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := a.Store.Get(r, "auth-session")
	session.Values["authed"] = false
	session.Save(r, w)
}
