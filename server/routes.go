package server

import (
	"ainosite/models"
	"ainosite/views/adm"
	"ainosite/views/pub"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) Route(auth *Auth) {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.StripSlashes)
	s.Router.Mount("/res", http.FileServer(http.Dir("")))

	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		pub.Index().Render(context.Background(), w)
	})

	s.Router.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		pub.Login().Render(context.Background(), w)
	})

	s.Router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		correct := r.FormValue("password") == auth.Password
		if !correct {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		auth.Login(w, r)
		http.Redirect(w, r, "/adm", http.StatusFound)
	})

	s.Router.Route("/adm", func(r chi.Router) {
		r.Use(auth.Middleware)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := context.Background()
			moods, err := G[models.Mood](s.DB).Order("date").Find(ctx)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			adm.Index(moods).Render(context.Background(), w)
		})
		
		r.Get("/mood", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/adm/mood/" + time.Now().Format(models.DateFormat), http.StatusFound)
		})

		r.Get("/mood/{date}", func(w http.ResponseWriter, r *http.Request) {
			date, err := time.Parse(models.DateFormat, chi.URLParam(r, "date"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ctx := context.Background()
			mood, err := G[models.Mood](s.DB).Where("date = ?", date).First(ctx)
			if err != nil {
				mood = models.Mood{Date: date}
			}
			w.WriteHeader(http.StatusOK)
			adm.Mood(mood).Render(context.Background(), w)
		})

		r.Get("/mood/{date}/set/{value}", func(w http.ResponseWriter, r *http.Request) {
			date, err := time.Parse(models.DateFormat, chi.URLParam(r, "date"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ctx := context.Background()
			mood, err := G[models.Mood](s.DB).Where("date = ?", date).First(ctx)
			if err != nil {
				mood = models.Mood{Date: date}
			}
			mood.Value, err = strconv.Atoi(chi.URLParam(r, "value"))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if mood.ID == 0 {
				G[models.Mood](s.DB).Create(ctx, &mood)
			} else {
				G[models.Mood](s.DB).Updates(ctx, mood)
			}
			http.Redirect(w, r, "/adm/mood/" + mood.Date.Format(models.DateFormat), http.StatusFound)
		})

		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
			auth.Logout(w, r)
			http.Redirect(w, r, "/", http.StatusFound)
		})
	})
}
