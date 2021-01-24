package router

import (
	"fmt"
	"message_api/interface/handler"
	mid "message_api/interface/middleware"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// CreateRouter .
func CreateRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("satella coffee"))
	})

	r.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/signup", handler.Signup)
		})

		r.Route("/groups", func(r chi.Router) {
			r.Use(mid.Auth)

			r.Post("/", handler.CreateGroup)

			r.Route("/{group_id}", func(r chi.Router) {
				r.Use(mid.GroupCtx)

				r.Route("/messages", func(r chi.Router) {
					r.Get("/", handler.GroupMessageList)
					r.Post("/", handler.CreateGroupMessage)

					r.Route("/{group_message_id}", func(r chi.Router) {
						r.Use(mid.GroupMessageCtx)

						r.Delete("/", handler.DeleteGroupMessage)
					})
				})
			})
		})

		r.Route("/users", func(r chi.Router) {
			r.Use(mid.Auth)

			r.Route("/{user_id}", func(r chi.Router) {
				r.Use(mid.UserCtx)

				r.Route("/direct_messages", func(r chi.Router) {
					r.Get("/", handler.DirectMessageList)
					r.Post("/", handler.CreateDirectMessage)

					r.Route("/{direct_message_id}", func(r chi.Router) {
						r.Use(mid.DirectMessageCtx)

						r.Delete("/", handler.DeleteDirectMessage)
					})
				})
			})
		})
	})

	return r
}

// PrintRoutes .
func PrintRoutes(r chi.Routes) {
	var printRoutes func(parentPattern string, r chi.Routes)
	printRoutes = func(parentPattern string, r chi.Routes) {
		rts := r.Routes()
		for _, rt := range rts {
			if rt.SubRoutes == nil {
				pattern := parentPattern + rt.Pattern
				pattern = strings.Replace(pattern, "/*", "", -1)
				if strings.HasSuffix(pattern, "/") {
					pattern = pattern[:len(pattern)-1]
				}
				for method := range rt.Handlers {
					fmt.Println(method, pattern)
				}
			} else {
				pat := rt.Pattern

				subRoutes := rt.SubRoutes
				printRoutes(parentPattern+pat, subRoutes)
			}
		}
	}
	printRoutes("", r)
}
